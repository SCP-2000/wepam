package adapter

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
)

type Adapter interface {
	Endpoint() oauth2.Endpoint
	Scopes() []string
	Data(ctx context.Context, client *http.Client) ([]byte, error)
}

type Config struct {
	ClientID string
	Adapter  Adapter
}

func New(provider string, client_id string) (*Config, error) {
	switch provider {
	case "github":
		return &Config{
			ClientID: client_id,
			Adapter:  &Github{},
		}, nil
	default:
		return nil, fmt.Errorf("unsupported provider: %s", provider)
	}
}

func (c *Config) Auth(ctx context.Context, callback func(url string, code string) error) ([]byte, error) {
	cfg := &oauth2.Config{
		ClientID: c.ClientID,
		Endpoint: c.Adapter.Endpoint(),
		Scopes:   c.Adapter.Scopes(),
	}
	auth, err := cfg.AuthDevice(ctx)
	if err != nil {
		return nil, err
	}
	err = callback(auth.VerificationURI, auth.UserCode)
	if err != nil {
		return nil, err
	}
	token, err := cfg.Poll(ctx, auth)
	if err != nil {
		return nil, err
	}
	return c.Adapter.Data(ctx, cfg.Client(ctx, token))
}
