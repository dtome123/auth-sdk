package oauthx

import (
	"sync"
)

type Registry struct {
	configs   Config
	providers map[string]Authenticator // stores initialized authenticators keyed by provider name
	mu        sync.Mutex               // ensures thread-safe access to maps
}

// Option defines a functional option for configuring the Registry
type Option func(r *Registry) error

func WithGoogleProvider(cfg GoogleConfig) Option {
	return func(r *Registry) error {
		if err := cfg.Validate(); err != nil {
			return err
		}
		r.configs.google = cfg

		provider, err := NewGoogle(cfg)
		if err != nil {
			return err
		}
		r.providers[string(Google)] = provider

		return nil
	}
}

func WithGitHubProvider(cfg GitHubConfig) Option {
	return func(r *Registry) error {
		if err := cfg.Validate(); err != nil {
			return err
		}
		r.configs.github = cfg

		provider, err := NewGitHub(cfg)
		if err != nil {
			return err
		}
		r.providers[string(GitHub)] = provider

		return nil
	}
}

func WithFacebookProvider(cfg FacebookConfig) Option {
	return func(r *Registry) error {
		if err := cfg.Validate(); err != nil {
			return err
		}
		r.configs.facebook = cfg

		provider, err := NewFacebook(cfg)
		if err != nil {
			return err
		}
		r.providers[string(Facebook)] = provider

		return nil
	}
}

func WithDiscordProvider(cfg DiscordConfig) Option {
	return func(r *Registry) error {
		if err := cfg.Validate(); err != nil {
			return err
		}
		r.configs.discord = cfg

		provider, err := NewDiscord(cfg)
		if err != nil {
			return err
		}
		r.providers[string(Discord)] = provider

		return nil
	}
}

func WithAppleProvider(cfg AppleConfig) Option {
	return func(r *Registry) error {
		if err := cfg.Validate(); err != nil {
			return err
		}
		r.configs.apple = cfg

		provider, err := NewApple(cfg)
		if err != nil {
			return err
		}
		r.providers[string(Apple)] = provider

		return nil
	}
}

// NewRegistry creates a new Registry with given options.
// Each option may register provider configs or customize Registry behavior.
func NewRegistry(opts ...Option) (*Registry, error) {
	r := &Registry{
		configs:   Config{},
		providers: make(map[string]Authenticator),
	}
	for _, opt := range opts {
		if err := opt(r); err != nil {
			return nil, err
		}
	}

	return r, nil
}

// GetGoogleProvider returns the Google Authenticator instance.
// If not yet initialized, it validates config and creates one.
func (r *Registry) GetGoogleProvider() (Authenticator, error) {

	if err := r.configs.google.Validate(); err != nil {
		return nil, err
	}

	return r.providers[string(Google)], nil
}

// GetGitHubProvider returns the GitHub Authenticator instance.
// If not yet initialized, it validates config and creates one.
func (r *Registry) GetGitHubProvider() (Authenticator, error) {

	if err := r.configs.google.Validate(); err != nil {
		return nil, err
	}

	return r.providers[string(GitHub)], nil
}

// GetFacebookProvider returns the Facebook Authenticator instance.
// If not yet initialized, it validates config and creates one.
func (r *Registry) GetFacebookProvider() (Authenticator, error) {

	if err := r.configs.facebook.Validate(); err != nil {
		return nil, err
	}

	return r.providers[string(Facebook)], nil
}

// GetDiscordProvider returns the Discord Authenticator instance.
// If not yet initialized, it validates config and creates one.
func (r *Registry) GetDiscordProvider() (Authenticator, error) {

	if err := r.configs.discord.Validate(); err != nil {
		return nil, err
	}

	return r.providers[string(Discord)], nil
}

// GetAppleProvider returns the Apple Authenticator instance.
// If not yet initialized, it validates config and creates one.
func (r *Registry) GetAppleProvider() (Authenticator, error) {

	if err := r.configs.apple.Validate(); err != nil {
		return nil, err
	}

	return r.providers[string(Apple)], nil
}
