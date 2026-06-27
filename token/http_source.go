package token

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dtome123/auth-sdk/jwt"
	"golang.org/x/oauth2"
)

// HttpTokenSourceConfig holds configuration parameters for HttpTokenSource.
type HttpTokenSourceConfig struct {
	AuthSvcEndpoint string
	Issuer          string
	Subject         string
	Audience        string
	Signer          jwt.Signer
	AssertionTTL    time.Duration
	Client          *http.Client
}

// Validate checks whether the config struct has valid required fields.
func (cfg *HttpTokenSourceConfig) Validate() error {
	if cfg.AuthSvcEndpoint == "" {
		return errors.New("authSvcEndpoint is required")
	}
	if cfg.Issuer == "" {
		return errors.New("issuer (iss) is required")
	}
	if cfg.Subject == "" {
		return errors.New("subject (sub) is required")
	}
	if cfg.Audience == "" {
		return errors.New("audience (aud) is required")
	}
	if cfg.Signer == nil {
		return errors.New("signer is required")
	}
	if cfg.AssertionTTL <= 0 {
		return errors.New("assertionTTL must be greater than zero")
	}
	return nil
}

// HttpTokenSource implements oauth2.TokenSource backed by HTTP POST request to AuthService.
type HttpTokenSource struct {
	authSvcEndpoint string
	signer          jwt.Signer
	iss             string
	sub             string
	aud             string
	assertionTTL    time.Duration
	client          *http.Client
}

// NewHttpTokenSource constructs a token source from HTTP AuthService using HttpTokenSourceConfig.
func NewHttpTokenSource(cfg HttpTokenSourceConfig) (*HttpTokenSource, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid http token source config: %w", err)
	}

	cli := cfg.Client
	if cli == nil {
		cli = &http.Client{Timeout: 10 * time.Second}
	}

	return &HttpTokenSource{
		authSvcEndpoint: cfg.AuthSvcEndpoint,
		signer:          cfg.Signer,
		iss:             cfg.Issuer,
		sub:             cfg.Subject,
		aud:             cfg.Audience,
		assertionTTL:    cfg.AssertionTTL,
		client:          cli,
	}, nil
}

// Close closes any idle connections on the HTTP client if applicable.
func (h *HttpTokenSource) Close() error {
	if h.client != nil {
		h.client.CloseIdleConnections()
	}
	return nil
}

// Token implements oauth2.TokenSource interface.
func (h *HttpTokenSource) Token() (*oauth2.Token, error) {
	assertion, err := jwt.GenerateClientAssertion(h.signer, h.iss, h.sub, h.aud, h.assertionTTL)
	if err != nil {
		return nil, err
	}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
	data.Set("client_assertion", assertion)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, h.authSvcEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http token request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("http token request returned status code: %d", resp.StatusCode)
	}

	var res struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int64  `json:"expires_in"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("failed to decode token response: %w", err)
	}

	return &oauth2.Token{
		AccessToken: res.AccessToken,
		TokenType:   res.TokenType,
		Expiry:      time.Now().Add(time.Duration(res.ExpiresIn) * time.Second),
	}, nil
}
