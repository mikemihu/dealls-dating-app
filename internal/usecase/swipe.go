package usecase

import (
	"context"
	"dealls-dating-app/internal"
	"dealls-dating-app/internal/config"
	"dealls-dating-app/internal/constant"
	"dealls-dating-app/internal/entity"
	"errors"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type SwipeUC struct {
	cfg         *config.Cfg
	logger      *zap.Logger
	redis       *redis.Client
	profileRepo internal.ProfileRepo
	likeRepo    internal.LikeRepo
	swipeRepo   internal.SwipeRepo
	userUC      internal.UserUC
}

func NewSwipeUC(
	cfg *config.Cfg,
	logger *zap.Logger,
	redis *redis.Client,
	profileRepo internal.ProfileRepo,
	likeRepo internal.LikeRepo,
	swipeRepo internal.SwipeRepo,
	userUC internal.UserUC,
) internal.SwipeUC {
	return &SwipeUC{
		cfg:         cfg,
		logger:      logger,
		redis:       redis,
		profileRepo: profileRepo,
		likeRepo:    likeRepo,
		swipeRepo:   swipeRepo,
		userUC:      userUC,
	}
}

func (s *SwipeUC) NextProfile(ctx context.Context, req entity.NextProfileRequest) (entity.NextProfileResponse, error) {
	var excludeUserIDs uuid.UUIDs
	excludeUserIDs = append(excludeUserIDs, req.User.ID)

	// get today's user id from redis
	swipedUserIDs, err := s.swipeRepo.TodayProfiles(ctx, req.User.ID)
	if err != nil {
		s.logger.Error("failed to get swiped user ids", zap.Error(err))
		return entity.NextProfileResponse{}, err
	}
	excludeUserIDs = append(excludeUserIDs, swipedUserIDs...)

	// check swipe limit
	if len(swipedUserIDs) >= req.User.SwipeLimit() {
		return entity.NextProfileResponse{}, constant.ErrSwipeLimitReached
	}

	// query next profile
	filter := entity.ProfileFilter{
		NotUserIDs:         excludeUserIDs,
		ExcludeLikedUserID: req.User.ID,
		OrderByRandom:      constant.FilterBoolTrue,
		Limit:              1,
	}
	nextProfiles, err := s.profileRepo.Get(ctx, filter)
	if err != nil {
		s.logger.Error("failed to get next profile", zap.Error(err))
		return entity.NextProfileResponse{}, err
	}
	nextProfile := nextProfiles[0]

	// assume user do swiped when requesting next profile, prevent abuse
	err = s.swipeRepo.Store(ctx, req.User.ID, nextProfile.UserID)
	if err != nil {
		s.logger.Error("failed to store swipe", zap.Error(err))
		return entity.NextProfileResponse{}, err
	}

	// get is verified status
	isVerified, err := s.userUC.GetVerifiedStatus(ctx, nextProfile.UserID)
	if err != nil {
		s.logger.Error("failed to get verified status", zap.Error(err))
		return entity.NextProfileResponse{}, err
	}

	// response
	profileResp := nextProfile.ToResponse()
	profileResp.IsVerified = isVerified

	resp := entity.NextProfileResponse{
		Profile: profileResp,
	}
	return resp, nil
}

func (s *SwipeUC) Like(ctx context.Context, req entity.SwipeRequest) (entity.SwipeResponse, error) {
	// store to db
	_, err := s.likeRepo.Store(ctx, entity.Like{
		UserID:       req.UserID,
		TargetUserID: req.TargetUserID,
	})
	if err != nil {
		s.logger.Error("failed to store likes", zap.Error(err))
		return entity.SwipeResponse{}, err
	}

	// check if matched profile
	filterLike := entity.LikeFilter{
		Like: entity.Like{
			UserID:       req.TargetUserID,
			TargetUserID: req.UserID,
		},
	}
	existLike, err := s.likeRepo.Get(ctx, filterLike)
	if err != nil && !errors.Is(err, constant.ErrNotFound) {
		s.logger.Error("failed to get matched profile", zap.Error(err))
		return entity.SwipeResponse{}, err
	}

	resp := entity.SwipeResponse{}
	if len(existLike) != 0 {
		resp.IsMatched = true
	}
	return resp, nil
}
