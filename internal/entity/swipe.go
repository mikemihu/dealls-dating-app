package entity

import "github.com/google/uuid"

type NextProfileRequest struct {
	User User
}

type NextProfileResponse struct {
	Profile ProfileResponse `json:"profile"`
}

type SwipeRequest struct {
	UserID       uuid.UUID `json:"-"`
	TargetUserID uuid.UUID `json:"target_user_id"`
}

type SwipeResponse struct {
	IsMatched bool `json:"is_matched"`
}
