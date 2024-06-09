package usecase

import (
	"context"
	"dealls-dating-app/internal"
	"dealls-dating-app/internal/constant"
	"dealls-dating-app/internal/entity"
	"dealls-dating-app/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestSwipeUC_NextProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		userAid = uuid.New()
		userBid = uuid.New()
		userCid = uuid.New()
	)

	type fields struct {
		profileRepo internal.ProfileRepo
		likeRepo    internal.LikeRepo
		swipeRepo   internal.SwipeRepo
		userUC      internal.UserUC
	}
	type args struct {
		req entity.NextProfileRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.NextProfileResponse
		wantErr error
	}{
		{
			name: "success - return next profile",
			fields: fields{
				profileRepo: func() *mocks.MockProfileRepo {
					mock := mocks.NewMockProfileRepo(ctrl)
					mock.EXPECT().Get(gomock.Any(), gomock.Any()).Return([]entity.Profile{
						{
							UserID: userCid,
							Name:   "User C",
							Gender: "female",
						},
					}, nil)
					return mock
				}(),
				swipeRepo: func() *mocks.MockSwipeRepo {
					mock := mocks.NewMockSwipeRepo(ctrl)
					todaySwipedProfiles := []uuid.UUID{userBid}
					mock.EXPECT().TodayProfiles(gomock.Any(), userAid).Return(todaySwipedProfiles, nil)
					mock.EXPECT().Store(gomock.Any(), userAid, userCid).Return(nil)
					return mock
				}(),
				userUC: func() *mocks.MockUserUC {
					mock := mocks.NewMockUserUC(ctrl)
					mock.EXPECT().GetVerifiedStatus(gomock.Any(), userCid).Return(true, nil)
					return mock
				}(),
			},
			args: args{
				req: entity.NextProfileRequest{
					User: entity.User{
						BaseModel: entity.BaseModel{
							ID: userAid,
						},
					},
				},
			},
			want: entity.NextProfileResponse{
				Profile: entity.ProfileResponse{
					UserID:     userCid,
					Name:       "User C",
					Gender:     "female",
					IsVerified: true,
				},
			},
			wantErr: nil,
		},
		{
			name: "error - day limit reached",
			fields: fields{
				profileRepo: func() *mocks.MockProfileRepo {
					mock := mocks.NewMockProfileRepo(ctrl)
					return mock
				}(),
				swipeRepo: func() *mocks.MockSwipeRepo {
					mock := mocks.NewMockSwipeRepo(ctrl)
					todaySwipedProfiles := []uuid.UUID{}
					for _ = range constant.SwipeDailyLimit {
						todaySwipedProfiles = append(todaySwipedProfiles, uuid.New())
					}
					mock.EXPECT().TodayProfiles(gomock.Any(), userAid).Return(todaySwipedProfiles, nil)
					return mock
				}(),
				userUC: func() *mocks.MockUserUC {
					mock := mocks.NewMockUserUC(ctrl)
					return mock
				}(),
			},
			args: args{
				req: entity.NextProfileRequest{
					User: entity.User{
						BaseModel: entity.BaseModel{
							ID: userAid,
						},
					},
				},
			},
			want:    entity.NextProfileResponse{},
			wantErr: constant.ErrSwipeLimitReached,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SwipeUC{
				logger:      zap.NewNop(),
				profileRepo: tt.fields.profileRepo,
				likeRepo:    tt.fields.likeRepo,
				swipeRepo:   tt.fields.swipeRepo,
				userUC:      tt.fields.userUC,
			}
			got, err := s.NextProfile(context.Background(), tt.args.req)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
