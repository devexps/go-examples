package middleware

import (
	"context"
	casbinV1 "github.com/devexps/go-casbin/api/gen/go/casbin_service/v1"
	"github.com/devexps/go-micro/middleware/authz/engine/casbin/v2"
	"github.com/devexps/go-micro/v2/middleware/authz/engine"
)

type authorizer struct {
	client casbinV1.CasbinServiceClient
}

// IsAuthorized .
func (a *authorizer) IsAuthorized(ctx context.Context) error {
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

// NewAuthorizer .
func NewAuthorizer(client casbinV1.CasbinServiceClient) engine.Authorizer {
	return &authorizer{
		client: client,
	}
}

type authzClaimsMsg struct {
	subject, resource, action, project string
}

// NewAuthzClaims .
func NewAuthzClaims(subject, resource, action, project string) engine.Claims {
	return &authzClaimsMsg{
		subject:  subject,
		resource: resource,
		action:   action,
		project:  project,
	}
}

// GetSubject .
func (a *authzClaimsMsg) GetSubject() string {
	return a.subject
}

// GetAction .
func (a *authzClaimsMsg) GetAction() string {
	return a.action
}

// GetResource .
func (a *authzClaimsMsg) GetResource() string {
	return a.resource
}

// GetProject .
func (a *authzClaimsMsg) GetProject() string {
	return a.project
}
