package internal

import (
	"context"

	"dealls-dating-app/internal/entity"

	"github.com/google/uuid"
)

type UserUC interface {
	// Register creates new user
	Register(ctx context.Context, email, password string) error
	// Login returns user's token
	Login(ctx context.Context, email, password string) (string, error)
	// Get returns single user
	Get(ctx context.Context, id uuid.UUID) (entity.User, error)
	// GetVerifiedStatus returns user's verified status
	GetVerifiedStatus(ctx context.Context, id uuid.UUID) (bool, error)
	// BuyPackage activates user's package
	BuyPackage(ctx context.Context, req entity.BuyPackageRequest) error
}
