package entity

import (
	"dealls-dating-app/internal/constant"
	"time"

	"github.com/google/uuid"
)

type User struct {
	BaseModel
	Email          string        `json:"email"`
	Password       string        `json:"-"`
	Profile        Profile       `gorm:"foreignKey:UserID" json:"profile"`
	ActivePackages []UserPackage `gorm:"foreignKey:UserID" json:"active_packages"`
}

type UserPackage struct {
	BaseModel
	UserID    uuid.UUID `json:"user_id"`
	PackageID uuid.UUID `json:"package_id"`
	ExpiredAt time.Time `json:"expired_at"`
	Package   Package   `gorm:"foreignKey:PackageID" json:"package"`
}

type UserFilter struct {
	User
}

// SwipeLimit return the swipe limit for the user based on user's active packages
func (u *User) SwipeLimit() int {
	for _, p := range u.ActivePackages {
		if p.Package.IsUnlimitedSwipe {
			// fair limit for the unlimited swipe
			return 99999
		}
	}
	return constant.SwipeDailyLimit
}

// IsHaveVerifiedPackage check if the user has verified package
func (u *User) IsHaveVerifiedPackage() bool {
	for _, p := range u.ActivePackages {
		if p.Package.IsVerified {
			return true
		}
	}
	return false
}

type BuyPackageRequest struct {
	UserID    uuid.UUID `json:"-"`
	PackageID uuid.UUID `json:"package_id"`
}
