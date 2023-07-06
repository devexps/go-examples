package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2/model"
	fileAdapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/devexps/go-micro/middleware/authz/engine/casbin/v2"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware/authz/engine"
)

// NewAuthorizer .
func NewAuthorizer() engine.Authorizer {
	m, _ := model.NewModelFromFile("./configs/authz/authz_model.conf")
	p := fileAdapter.NewAdapter("./configs/authz/authz_policy.csv")
	a, err := casbin.NewAuthorizer(casbin.WithModel(m), casbin.WithPolicy(p))
	if err != nil {
		log.Fatal("new authorizer error: ", err)
	}
	return a
}

type authzClaimsMsg struct {
	subject, resource, action, project string
}

// NewAuthzClaims .
func NewAuthzClaims(subject, resource, action, project string) engine.Claims {
	fmt.Println(subject, resource, action, project)
	return &authzClaimsMsg{
		subject:  subject,
		resource: resource,
		action:   action,
		project:  project,
	}
}

func (a *authzClaimsMsg) GetSubject() string {
	return a.subject
}

func (a *authzClaimsMsg) GetAction() string {
	return a.action
}

func (a *authzClaimsMsg) GetResource() string {
	return a.resource
}

func (a *authzClaimsMsg) GetProject() string {
	return a.project
}
