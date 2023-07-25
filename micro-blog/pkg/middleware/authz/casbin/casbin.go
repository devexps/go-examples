package casbin

import (
	"context"
	casbinV1 "github.com/devexps/go-casbin/api/gen/go/casbin_service/v1"
	"github.com/devexps/go-micro/middleware/authz/engine/casbin/v2"
	"github.com/devexps/go-micro/v2/errors"
	"github.com/devexps/go-micro/v2/middleware/authz/engine"
)

type authorizer struct {
	client casbinV1.CasbinServiceClient
}

var _ engine.Authorizer = (*authorizer)(nil)

var (
	ErrMissingClient = errors.Unauthorized(engine.Reason, "missing client")
)

// NewAuthorizer .
func NewAuthorizer(opts ...Option) engine.Authorizer {
	auth := &authorizer{}
	for _, o := range opts {
		o(auth)
	}
	return auth
}

// Option .
type Option func(*authorizer)

// WithClient .
func WithClient(client casbinV1.CasbinServiceClient) Option {
	return func(a *authorizer) {
		a.client = client
	}
}

// IsAuthorized .
func (a *authorizer) IsAuthorized(ctx context.Context) error {
	if a.client == nil {
		return ErrMissingClient
	}
	claims, ok := engine.FromContext(ctx)
	if !ok {
		return casbin.ErrMissingClaims
	}
	if len(claims.GetSubject()) == 0 || len(claims.GetResource()) == 0 || len(claims.GetAction()) == 0 {
		return casbin.ErrInvalidClaims
	}
	var (
		allowed *casbinV1.BoolReply
		err     error
	)
	if len(claims.GetProject()) > 0 {
		allowed, err = a.client.Enforce(ctx, &casbinV1.EnforceRequest{
			Params: []string{claims.GetSubject(), claims.GetResource(), claims.GetAction(), claims.GetProject()},
		})
	} else {
		allowed, err = a.client.Enforce(ctx, &casbinV1.EnforceRequest{
			Params: []string{claims.GetSubject(), claims.GetResource(), claims.GetAction()},
		})
	}
	if err != nil {
		return casbin.ErrUnauthorized
	}
	if !allowed.Res {
		return casbin.ErrUnauthorized
	}
	return nil
}
