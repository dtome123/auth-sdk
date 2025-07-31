package jwtutils

import (
	"errors"
	"fmt"
	"time"
)

// validateOptions holds options for Validate method.
type validateOptions struct {
	expectedAudience     string
	allowedAudiences     map[string]struct{}
	requireIssuerSubject bool
	replayChecker        ReplayChecker
	timeNow              func() time.Time
	customValidators     []validateFunc
}

// ValidateOption is a function option type for Validate.
type ValidateOption func(*validateOptions)

// Option setters

func WithExpectedAudience(aud string) ValidateOption {
	return func(o *validateOptions) {
		o.expectedAudience = aud
	}
}

func WithAllowedAudiences(auds []string) ValidateOption {
	return func(o *validateOptions) {
		o.allowedAudiences = make(map[string]struct{}, len(auds))
		for _, a := range auds {
			o.allowedAudiences[a] = struct{}{}
		}
	}
}

func WithRequireIssuerSubject(require bool) ValidateOption {
	return func(o *validateOptions) {
		o.requireIssuerSubject = require
	}
}

func WithReplayChecker(rc ReplayChecker) ValidateOption {
	return func(o *validateOptions) {
		o.replayChecker = rc
	}
}

func WithTimeNow(fn func() time.Time) ValidateOption {
	return func(o *validateOptions) {
		o.timeNow = fn
	}
}

func WithCustomValidators(validators ...validateFunc) ValidateOption {
	return func(o *validateOptions) {
		o.customValidators = append(o.customValidators, validators...)
	}
}

// validateFunc defines signature of individual validators
type validateFunc func(OauthClaims, *validateOptions) error

// Default validators

func validateRequiredFields(c OauthClaims, o *validateOptions) error {
	if c.Iss == "" {
		return errors.New("missing issuer (iss) claim")
	}
	if c.Sub == "" {
		return errors.New("missing subject (sub) claim")
	}
	if len(c.Aud) == 0 {
		return errors.New("missing audience (aud) claim")
	}
	return nil
}

func validateIssuerSubject(c OauthClaims, o *validateOptions) error {
	if o.requireIssuerSubject && c.Iss != c.Sub {
		return errors.New("issuer (iss) and subject (sub) do not match")
	}
	return nil
}

func validateExpiration(c OauthClaims, o *validateOptions) error {
	now := time.Now
	if o.timeNow != nil {
		now = o.timeNow
	}
	unixNow := now().Unix()

	if c.Exp != 0 && unixNow > c.Exp {
		return errors.New("token expired")
	}
	if c.Iat != 0 && unixNow < c.Iat {
		return errors.New("token used before issued")
	}
	return nil
}

func validateAudience(c OauthClaims, o *validateOptions) error {
	if o.expectedAudience != "" && c.Aud != o.expectedAudience {
		return fmt.Errorf("audience mismatch: want %q, got %q", o.expectedAudience, c.Aud)
	}
	if len(o.allowedAudiences) > 0 {
		if _, found := o.allowedAudiences[c.Aud]; !found {
			return fmt.Errorf("audience %q not allowed", c.Aud)
		}
	}
	return nil
}

func validateReplay(c OauthClaims, o *validateOptions) error {
	if o.replayChecker != nil {
		if o.replayChecker.Exists(c.Jti) {
			return errors.New("replay detected: jti already used")
		}

		if err := o.replayChecker.Set(c.Jti, time.Duration(c.Exp-time.Now().Unix())*time.Second); err != nil {
			return err
		}
	}

	return nil
}
