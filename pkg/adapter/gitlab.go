package adapter

import (
	"context"
	"fmt"
	"github.com/valyala/fastjson"
	"golang.org/x/oauth2"
	oauth2_gitlab "golang.org/x/oauth2/gitlab"
	"io"
	"net/http"
)

type Gitlab struct{}

func (g *Gitlab) Endpoint() oauth2.Endpoint {
	return oauth2_gitlab.Endpoint
}

func (g *Gitlab) Scopes() []string {
	return []string{"read_user"}
}

func (g *Gitlab) UserID(ctx context.Context, client *http.Client) (string, error) {
	// TODO: require some rework
	resp, err := client.Get("https://gitlab.com/api/v4/user")
	if err != nil {
		return "", err
	}
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	decoded, err := fastjson.ParseBytes(payload)
	if err != nil {
		return "", err
	}
	username := decoded.GetStringBytes("username")
	if username == nil {
		return "", fmt.Errorf("no username found in resp")
	}
	return string(username), nil
}
