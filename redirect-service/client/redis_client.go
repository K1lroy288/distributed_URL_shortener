package client

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	SaveLink(ctx context.Context, shortCode, longLink string, owner_id int, ttl time.Duration) error
	GetLink(ctx context.Context, shortCode string) (string, error)
}

type redisClient struct {
	client *redis.Client
}

func NewRedisClient(addr string) *redisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &redisClient{client: rdb}
}

func (r *redisClient) SaveLink(ctx context.Context, shortCode, longLink string, owner_id int, ttl time.Duration) error {
	_, err := r.client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.HSet(ctx, shortCode,
			"url", longLink,
			"owner_id", owner_id,
			"created_at", time.Now().Unix(),
		)
		pipe.Expire(ctx, shortCode, ttl)
		return nil
	})

	return err
}

func (r *redisClient) GetLink(ctx context.Context, shortCode string) (string, error) {
	var url string
	if err := r.client.HGet(ctx, shortCode, "url").Scan(url); err != nil {
		return "", err
	}

	return url, nil
}
