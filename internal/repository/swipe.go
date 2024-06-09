package repository

import (
	"context"
	"dealls-dating-app/internal"
	"dealls-dating-app/internal/config"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type SwipeRepo struct {
	cfg   *config.Cfg
	redis *redis.Client
}

func NewSwipeRepo(
	cfg *config.Cfg,
	redis *redis.Client,
) internal.SwipeRepo {
	return &SwipeRepo{
		cfg:   cfg,
		redis: redis,
	}
}

const _swipeLifetime = 24 * time.Hour

func ttlKey(userID string, targetUserID string) string {
	return fmt.Sprintf("swipe:ttl:%s:%s", userID, targetUserID)
}

func setsKey(userID string) string {
	return fmt.Sprintf("swipe:sets:%s", userID)
}

func (p *SwipeRepo) Store(ctx context.Context, userID uuid.UUID, targetUserID uuid.UUID) error {
	err := p.redis.Set(ctx, ttlKey(userID.String(), targetUserID.String()), 1, _swipeLifetime).Err()
	if err != nil {
		return err
	}

	err = p.redis.SAdd(ctx, setsKey(userID.String()), targetUserID.String()).Err()
	if err != nil {
		return err
	}

	return nil
}

func (p *SwipeRepo) TodayProfiles(ctx context.Context, userID uuid.UUID) (uuid.UUIDs, error) {
	err := p.cleanUpSets(ctx, userID.String())
	if err != nil {
		return nil, err
	}

	targetUserIDs, err := p.redis.SMembers(ctx, setsKey(userID.String())).Result()
	if err != nil {
		return nil, err
	}

	var results uuid.UUIDs
	for _, targetUserID := range targetUserIDs {
		res, err := uuid.Parse(targetUserID)
		if err != nil {
			return nil, err
		}
		results = append(results, res)
	}

	return results, nil
}

// cleanUpSets cleans up the sets from expired swipes
func (p *SwipeRepo) cleanUpSets(ctx context.Context, userID string) error {
	targetUserIDs, err := p.redis.SMembers(ctx, setsKey(userID)).Result()
	if err != nil {
		return err
	}

	for _, targetUserID := range targetUserIDs {
		_, err := p.redis.Get(ctx, ttlKey(userID, targetUserID)).Result()
		if errors.Is(err, redis.Nil) {
			err = p.redis.SRem(ctx, setsKey(userID), targetUserID).Err()
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}
