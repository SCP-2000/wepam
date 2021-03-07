package adapter

import (
	"context"
	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
	"net/http"
)

type Github struct{}

func (g *Github) Endpoint() oauth2.Endpoint {
	return oauth2.Endpoint{
		AuthURL:       "https://github.com/login/oauth/authorize",
		DeviceAuthURL: "https://github.com/login/device/code",
		TokenURL:      "https://github.com/login/oauth/access_token",
	}
}

func (g *Github) Scopes() []string {
	return []string{"read:user"}
}

func (g *Github) UserID(ctx context.Context, client *http.Client) (string, error) {
	user, _, err := github.NewClient(client).Users.Get(ctx, "")
	if err != nil {
		return "", err
	}
	return user.GetLogin(), nil
}
