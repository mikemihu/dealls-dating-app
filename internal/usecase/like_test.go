package usecase

import (
	"context"
	"dealls-dating-app/internal"
	"dealls-dating-app/internal/entity"
	"dealls-dating-app/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestLikeUC_MutualList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		userAid = uuid.New()
		userBid = uuid.New()
		userCid = uuid.New()
	)

	type fields struct {
		likeRepo internal.LikeRepo
		userUC   internal.UserUC
	}
	type args struct {
		req entity.MutualListRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.MutualListResponse
		wantErr error
	}{
		{
			name: "success - few mutual",
			fields: fields{
				likeRepo: func() *mocks.MockLikeRepo {
					mock := mocks.NewMockLikeRepo(ctrl)
					mutual := []entity.Like{
						{
							UserID:       userAid,
							TargetUserID: userBid,
							TargetProfile: entity.Profile{
								UserID: userBid,
								Name:   "User B",
								Gender: "female",
							},
						},
						{
							UserID:       userAid,
							TargetUserID: userCid,
							TargetProfile: entity.Profile{
								UserID: userCid,
								Name:   "User C",
								Gender: "female",
							},
						},
					}
					mock.EXPECT().Get(gomock.Any(), gomock.Any()).Return(mutual, nil)
					return mock
				}(),
				userUC: func() *mocks.MockUserUC {
					mock := mocks.NewMockUserUC(ctrl)
					mock.EXPECT().GetVerifiedStatus(gomock.Any(), userBid).Return(true, nil)
					mock.EXPECT().GetVerifiedStatus(gomock.Any(), userCid).Return(false, nil)
					return mock
				}(),
			},
			args: args{
				req: entity.MutualListRequest{
					UserID: userAid,
				},
			},
			want: entity.MutualListResponse{
				Profiles: []entity.ProfileResponse{
					{
						UserID:     userBid,
						Name:       "User B",
						Gender:     "female",
						IsVerified: true,
					},
					{
						UserID:     userCid,
						Name:       "User C",
						Gender:     "female",
						IsVerified: false,
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "success - no mutual",
			fields: fields{
				likeRepo: func() *mocks.MockLikeRepo {
					mock := mocks.NewMockLikeRepo(ctrl)
					mutual := []entity.Like{}
					mock.EXPECT().Get(gomock.Any(), gomock.Any()).Return(mutual, nil)
					return mock
				}(),
				userUC: func() *mocks.MockUserUC {
					mock := mocks.NewMockUserUC(ctrl)
					return mock
				}(),
			},
			args: args{
				req: entity.MutualListRequest{
					UserID: userAid,
				},
			},
			want: entity.MutualListResponse{
				Profiles: []entity.ProfileResponse{},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LikeUC{
				logger:   zap.NewNop(),
				likeRepo: tt.fields.likeRepo,
				userUC:   tt.fields.userUC,
			}
			got, err := s.MutualList(context.Background(), tt.args.req)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
