package service

import (
	"context"
	"fmt"
	"log"
	"redirect-service/client"
	"time"
)

type RedirectService struct {
	redisClient     client.RedisClient
	shortenerClient client.ShortenerClient
	cacheTTL        time.Duration
}

func NewRedirectService(redis client.RedisClient, shortener client.ShortenerClient) *RedirectService {
	return &RedirectService{
		redisClient:     redis,
		shortenerClient: shortener,
		cacheTTL:        24 * time.Hour,
	}
}

func (s *RedirectService) Resolve(ctx context.Context, shortCode string) (string, error) {
	if longURL, err := s.redisClient.GetLink(ctx, shortCode); err == nil {
		return longURL, nil
	}

	res, err := s.shortenerClient.Resolve(ctx, shortCode)
	if err != nil {
		return "", fmt.Errorf("error shortener resolve long URL: %w", err)
	}

	err = s.redisClient.SaveLink(ctx, shortCode, res.LongURL, res.Owner_id, s.cacheTTL)
	if err != nil {
		log.Printf("error save URL to redis: %v", err)
	}

	return res.LongURL, nil
}
