package middleware

import (
	"context"
	tokenV1 "github.com/devexps/go-examples/micro-blog/api/gen/go/token_service/v1"
	"github.com/devexps/go-examples/micro-blog/pkg/middleware/authn/jwt"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware/authn/engine"
)

type claimsUtil struct {
	client tokenV1.TokenServiceClient
}

// NewAuthenticator .
func NewAuthenticator(client tokenV1.TokenServiceClient) engine.Authenticator {
	c := &claimsUtil{
		client: client,
	}
	a, err := jwt.NewAuthenticator(
		jwt.WithClaimsFn(c.execute),
	)
	if err != nil {
		log.Error("new authenticator failed: ", err.Error())
	}
	return a
}

func (m *claimsUtil) execute(ctx context.Context, token string) (engine.Claims, error) {
	ret, err := m.client.ValidateToken(ctx, &tokenV1.ValidateTokenRequest{
		Token: token,
	})
	if err != nil {
		return nil, err
	}
	return ret.GetTokenInfo(), nil
}
