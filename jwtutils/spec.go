package jwtutils

import "time"

type Signer interface {
	Sign(claims Claims, expiry time.Duration) (string, error)
}

type Verifier interface {
	Verify(tokenStr string) (Claims, error)
}

// TokenBuilder defines the interface for generating JWT bearer tokens.
type TokenBuilder interface {
	Build() (BearerToken, error)
}

// ReplayChecker is the interface for checking JWT replay by jti.
type ReplayChecker interface {
	Exists(jti string) bool
	Set(jti string, ttl time.Duration) error
}
