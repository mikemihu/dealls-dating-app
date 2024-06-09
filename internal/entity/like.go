package entity

import (
	"dealls-dating-app/internal/constant"

	"github.com/google/uuid"
)

type Like struct {
	BaseModel
	UserID        uuid.UUID `json:"user_id"`
	TargetUserID  uuid.UUID `json:"target_user_id"`
	TargetProfile Profile   `gorm:"foreignKey:TargetUserID;references:UserID"`
}

type LikeFilter struct {
	Like
	IsMutual constant.FilterBool
}

type MutualListRequest struct {
	UserID uuid.UUID
}

type MutualListResponse struct {
	Profiles []ProfileResponse `json:"profiles"`
}
