package usecase

import (
	"context"
	"dealls-dating-app/internal/constant"
	"errors"

	"dealls-dating-app/internal"
	"dealls-dating-app/internal/entity"

	"go.uber.org/zap"
)

type ProfileUC struct {
	logger      *zap.Logger
	profileRepo internal.ProfileRepo
}

func NewProfileUC(
	logger *zap.Logger,
	profileRepo internal.ProfileRepo,
) internal.ProfileUC {
	return &ProfileUC{
		logger:      logger,
		profileRepo: profileRepo,
	}
}

func (p *ProfileUC) GetList(ctx context.Context, filter entity.ProfileFilter) ([]entity.Profile, error) {
	return p.profileRepo.Get(ctx, filter)
}

func (p *ProfileUC) Get(ctx context.Context, filter entity.ProfileFilter) (entity.Profile, error) {
	profiles, err := p.profileRepo.Get(ctx, filter)
	if err != nil {
		return entity.Profile{}, err
	}
	profile := profiles[0]

	return profile, nil
}

func (p *ProfileUC) Update(ctx context.Context, req entity.ProfileUpdateRequest) error {
	filter := entity.ProfileFilter{
		Profile: entity.Profile{
			UserID: req.UserID,
		},
	}
	profiles, err := p.profileRepo.Get(ctx, filter)
	if err != nil && !errors.Is(err, constant.ErrNotFound) {
		p.logger.Error("failed get profile", zap.Error(err))
		return err
	}

	var profile entity.Profile
	if len(profiles) != 0 {
		profile = profiles[0]
	}

	profile.UserID = req.UserID
	profile.Name = req.Name
	profile.Gender = req.Gender
	profile.Bio = req.Bio

	_, err = p.profileRepo.Store(ctx, profile)
	if err != nil {
		p.logger.Error("failed update profile", zap.Error(err))
		return err
	}

	return nil
}
