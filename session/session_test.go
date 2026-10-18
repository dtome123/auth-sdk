package session

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

type mockCmdable struct {
	redis.Cmdable
	store map[string]string
	ttls  map[string]time.Duration
}

func newMockCmdable() *mockCmdable {
	return &mockCmdable{
		store: make(map[string]string),
		ttls:  make(map[string]time.Duration),
	}
}

func (m *mockCmdable) Get(ctx context.Context, key string) *redis.StringCmd {
	val, ok := m.store[key]
	if !ok {
		return redis.NewStringResult("", redis.Nil)
	}
	return redis.NewStringResult(val, nil)
}

func (m *mockCmdable) Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd {
	switch v := value.(type) {
	case string:
		m.store[key] = v
	case []byte:
		m.store[key] = string(v)
	default:
		m.store[key] = ""
	}
	m.ttls[key] = expiration
	return redis.NewStatusResult("OK", nil)
}

func (m *mockCmdable) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	count := int64(0)
	for _, k := range keys {
		if _, ok := m.store[k]; ok {
			delete(m.store, k)
			delete(m.ttls, k)
			count++
		}
	}
	return redis.NewIntResult(count, nil)
}

type UserPayload struct {
	Role  string   `json:"role"`
	Perms []string `json:"perms"`
}

func TestSessionWithBaseStruct(t *testing.T) {
	mock := newMockCmdable()
	ctx := context.Background()

	client := New[UserPayload](mock)

	sess := Session[UserPayload]{
		BaseSession: BaseSession{
			ID:       "sess-123",
			JTI:      "jti-abc",
			UserID:   "user-001",
			DeviceID: "device-web",
			Type:     "user",
			TTL:      time.Now().Add(1 * time.Hour),
		},
		Data: UserPayload{
			Role:  "admin",
			Perms: []string{"read", "write"},
		},
	}

	// 1. UpsertSession
	err := client.UpsertSession(ctx, sess)
	if err != nil {
		t.Fatalf("UpsertSession failed: %v", err)
	}

	key := BuildSessionKey("user-001", "device-web")
	if _, ok := mock.store[key]; !ok {
		t.Fatalf("expected key %s in store", key)
	}

	// 2. GetSession
	retrieved, err := client.GetSession(ctx, "user-001", "device-web")
	if err != nil {
		t.Fatalf("GetSession failed: %v", err)
	}
	if retrieved == nil {
		t.Fatalf("expected retrieved session, got nil")
	}
	if retrieved.UserID != "user-001" || retrieved.JTI != "jti-abc" || retrieved.Data.Role != "admin" {
		t.Errorf("unexpected retrieved session data: %+v", retrieved)
	}

	// 3. DeleteBySession
	err = client.DeleteBySession(ctx, sess)
	if err != nil {
		t.Fatalf("DeleteBySession failed: %v", err)
	}

	retrievedAfterDel, err := client.GetSession(ctx, "user-001", "device-web")
	if err != nil {
		t.Fatalf("expected nil error on deleted session, got: %v", err)
	}
	if retrievedAfterDel != nil {
		t.Errorf("expected nil session after deletion, got: %+v", retrievedAfterDel)
	}
}

func TestExpiredTTLDeletesSession(t *testing.T) {
	mock := newMockCmdable()
	ctx := context.Background()

	client := New[any](mock)

	key := BuildSessionKey("u-exp", "dev-1")
	mock.store[key] = `{"id":"old"}`

	sess := Session[any]{
		BaseSession: BaseSession{
			UserID:   "u-exp",
			DeviceID: "dev-1",
			TTL:      time.Now().Add(-5 * time.Minute),
		},
	}

	err := client.UpsertSession(ctx, sess)
	if err != nil {
		t.Fatalf("UpsertSession with expired TTL failed: %v", err)
	}

	if _, exists := mock.store[key]; exists {
		t.Errorf("expected key to be deleted when TTL is in past")
	}
}

func TestUninitializedClient(t *testing.T) {
	var client Client[any]
	ctx := context.Background()

	_, err := client.GetSession(ctx, "u", "d")
	if err == nil {
		t.Errorf("expected error with uninitialized client on GetSession")
	}

	err = client.UpsertSession(ctx, Session[any]{})
	if err == nil {
		t.Errorf("expected error with uninitialized client on UpsertSession")
	}

	err = client.DeleteSession(ctx, "u", "d")
	if err == nil {
		t.Errorf("expected error with uninitialized client on DeleteSession")
	}
}
