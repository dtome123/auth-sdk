package oauthx

import (
	"context"
	"encoding/json"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type githubAuth struct{ auth }

type GitHubConfig struct {
	RedirectURL  string
	ClientID     string
	ClientSecret string
}

func (c *GitHubConfig) Validate() error {
	if c.ClientID == "" || c.ClientSecret == "" || c.RedirectURL == "" {
		return fmt.Errorf("provider %s config invalid", GitHub)
	}
	return nil
}

func NewGitHub(cfg GitHubConfig) (Authenticator, error) {
	return &githubAuth{auth{
		oc: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			RedirectURL:  cfg.RedirectURL,
			Scopes:       []string{"read:user", "user:email"},
			Endpoint:     github.Endpoint,
		},
	}}, nil
}

func (g *githubAuth) Provider() Provider { return GitHub }
func (g *githubAuth) AuthCodeURLWithState(state, nonce string, challenge string) string {
	return g.authCodeURL(state, nonce, challenge)
}
func (g *githubAuth) ExchangeWithVerifier(ctx context.Context, code, verifier string) (*oauth2.Token, error) {
	return g.exchange(ctx, code, verifier)
}

func (g *githubAuth) FetchUser(ctx context.Context, tok *oauth2.Token) (User, error) {
	cli := oauth2.NewClient(ctx, oauth2.StaticTokenSource(tok))
	resp, err := cli.Get("https://api.github.com/user")
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()
	var u struct {
		ID                     int64
		Login, Name, AvatarURL string
	}
	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		return User{}, err
	}
	// emails
	eResp, err := cli.Get("https://api.github.com/user/emails")
	if err != nil {
		return User{}, err
	}
	defer eResp.Body.Close()
	var emails []struct {
		Email             string
		Primary, Verified bool
	}
	_ = json.NewDecoder(eResp.Body).Decode(&emails)
	email := ""
	var verified *bool
	for _, e := range emails {
		if e.Primary {
			email = e.Email
			v := e.Verified
			verified = &v
			break
		}
	}
	name := u.Name
	if name == "" {
		name = u.Login
	}
	return User{ID: fmt.Sprint(u.ID), Email: email, EmailVerified: verified, Name: name, AvatarURL: u.AvatarURL, Raw: u}, nil
}
