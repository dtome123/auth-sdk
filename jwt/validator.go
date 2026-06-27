package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dtome123/auth-sdk/assertion"
)

type validateOptions struct {
	expectedAudience     string
	requireIssuerSubject bool
	replayChecker        assertion.ReplayChecker
	clockSkew            time.Duration
}

// ValidateOption is a function option type for Validate.
type ValidateOption func(*validateOptions)

// Option setters

func WithExpectedAudience(aud string) ValidateOption {
	return func(o *validateOptions) {
		o.expectedAudience = aud
	}
}

func WithRequireIssuerSubject(require bool) ValidateOption {
	return func(o *validateOptions) {
		o.requireIssuerSubject = require
	}
}

func WithReplayChecker(rc assertion.ReplayChecker) ValidateOption {
	return func(o *validateOptions) {
		o.replayChecker = rc
	}
}

func WithClockSkew(skew time.Duration) ValidateOption {
	return func(o *validateOptions) {
		o.clockSkew = skew
	}
}

// Validate performs sequential, allocation-free claim validation.
func (c OauthClaims) Validate(opts ...ValidateOption) error {
	var o validateOptions
	for _, opt := range opts {
		opt(&o)
	}

	// 1. Check required fields
	if c.Iss == "" {
		return errors.New("missing issuer (iss) claim")
	}
	if c.Sub == "" {
		return errors.New("missing subject (sub) claim")
	}
	if c.Aud == "" && len(c.Audiences) == 0 {
		return errors.New("missing audience (aud) claim")
	}

	// 2. Check issuer == subject (Service token requirement)
	if o.requireIssuerSubject && c.Iss != c.Sub {
		return errors.New("issuer (iss) and subject (sub) do not match")
	}

	// 3. Check expiration & iat with clock skew
	now := time.Now().Unix()
	skewSec := int64(o.clockSkew.Seconds())

	if c.Exp != 0 && (now-skewSec) > c.Exp {
		return errors.New("token expired")
	}
	if c.Iat != 0 && (now+skewSec) < c.Iat {
		return errors.New("token used before issued")
	}

	// 4. Check audience
	if o.expectedAudience != "" {
		auds := c.Audiences
		if len(auds) == 0 && c.Aud != "" {
			auds = []string{c.Aud}
		}
		matched := false
		for _, a := range auds {
			if a == o.expectedAudience {
				matched = true
				break
			}
		}
		if !matched {
			return fmt.Errorf("audience mismatch: want %q, got %v", o.expectedAudience, auds)
		}
	}

	// 5. Check replay attack
	if o.replayChecker != nil {
		if c.Jti == "" {
			return errors.New("missing jti claim required for replay check")
		}
		if o.replayChecker.Exists(c.Jti) {
			return errors.New("replay detected: jti already used")
		}
		ttlSec := c.Exp - now
		if ttlSec <= 0 {
			return errors.New("cannot check replay: token is already expired")
		}
		if err := o.replayChecker.Set(c.Jti, time.Duration(ttlSec)*time.Second); err != nil {
			return err
		}
	}

	return nil
}
