package session

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// DeleteSession removes a session from Redis by userID and deviceID.
func (repo *Client[T]) DeleteSession(ctx context.Context, userID, deviceID string) error {
	if repo.client == nil {
		return fmt.Errorf("redis client is not initialized")
	}

	key := BuildSessionKey(userID, deviceID)
	return repo.client.Del(ctx, key).Err()
}

// Delete is an alias for DeleteSession.
func (repo *Client[T]) Delete(ctx context.Context, userID, deviceID string) error {
	return repo.DeleteSession(ctx, userID, deviceID)
}

// DeleteBySession deletes a session by extracting userID and deviceID directly from Session[T].
func (repo *Client[T]) DeleteBySession(ctx context.Context, session Session[T]) error {
	return repo.DeleteSession(ctx, session.UserID, session.DeviceID)
}

// DeleteSession removes a session from Redis directly via redis.Cmdable.
func DeleteSession(ctx context.Context, client redis.Cmdable, userID, deviceID string) error {
	if client == nil {
		return fmt.Errorf("redis client is not initialized")
	}

	key := BuildSessionKey(userID, deviceID)
	return client.Del(ctx, key).Err()
}
