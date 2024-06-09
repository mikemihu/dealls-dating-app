package entity

import (
	"dealls-dating-app/internal/constant"

	"github.com/google/uuid"
)

type Profile struct {
	BaseModel
	UserID uuid.UUID       `json:"user_id"`
	Name   string          `json:"name"`
	Gender constant.Gender `json:"gender"`
	Bio    string          `json:"bio"`
}

type ProfileFilter struct {
	Profile
	IDs                uuid.UUIDs
	NotUserIDs         uuid.UUIDs
	ExcludeLikedUserID uuid.UUID
	Limit              int
	OrderByRandom      constant.FilterBool
}

type ProfileUpdateRequest struct {
	UserID uuid.UUID       `json:"-"`
	Name   string          `json:"name"`
	Gender constant.Gender `json:"gender"`
	Bio    string          `json:"bio"`
}

type ProfileResponse struct {
	ID         uuid.UUID       `json:"id"`
	UserID     uuid.UUID       `json:"user_id"`
	Name       string          `json:"name"`
	Gender     constant.Gender `json:"gender"`
	Bio        string          `json:"bio"`
	IsVerified bool            `json:"is_verified"`
}

func (p Profile) ToResponse() ProfileResponse {
	return ProfileResponse{
		ID:     p.ID,
		UserID: p.UserID,
		Name:   p.Name,
		Gender: p.Gender,
		Bio:    p.Bio,
	}
}
