package jwt

import (
	"errors"
)

// StaticTokenVerifier verifies a token string against an expected static secret or API key.
type StaticTokenVerifier struct {
	expectedToken string
}

// NewStaticTokenVerifier creates a new StaticTokenVerifier with the expected token.
func NewStaticTokenVerifier(expectedToken string) *StaticTokenVerifier {
	return &StaticTokenVerifier{expectedToken: expectedToken}
}

// Verify validates if the token string matches the expected static token.
func (s *StaticTokenVerifier) Verify(tokenStr string) (Claims, error) {
	if tokenStr == "" {
		return nil, errors.New("missing token")
	}
	if tokenStr != s.expectedToken {
		return nil, errors.New("invalid static token or api key")
	}
	return Claims{"sub": Value{Value: "static-key"}}, nil
}

// FuncVerifier adapts a custom verification function to the Verifier interface.
type FuncVerifier struct {
	verifyFn func(tokenStr string) (Claims, error)
}

// NewFuncVerifier creates a new FuncVerifier with a custom verification function.
func NewFuncVerifier(verifyFn func(tokenStr string) (Claims, error)) *FuncVerifier {
	return &FuncVerifier{verifyFn: verifyFn}
}

// Verify executes the custom verification function.
func (f *FuncVerifier) Verify(tokenStr string) (Claims, error) {
	if f.verifyFn == nil {
		return nil, errors.New("verify function is nil")
	}
	return f.verifyFn(tokenStr)
}
