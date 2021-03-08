package oauth2

import (
	"context"
	"github.com/SCP-2000/wepam/pkg/adapter"
	"golang.org/x/oauth2"
)

type Challenge struct {
	DeviceAuth *oauth2.DeviceAuth
	oa         *oauth2.Config
	ad         adapter.Adapter
}

func NewChallenge(ctx context.Context, provider string, client_id string) (*Challenge, error) {
	ad, err := adapter.GetAdapter(provider)
	if err != nil {
		return nil, err
	}

	oa := &oauth2.Config{
		ClientID: client_id,
		Endpoint: ad.Endpoint(),
		Scopes:   ad.Scopes(),
	}

	auth, err := oa.AuthDevice(ctx)
	if err != nil {
		return nil, err
	}

	return &Challenge{
		DeviceAuth: auth,
		oa:         oa,
		ad:         ad,
	}, nil
}

func (c *Challenge) Resolve(ctx context.Context) ([]byte, error) {
	token, err := c.oa.Poll(ctx, c.DeviceAuth)
	if err != nil {
		return nil, err
	}
	return c.ad.Data(ctx, c.oa.Client(ctx, token))
}
