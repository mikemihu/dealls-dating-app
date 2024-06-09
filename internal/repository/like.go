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

type LikeRepo struct {
	cfg *config.Cfg
	db  *gorm.DB
}

func NewLikeRepo(
	cfg *config.Cfg,
	db *gorm.DB,
) internal.LikeRepo {
	return &LikeRepo{
		cfg: cfg,
		db:  db,
	}
}

func (p *LikeRepo) Get(ctx context.Context, filter entity.LikeFilter) ([]entity.Like, error) {
	tx := p.applyFilter(p.db.WithContext(ctx), filter)

	// multiple records
	if filter.ID == uuid.Nil {
		var likes []entity.Like
		err := tx.Preload("TargetProfile").
			Find(&likes).
			Error
		if err != nil {
			return nil, err
		}
		if len(likes) == 0 {
			return nil, constant.ErrNotFound
		}
		return likes, nil
	}

	// single record
	var like entity.Like
	err := p.db.WithContext(ctx).
		Preload("TargetProfile").
		Take(&like, filter.ID).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constant.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return []entity.Like{like}, nil
}

func (p *LikeRepo) applyFilter(tx *gorm.DB, filter entity.LikeFilter) *gorm.DB {
	if filter.UserID != uuid.Nil {
		tx = tx.Where("likes.user_id = ?", filter.UserID)
	}
	if filter.TargetUserID != uuid.Nil {
		tx = tx.Where("likes.target_user_id = ?", filter.TargetUserID)
	}
	if filter.IsMutual == constant.FilterBoolTrue {
		tx = tx.Joins("JOIN likes AS l2 ON likes.user_id = l2.target_user_id AND likes.target_user_id = l2.user_id")
	}
	return tx
}

func (p *LikeRepo) Store(ctx context.Context, like entity.Like) (uuid.UUID, error) {
	tx := p.db.WithContext(ctx)

	if like.ID == uuid.Nil {
		// create new record
		err := tx.Omit("ID").Create(&like).Error
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return uuid.Nil, constant.ErrDuplicateRecord
		}
		if err != nil {
			return uuid.Nil, err
		}

		return like.ID, nil
	}

	// update existing record
	tx = tx.Model(&like).Updates(&like)
	err := tx.Error
	if err != nil {
		return uuid.Nil, err
	}
	if tx.RowsAffected == 0 {
		return uuid.Nil, constant.ErrNotFound
	}

	return like.ID, nil
}
