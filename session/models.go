package session

import "time"

// BaseSession contains the core standard fields for any authenticated session.
type BaseSession struct {
	ID       string    `json:"id,omitempty"`
	JTI      string    `json:"jti,omitempty"`
	UserID   string    `json:"user_id"`
	DeviceID string    `json:"device_id"`
	Type     string    `json:"type,omitempty"`
	TTL      time.Time `json:"ttl"`
}

// Session is a generic session model embedding BaseSession with custom Data payload.
type Session[T any] struct {
	BaseSession
	Data T `json:"data,omitempty"`
}
