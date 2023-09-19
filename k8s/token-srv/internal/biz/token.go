package biz

import (
	"context"
	"github.com/devexps/go-examples/k8s/api/gen/go/common/authn"
	"github.com/devexps/go-examples/k8s/token-srv/internal/data"
	"github.com/devexps/go-micro/v2/log"
)

// TokenUseCase is a Token use case interface.
type TokenUseCase interface {
	GenerateToken(context.Context, *authn.Claims) (string, error)
}

// tokenUseCase is a Token use case.
type tokenUseCase struct {
	repo data.TokenRepo
	log  *log.Helper
}

// NewTokenUseCase new a Token use case.
func NewTokenUseCase(repo data.TokenRepo, logger log.Logger) TokenUseCase {
	l := log.NewHelper(log.With(logger, "module", "token_service/usecase/token"))
	return &tokenUseCase{repo: repo, log: l}
}

// GenerateToken .
func (u *tokenUseCase) GenerateToken(ctx context.Context, info *authn.Claims) (string, error) {
	return u.repo.GenerateToken(ctx, info)
}
