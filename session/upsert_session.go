package session

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// UpsertSession saves or updates a Session[T] in Redis.
// If TTL is in the past, any existing session is deleted.
func (repo *Client[T]) UpsertSession(ctx context.Context, session Session[T]) error {
	if repo.client == nil {
		return fmt.Errorf("redis client is not initialized")
	}

	key := repo.keyBuilder(session.UserID, session.DeviceID)
	data, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("failed to marshal session: %w", err)
	}

	var expiration time.Duration
	if !session.TTL.IsZero() {
		expiration = time.Until(session.TTL)
		if expiration <= 0 {
			return repo.client.Del(ctx, key).Err()
		}
	}

	return repo.client.Set(ctx, key, data, expiration).Err()
}

// UpsertSession saves or updates a Session[T] directly via redis.Cmdable.
func UpsertSession[T any](ctx context.Context, client redis.Cmdable, session Session[T], opts ...Option[T]) error {
	c := NewClient[T](client, opts...)
	return c.UpsertSession(ctx, session)
}
