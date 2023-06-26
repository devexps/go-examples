package jwt

import (
	"context"
	"github.com/devexps/go-micro/v2/middleware/authn/engine"
)

type fn func(context.Context, string) (engine.Claims, error)

type Option func(*options)

type options struct {
	claimsFn fn
}

// WithClaimsFn .
func WithClaimsFn(f fn) Option {
	return func(o *options) {
		o.claimsFn = f
	}
}
