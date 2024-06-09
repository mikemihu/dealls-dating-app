package entity

import "github.com/google/uuid"

type Package struct {
	BaseModel
	Name             string  `json:"name"`
	IsUnlimitedSwipe bool    `json:"is_unlimited_swipe"`
	IsVerified       bool    `json:"is_verified"`
	Price            float64 `json:"price"`
}

type PackageFilter struct {
	Package
	IDs uuid.UUIDs
}
