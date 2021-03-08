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

func GetAdapter(provider string) (Adapter, error) {
	switch provider {
	case "github":
		return &githubAdapter{}, nil
	default:
		return nil, fmt.Errorf("unsupported provider: %s", provider)
	}
}
