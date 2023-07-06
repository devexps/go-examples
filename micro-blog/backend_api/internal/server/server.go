package server

import (
	"context"
	v1 "github.com/devexps/go-examples/micro-blog/api/gen/go/backend_api/v1"
	pkgMid "github.com/devexps/go-examples/micro-blog/pkg/middleware"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware"
	authnEngine "github.com/devexps/go-micro/v2/middleware/authn/engine"
	authzEngine "github.com/devexps/go-micro/v2/middleware/authz/engine"
	"github.com/devexps/go-micro/v2/middleware/selector"
)

// NewWhiteListMatcher creates an authentication whitelist
func newWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]bool)
	whiteList[v1.OperationAuthenticationServiceLogin] = true
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewMiddlewares .
func NewMiddlewares(logger log.Logger,
	authenticator authnEngine.Authenticator,
	authorizer authzEngine.Authorizer,
) []middleware.Middleware {
	return pkgMid.NewMiddlewares(
		pkgMid.WithLogger(logger),
		pkgMid.WithAuthenticator(authenticator),
		pkgMid.WithAuthorizer(authorizer),
		pkgMid.WithMatcher(newWhiteListMatcher()),
	)
}
