package usecase

import (
	"context"
	"dealls-dating-app/internal/config"
	"dealls-dating-app/internal/pkg/authentication"
	"errors"
	"time"

	"dealls-dating-app/internal"
	"dealls-dating-app/internal/constant"
	"dealls-dating-app/internal/entity"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserUC struct {
	cfg         *config.Cfg
	logger      *zap.Logger
	userRepo    internal.UserRepo
	packageRepo internal.PackageRepo
}

func NewUserUC(
	cfg *config.Cfg,
	logger *zap.Logger,
	userRepo internal.UserRepo,
	packageRepo internal.PackageRepo,
) internal.UserUC {
	return &UserUC{
		cfg:         cfg,
		logger:      logger,
		userRepo:    userRepo,
		packageRepo: packageRepo,
	}
}

func (u *UserUC) Register(ctx context.Context, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := entity.User{
		Email:    email,
		Password: string(hashedPassword),
	}
	user.ID, err = u.userRepo.Store(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUC) Login(ctx context.Context, email, password string) (string, error) {
	filter := entity.UserFilter{
		User: entity.User{Email: email},
	}
	users, err := u.userRepo.Get(ctx, filter)
	if errors.Is(err, constant.ErrNotFound) || len(users) == 0 {
		return "", constant.ErrUserNotFound
	}
	if err != nil {
		return "", err
	}
	user := users[0]

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	signed, err := authentication.GenerateToken(u.cfg.Auth.Jwt.Secret, user.ID)
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (u *UserUC) Get(ctx context.Context, id uuid.UUID) (entity.User, error) {
	filter := entity.UserFilter{
		User: entity.User{
			BaseModel: entity.BaseModel{
				ID: id,
			},
		},
	}
	users, err := u.userRepo.Get(ctx, filter)
	if errors.Is(err, constant.ErrNotFound) {
		return entity.User{}, constant.ErrNotFound
	}
	if err != nil {
		return entity.User{}, err
	}
	return users[0], nil
}

func (u *UserUC) GetVerifiedStatus(ctx context.Context, id uuid.UUID) (bool, error) {
	filterUser := entity.UserFilter{
		User: entity.User{
			BaseModel: entity.BaseModel{
				ID: id,
			},
		},
	}

	users, err := u.userRepo.Get(ctx, filterUser)
	if err != nil {
		u.logger.Error("failed to get user", zap.Error(err))
		return false, err
	}
	if len(users) == 0 {
		return false, constant.ErrNotFound
	}
	return users[0].IsHaveVerifiedPackage(), nil
}

func (u *UserUC) BuyPackage(ctx context.Context, req entity.BuyPackageRequest) error {
	filterPackage := entity.PackageFilter{
		Package: entity.Package{
			BaseModel: entity.BaseModel{
				ID: req.PackageID,
			},
		},
	}
	packages, err := u.packageRepo.Get(ctx, filterPackage)
	if err != nil {
		u.logger.Error("failed to get package", zap.Error(err))
		return err
	}

	_, err = u.userRepo.StorePackage(ctx, entity.UserPackage{
		UserID:    req.UserID,
		PackageID: packages[0].ID,
		ExpiredAt: time.Now().AddDate(0, 1, 0),
	})
	if err != nil {
		u.logger.Error("failed to store package", zap.Error(err))
		return err
	}
	return nil
}
