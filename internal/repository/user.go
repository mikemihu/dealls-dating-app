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

type UserRepo struct {
	cfg *config.Cfg
	db  *gorm.DB
}

func NewUserRepo(
	cfg *config.Cfg,
	db *gorm.DB,
) internal.UserRepo {
	return &UserRepo{
		cfg: cfg,
		db:  db,
	}
}

func (u *UserRepo) Get(ctx context.Context, filter entity.UserFilter) ([]entity.User, error) {
	tx := u.applyFilter(u.db.WithContext(ctx), filter)

	// multiple records
	if filter.ID == uuid.Nil {
		var users []entity.User
		err := tx.Find(&users).Error
		if err != nil {
			return nil, err
		}
		if len(users) == 0 {
			return nil, constant.ErrNotFound
		}
		return users, nil
	}

	// single record
	var user entity.User
	err := u.db.WithContext(ctx).
		Preload("Profile").
		Preload("ActivePackages", "expired_at > NOW()").
		Preload("ActivePackages.Package").
		Take(&user, filter.ID).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constant.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return []entity.User{user}, nil
}

func (u *UserRepo) applyFilter(tx *gorm.DB, filter entity.UserFilter) *gorm.DB {
	if len(filter.Email) != 0 {
		tx = tx.Where("email = ?", filter.Email)
	}
	if len(filter.Password) != 0 {
		tx = tx.Where("password = ?", filter.Password)
	}
	return tx
}

func (u *UserRepo) Store(ctx context.Context, user entity.User) (uuid.UUID, error) {
	tx := u.db.WithContext(ctx)

	if user.ID == uuid.Nil {
		// create new record
		err := tx.Omit("ID").Create(&user).Error
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return uuid.Nil, constant.ErrDuplicateRecord
		}
		if err != nil {
			return uuid.Nil, err
		}

		return user.ID, nil
	}

	// update existing record
	tx = tx.Model(&user).Updates(&user)
	err := tx.Error
	if err != nil {
		return uuid.Nil, err
	}
	if tx.RowsAffected == 0 {
		return uuid.Nil, constant.ErrNotFound
	}

	return user.ID, nil
}

func (u *UserRepo) StorePackage(ctx context.Context, userPackage entity.UserPackage) (uuid.UUID, error) {
	tx := u.db.WithContext(ctx)

	if userPackage.ID == uuid.Nil {
		// create new record
		err := tx.Omit("ID").Create(&userPackage).Error
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return uuid.Nil, constant.ErrDuplicateRecord
		}
		if err != nil {
			return uuid.Nil, err
		}

		return userPackage.ID, nil
	}

	// update existing record
	tx = tx.Model(&userPackage).Updates(&userPackage)
	err := tx.Error
	if err != nil {
		return uuid.Nil, err
	}
	if tx.RowsAffected == 0 {
		return uuid.Nil, constant.ErrNotFound
	}

	return userPackage.ID, nil
}
