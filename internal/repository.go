package internal

import (
	"context"

	"dealls-dating-app/internal/entity"

	"github.com/google/uuid"
)

type UserRepo interface {
	// Get gets multiple or single row if id provided
	Get(ctx context.Context, filter entity.UserFilter) ([]entity.User, error)
	// Store create new user record
	Store(ctx context.Context, user entity.User) (uuid.UUID, error)
	// StorePackage stores a user's active package
	StorePackage(ctx context.Context, userPackage entity.UserPackage) (uuid.UUID, error)
}

type ProfileRepo interface {
	// Get gets multiple or single row if id provided
	Get(ctx context.Context, filter entity.ProfileFilter) ([]entity.Profile, error)
	// Store create new profile record or update if id provided, will return affected id
	Store(ctx context.Context, profile entity.Profile) (uuid.UUID, error)
}

type PackageRepo interface {
	// Get gets multiple or single row if id provided
	Get(ctx context.Context, filter entity.PackageFilter) ([]entity.Package, error)
	// Store create new package record or update if id provided, will return affected id
	Store(ctx context.Context, pkg entity.Package) (uuid.UUID, error)
}

type LikeRepo interface {
	// Get gets multiple or single row if id provided
	Get(ctx context.Context, filter entity.LikeFilter) ([]entity.Like, error)
	// Store create new like record or update if id provided, will return affected id
	Store(ctx context.Context, like entity.Like) (uuid.UUID, error)
}

type SwipeRepo interface {
	// Store stores a swipe record
	Store(ctx context.Context, userID uuid.UUID, targetUserID uuid.UUID) error
	// TodayProfiles gets all userIDs that swiped today
	TodayProfiles(ctx context.Context, userID uuid.UUID) (uuid.UUIDs, error)
}
