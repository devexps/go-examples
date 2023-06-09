package middleware

import (
	"github.com/devexps/go-micro/v2/log"
	authnEngine "github.com/devexps/go-micro/v2/middleware/authn/engine"
	authzEngine "github.com/devexps/go-micro/v2/middleware/authz/engine"
	"github.com/devexps/go-micro/v2/middleware/selector"
)

// Option is auth option.
type Option func(*options)

type options struct {
	logger  log.Logger
	matcher selector.MatchFunc

	authenticator authnEngine.Authenticator
	authorizer    authzEngine.Authorizer
}

// WithLogger .
func WithLogger(logger log.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

// WithMatcher .
func WithMatcher(matcher selector.MatchFunc) Option {
	return func(o *options) {
		o.matcher = matcher
	}
}

// WithAuthenticator .
func WithAuthenticator(authenticator authnEngine.Authenticator) Option {
	return func(o *options) {
		o.authenticator = authenticator
	}
}

// WithAuthorizer .
func WithAuthorizer(authorizer authzEngine.Authorizer) Option {
	return func(o *options) {
		o.authorizer = authorizer
	}
}
