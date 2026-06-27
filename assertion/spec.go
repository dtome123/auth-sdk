package assertion

import "time"

// ReplayChecker is the interface for checking JWT replay by jti.
type ReplayChecker interface {
	Exists(jti string) bool
	Set(jti string, ttl time.Duration) error
}
