package adapter

import (
	"context"
	"golang.org/x/oauth2"
	"net/http"
)

type Adapter interface {
	Endpoint() oauth2.Endpoint
	Scopes() []string
	UserID(ctx context.Context, client *http.Client) (string, error)
}

type Config struct {
	ClientID     string
	ClientSecret string // required only on server side
	RedirectURL  string
	Adapter      Adapter
}

func (c *Config) UserID(ctx context.Context, code string) (string, error) {
	oAuth2Config := &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		Endpoint:     c.Adapter.Endpoint(),
		RedirectURL:  c.RedirectURL,
	}
	token, err := oAuth2Config.Exchange(ctx, code)
	if err != nil {
		return "", err
	}
	return c.Adapter.UserID(ctx, oAuth2Config.Client(ctx, token))
}

func (c *Config) AuthCodeURL(state string) string {
	oAuth2Config := &oauth2.Config{
		ClientID:    c.ClientID,
		Endpoint:    c.Adapter.Endpoint(),
		RedirectURL: c.RedirectURL,
		Scopes:      c.Adapter.Scopes(),
	}
	return oAuth2Config.AuthCodeURL(state)
}
