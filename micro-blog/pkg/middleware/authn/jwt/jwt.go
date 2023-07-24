package jwt

import (
	"context"
	"github.com/devexps/go-micro/v2/errors"
	"github.com/devexps/go-micro/v2/middleware/authn/engine"
)

type fn func(context.Context, string) (engine.Claims, error)

type authenticator struct {
	claimsFn fn
}

var _ engine.Authenticator = (*authenticator)(nil)

var (
	ErrMissingBearerToken = errors.Unauthorized(engine.Reason, "missing bearer token")
	ErrMissingClaimsFunc  = errors.Unauthorized(engine.Reason, "missing claims fun")
	ErrUnauthenticated    = errors.Unauthorized(engine.Reason, "unauthenticated")
)

// NewAuthenticator .
func NewAuthenticator(opts ...Option) engine.Authenticator {
	auth := &authenticator{}
	for _, o := range opts {
		o(auth)
	}
	return auth
}

// Option .
type Option func(*authenticator)

// WithClaimsFn .
func WithClaimsFn(f fn) Option {
	return func(o *authenticator) {
		o.claimsFn = f
	}
}

// Authenticate .
func (a *authenticator) Authenticate(ctx context.Context, ctxType engine.ContextType) (engine.Claims, error) {
	tokenString, err := engine.AuthFromMD(ctx, ctxType)
	if err != nil {
		return nil, ErrMissingBearerToken
	}
	if a.claimsFn == nil {
		return nil, ErrMissingClaimsFunc
	}
	claims, err := a.claimsFn(ctx, tokenString)
	if err != nil {
		return nil, ErrUnauthenticated
	}
	return claims, nil
}

// CreateIdentity .
func (a *authenticator) CreateIdentity(ctx context.Context, ctxType engine.ContextType, claims engine.Claims) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}
