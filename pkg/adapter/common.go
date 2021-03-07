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
	ClientID string
	Adapter  Adapter
}

func (c *Config) Auth(ctx context.Context, callback func(url string, code string) error) (string, error) {
	cfg := &oauth2.Config{
		ClientID: c.ClientID,
		Endpoint: c.Adapter.Endpoint(),
		Scopes:   c.Adapter.Scopes(),
	}
	auth, err := cfg.AuthDevice(ctx)
	if err != nil {
		return "", err
	}
	err = callback(auth.VerificationURI, auth.UserCode)
	if err != nil {
		return "", err
	}
	token, err := cfg.Poll(ctx, auth)
	if err != nil {
		return "", err
	}
	return c.Adapter.UserID(ctx, cfg.Client(ctx, token))
}
