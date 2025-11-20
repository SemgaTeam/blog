package repository

import (
	e "github.com/SemgaTeam/blog/internal/error"
	"github.com/redis/go-redis/v9"

	"context"
	"time"
)

type RedisRepository interface {
	GetToken(context.Context, string) (string, error)
	SetToken(context.Context, string, string, time.Duration) error
	DeleteToken(context.Context, string) (int64, error)
}

type redisRepository struct {
	rdb *redis.Client
}

func NewRedisRepository(rdb *redis.Client) RedisRepository {
	return &redisRepository{
		rdb: rdb,
	}
}

func (r *redisRepository) GetToken(ctx context.Context, key string) (string, error) {
	v, err := r.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", e.ErrRedisTokenNotFound
	} else if err != nil {
		return "", e.Internal(err)
	}

	return v, nil
}

func (r *redisRepository) SetToken(ctx context.Context, key, value string, expiration time.Duration) error {
	if err := r.rdb.Set(ctx, key, value, expiration).Err(); err != nil {
		return e.ErrRedisTokenSetFailed
	}

	return nil
}

func (r *redisRepository) DeleteToken(ctx context.Context, key string) (int64, error) {
	countDeleted, err := r.rdb.Del(ctx, key).Result()
	if err != nil {
		return 0, e.ErrRedisTokenDeleteFailed
	}

	return countDeleted, nil
}
