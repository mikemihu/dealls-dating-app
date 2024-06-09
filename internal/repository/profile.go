package repository

import (
	"context"
	"errors"

	"dealls-dating-app/internal"
	"dealls-dating-app/internal/config"
	"dealls-dating-app/internal/constant"
	"dealls-dating-app/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfileRepo struct {
	cfg *config.Cfg
	db  *gorm.DB
}

func NewProfileRepo(
	cfg *config.Cfg,
	db *gorm.DB,
) internal.ProfileRepo {
	return &ProfileRepo{
		cfg: cfg,
		db:  db,
	}
}

func (p *ProfileRepo) Get(ctx context.Context, filter entity.ProfileFilter) ([]entity.Profile, error) {
	tx := p.applyFilter(p.db.WithContext(ctx), filter)

	// single / multiple records
	var profiles []entity.Profile
	err := tx.Find(&profiles).Error
	if err != nil {
		return nil, err
	}
	if len(profiles) == 0 {
		return nil, constant.ErrNotFound
	}
	return profiles, nil
}

func (p *ProfileRepo) applyFilter(tx *gorm.DB, filter entity.ProfileFilter) *gorm.DB {
	if filter.ID != uuid.Nil {
		tx = tx.Where("id = ?", filter.ID)
	}
	if len(filter.IDs) != 0 {
		tx = tx.Where("id IN ?", filter.IDs)
	}
	if filter.UserID != uuid.Nil {
		tx = tx.Where("user_id = ?", filter.UserID)
	}
	if len(filter.NotUserIDs) != 0 {
		tx = tx.Where("user_id NOT IN ?", filter.NotUserIDs)
	}
	if filter.ExcludeLikedUserID != uuid.Nil {
		subQuery := p.db.Table("likes").
			Select("target_user_id").
			Where("user_id = ?", filter.ExcludeLikedUserID)
		tx = tx.Where("user_id NOT IN (?)", subQuery)
	}

	if filter.OrderByRandom == constant.FilterBoolTrue {
		tx = tx.Order("RANDOM()")
	}
	if filter.Limit != 0 {
		tx = tx.Limit(filter.Limit)
	}

	return tx
}

func (p *ProfileRepo) Store(ctx context.Context, profile entity.Profile) (uuid.UUID, error) {
	tx := p.db.WithContext(ctx)

	if profile.ID == uuid.Nil {
		// create new record
		err := tx.Omit("ID").Create(&profile).Error
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return uuid.Nil, constant.ErrDuplicateRecord
		}
		if err != nil {
			return uuid.Nil, err
		}

		return profile.ID, nil
	}

	// update existing record
	tx = tx.Model(&profile).Updates(&profile)
	err := tx.Error
	if err != nil {
		return uuid.Nil, err
	}
	if tx.RowsAffected == 0 {
		return uuid.Nil, constant.ErrNotFound
	}

	return profile.ID, nil
}
