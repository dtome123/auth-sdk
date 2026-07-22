package client

import (
	"crypto/rsa"
	"errors"
	"time"

	"github.com/dtome123/auth-sdk/jwt"
	"github.com/dtome123/auth-sdk/pkgs/crypto"
)

type Option func(*impl) error

func WithHMACClientAssertion(secret string, assertionTTL time.Duration) Option {
	return func(c *impl) error {
		if secret == "" {
			return errors.New("client secret must not be nil for client assertion")
		}
		c.assertionTTL = assertionTTL
		c.clientServiceSigner = jwt.NewHMACSigner([]byte(secret))
		return nil
	}
}

func WithRS256ClientAssertion(privKey *rsa.PrivateKey, assertionTTL time.Duration) Option {
	return func(c *impl) error {
		if privKey == nil {
			return errors.New("client private key must not be nil for client assertion")
		}
		c.assertionTTL = assertionTTL
		c.clientServiceSigner = jwt.NewRS256SignerFromKey(privKey)
		return nil
	}
}

func WithRS256ClientAssertionFromPath(privKeyPath string, assertionTTL time.Duration) Option {
	return func(c *impl) error {
		if privKeyPath == "" {
			return errors.New("client private key path must not be empty")
		}
		key, err := crypto.LoadRSAPrivateKeyFromPath(privKeyPath)
		if err != nil {
			return err
		}
		c.assertionTTL = assertionTTL
		c.clientServiceSigner = jwt.NewRS256SignerFromKey(key)
		return nil
	}
}

func WithClientAssertionFromString(privKeyStr string, assertionTTL time.Duration) Option {
	return func(c *impl) error {
		if privKeyStr == "" {
			return errors.New("client private key string must not be empty")
		}
		key, err := crypto.LoadRSAPrivateKeyFromString(privKeyStr)
		if err != nil {
			return err
		}
		c.assertionTTL = assertionTTL
		c.clientServiceSigner = jwt.NewRS256SignerFromKey(key)
		return nil
	}
}
