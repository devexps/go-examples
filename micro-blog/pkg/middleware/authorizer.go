package middleware

import (
	casbinV1 "github.com/devexps/go-casbin/api/gen/go/casbin_service/v1"
	"github.com/devexps/go-examples/micro-blog/pkg/middleware/authz/casbin"
	"github.com/devexps/go-micro/v2/middleware/authz/engine"
)

// NewAuthorizer .
func NewAuthorizer(client casbinV1.CasbinServiceClient) engine.Authorizer {
	return casbin.NewAuthorizer(
		casbin.WithClient(client),
	)
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
