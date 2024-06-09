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

type PackageRepo struct {
	cfg *config.Cfg
	db  *gorm.DB
}

func NewPackageRepo(
	cfg *config.Cfg,
	db *gorm.DB,
) internal.PackageRepo {
	return &PackageRepo{
		cfg: cfg,
		db:  db,
	}
}

func (p *PackageRepo) Get(ctx context.Context, filter entity.PackageFilter) ([]entity.Package, error) {
	tx := p.applyFilter(p.db.WithContext(ctx), filter)

	// multiple records
	if filter.ID == uuid.Nil {
		var packages []entity.Package
		err := tx.Find(&packages).Error
		if err != nil {
			return nil, err
		}
		if len(packages) == 0 {
			return nil, constant.ErrNotFound
		}
		return packages, nil
	}

	// single record
	var pkg entity.Package
	err := p.db.WithContext(ctx).Take(&pkg, filter.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constant.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return []entity.Package{pkg}, nil
}

func (p *PackageRepo) applyFilter(tx *gorm.DB, filter entity.PackageFilter) *gorm.DB {
	if len(filter.IDs) != 0 {
		tx = tx.Where("id IN ?", filter.IDs)
	}
	return tx
}

func (p *PackageRepo) Store(ctx context.Context, pkg entity.Package) (uuid.UUID, error) {
	tx := p.db.WithContext(ctx)

	if pkg.ID == uuid.Nil {
		// create new record
		err := tx.Omit("ID").Create(&pkg).Error
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return uuid.Nil, constant.ErrDuplicateRecord
		}
		if err != nil {
			return uuid.Nil, err
		}

		return pkg.ID, nil
	}

	// update existing record
	tx = tx.Model(&pkg).Updates(&pkg)
	err := tx.Error
	if err != nil {
		return uuid.Nil, err
	}
	if tx.RowsAffected == 0 {
		return uuid.Nil, constant.ErrNotFound
	}

	return pkg.ID, nil
}
