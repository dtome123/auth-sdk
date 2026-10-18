package session

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// GetSession retrieves and unmarshals a Session[T] from Redis by userID and deviceID.
// If the session is not found (redis.Nil), it returns (nil, nil).
func (repo *Client[T]) GetSession(ctx context.Context, userID, deviceID string) (*Session[T], error) {
	if repo.client == nil {
		return nil, fmt.Errorf("redis client is not initialized")
	}

	key := repo.keyBuilder(userID, deviceID)
	return getSessionByKey[T](ctx, repo.client, key)
}

// Get is an alias for GetSession.
func (repo *Client[T]) Get(ctx context.Context, userID, deviceID string) (*Session[T], error) {
	return repo.GetSession(ctx, userID, deviceID)
}

// GetSession retrieves and unmarshals a Session[T] directly via redis.Cmdable.
func GetSession[T any](ctx context.Context, client redis.Cmdable, userID, deviceID string) (*Session[T], error) {
	if client == nil {
		return nil, fmt.Errorf("redis client is not initialized")
	}

	key := BuildSessionKey(userID, deviceID)
	return getSessionByKey[T](ctx, client, key)
}

func getSessionByKey[T any](ctx context.Context, client redis.Cmdable, key string) (*Session[T], error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get session from redis: %w", err)
	}

	var sess Session[T]
	if err := json.Unmarshal([]byte(val), &sess); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session: %w", err)
	}

	return &sess, nil
}
