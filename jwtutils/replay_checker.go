package jwtutils

import (
	"context"
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
)

// ======================
// In-Memory ReplayChecker (using go-cache)
// ======================

type memoryReplayChecker struct {
	cache *cache.Cache
}

// NewMemoryReplayChecker creates a new in-memory replay checker using go-cache.
func NewMemoryReplayChecker(defaultExpiration time.Duration) ReplayChecker {
	return &memoryReplayChecker{
		cache: cache.New(defaultExpiration, time.Minute), // cleanup interval = 1m
	}
}

// Exists checks if the jti is already in cache.
func (m *memoryReplayChecker) Exists(jti string) bool {
	_, found := m.cache.Get(jti)
	return found
}

// Set stores the jti in cache with expiration TTL.
func (m *memoryReplayChecker) Set(jti string, ttl time.Duration) error {
	m.cache.Set(jti, true, ttl)
	return nil
}

// ======================
// Redis ReplayChecker
// ======================

type redisReplayChecker struct {
	client *redis.Client
	prefix string
}

// NewRedisReplayChecker creates a new Redis-based replay checker.
func NewRedisReplayChecker(client *redis.Client) ReplayChecker {
	return &redisReplayChecker{
		client: client,
		prefix: "jwtjti:",
	}
}

// Exists checks if jti already exists in Redis.
func (r *redisReplayChecker) Exists(jti string) bool {
	ctx := context.Background()
	key := r.prefix + jti

	exists, err := r.client.Exists(ctx, key).Result()
	return err == nil && exists > 0
}

// Set stores the jti in Redis with expiration TTL.
func (r *redisReplayChecker) Set(jti string, ttl time.Duration) error {
	ctx := context.Background()
	key := r.prefix + jti

	ok, err := r.client.SetNX(ctx, key, "1", ttl).Result()
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("replay detected: jti already used (SetNX failed)")
	}
	return nil
}
