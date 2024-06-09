package usecase

import (
	"context"
	"dealls-dating-app/internal"
	"dealls-dating-app/internal/config"
	"dealls-dating-app/internal/constant"
	"dealls-dating-app/internal/entity"
	"errors"

	"go.uber.org/zap"
)

type LikeUC struct {
	cfg      *config.Cfg
	logger   *zap.Logger
	likeRepo internal.LikeRepo
	userUC   internal.UserUC
}

func NewLikeUC(
	cfg *config.Cfg,
	logger *zap.Logger,
	likeRepo internal.LikeRepo,
	userUC internal.UserUC,
) internal.LikeUC {
	return &LikeUC{
		cfg:      cfg,
		logger:   logger,
		likeRepo: likeRepo,
		userUC:   userUC,
	}
}

func (s *LikeUC) MutualList(ctx context.Context, req entity.MutualListRequest) (entity.MutualListResponse, error) {
	filter := entity.LikeFilter{
		Like: entity.Like{
			UserID: req.UserID,
		},
		IsMutual: constant.FilterBoolTrue,
	}
	likes, err := s.likeRepo.Get(ctx, filter)
	if err != nil && !errors.Is(err, constant.ErrNotFound) {
		s.logger.Error("failed to get likes", zap.Error(err))
		return entity.MutualListResponse{}, err
	}

	resp := entity.MutualListResponse{
		Profiles: []entity.ProfileResponse{},
	}
	for _, like := range likes {
		profileResp := like.TargetProfile.ToResponse()

		isVerified, err := s.userUC.GetVerifiedStatus(ctx, like.TargetProfile.UserID)
		if err != nil {
			s.logger.Error("failed to get verified status", zap.Error(err))
			return entity.MutualListResponse{}, err
		}

		profileResp.IsVerified = isVerified
		resp.Profiles = append(resp.Profiles, profileResp)
	}
	return resp, nil
}
