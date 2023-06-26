package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/devexps/go-examples/micro-blog/api/gen/go/common/authn"
	"github.com/devexps/go-examples/micro-blog/api/gen/go/common/conf"
	v1 "github.com/devexps/go-examples/micro-blog/api/gen/go/token_service/v1"
	"github.com/devexps/go-micro/v2/log"
	"github.com/go-redis/redis/v8"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	ClaimSubject  = "sub"
	ClaimIssuedAt = "iat"
	ClaimUserId   = "uid"
	ClaimPlatform = "platform"

	userTokenKeyPrefix = "ut_"
)

// TokenRepo is a Token repo interface.
type TokenRepo interface {
	GenerateToken(context.Context, *authn.Claims) (string, error)
	ValidateToken(context.Context, string) (*authn.Claims, error)
	RemoveToken(context.Context, *authn.Claims, bool) error
}

// tokenRepo is a Token repository
type tokenRepo struct {
	log  *log.Helper
	data *Data
	key  []byte
}

// NewTokenRepo .
func NewTokenRepo(logger log.Logger, data *Data, cfg *conf.Bootstrap) TokenRepo {
	return &tokenRepo{
		log:  log.NewHelper(logger),
		data: data,
		key:  []byte(cfg.Auth.ApiKey),
	}
}

// GenerateToken .
func (r *tokenRepo) GenerateToken(ctx context.Context, info *authn.Claims) (string, error) {
	token := r.createJwtToken(ctx, info)
	if len(token) == 0 {
		return "", v1.ErrorCreateTokenFailed("create token failed")
	}
	if err := r.setToken(ctx, info); err != nil {
		r.log.Errorf("fail to generate token: %v", err)
		return "", v1.ErrorSaveTokenFailed("save token failed")
	}
	return token, nil
}

// ValidateToken .
func (r *tokenRepo) ValidateToken(ctx context.Context, token string) (*authn.Claims, error) {
	tokenInfo, err := r.parseAccessJwtTokenFromString(token)
	if err != nil {
		return nil, v1.ErrorInvalidToken("invalid token")
	}
	tokenInfo, err = r.getToken(ctx, tokenInfo)
	if err != nil {
		return nil, err
	}
	if tokenInfo == nil {
		return nil, v1.ErrorTokenNotExist("token not exist")
	}
	return tokenInfo, nil
}

// RemoveToken .
func (r *tokenRepo) RemoveToken(ctx context.Context, tokenInfo *authn.Claims, clearAll bool) error {
	if clearAll {
		return r.deleteAllToken(ctx, tokenInfo)
	}
	return r.deleteToken(ctx, tokenInfo)
}

func (r *tokenRepo) createJwtToken(ctx context.Context, info *authn.Claims) string {
	claims := jwtV4.NewWithClaims(jwtV4.SigningMethodHS256,
		jwtV4.MapClaims{
			ClaimSubject:  info.GetUserName(),
			ClaimUserId:   info.GetUserId(),
			ClaimPlatform: info.GetPlatform(),
			ClaimIssuedAt: time.Now().Unix(),
		},
	)
	signedToken, err := claims.SignedString(r.key)
	if err != nil {
		r.log.Errorf("fail to create JWT token: %v", err)
		return ""
	}
	return signedToken
}

func (r *tokenRepo) parseAccessJwtTokenFromString(token string) (*authn.Claims, error) {
	parseAuth, err := jwtV4.Parse(token, func(*jwtV4.Token) (interface{}, error) {
		return r.key, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := parseAuth.Claims.(jwtV4.MapClaims)
	if !ok {
		return nil, v1.ErrorNoTokenInContext("no jwt token in context")
	}
	tokenInfo, err := r.parseAccessJwtToken(claims)
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

func (r *tokenRepo) parseAccessJwtToken(claims jwtV4.Claims) (*authn.Claims, error) {
	if claims == nil {
		return nil, v1.ErrorClaimsIsNil("claims is nil")
	}
	mc, ok := claims.(jwtV4.MapClaims)
	if !ok {
		return nil, v1.ErrorClaimsIsNotMap("claims is not map claims")
	}
	var strUserName string
	if userName, ok := mc[ClaimSubject]; ok {
		strUserName = userName.(string)
	}
	var strUserId string
	if userId, ok := mc[ClaimUserId]; ok {
		strUserId = userId.(string)
	}
	var strPlatform string
	if platform, ok := mc[ClaimPlatform]; ok {
		strPlatform = platform.(string)
	}
	return &authn.Claims{
		Subject:  strUserName,
		UserId:   strUserId,
		UserName: strUserName,
		Platform: strPlatform,
	}, nil
}

func (r *tokenRepo) setToken(ctx context.Context, info *authn.Claims) error {
	data, err := r.stringFromTokenInfo(info)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("%s%s", userTokenKeyPrefix, info.GetUserId())
	field := info.GetPlatform()
	return r.data.rdb.HSet(ctx, key, field, data).Err()
}

func (r *tokenRepo) getToken(ctx context.Context, tokenInfo *authn.Claims) (*authn.Claims, error) {
	key := fmt.Sprintf("%s%s", userTokenKeyPrefix, tokenInfo.GetUserId())
	field := tokenInfo.GetPlatform()
	result, err := r.data.rdb.HGet(ctx, key, field).Result()
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("get redis user token failed: %s", err.Error())
		}
		return nil, nil
	}
	tokenInfo, err = r.tokenInfoFromString(result)
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

func (r *tokenRepo) deleteToken(ctx context.Context, tokenInfo *authn.Claims) error {
	key := fmt.Sprintf("%s%s", userTokenKeyPrefix, tokenInfo.GetUserId())
	field := tokenInfo.GetPlatform()
	return r.data.rdb.HDel(ctx, key, field).Err()
}

func (r *tokenRepo) deleteAllToken(ctx context.Context, tokenInfo *authn.Claims) error {
	key := fmt.Sprintf("%s%s", userTokenKeyPrefix, tokenInfo.GetUserId())
	return r.data.rdb.Del(ctx, key).Err()
}

func (r *tokenRepo) stringFromTokenInfo(info *authn.Claims) (string, error) {
	data, err := json.Marshal(info)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (r *tokenRepo) tokenInfoFromString(s string) (*authn.Claims, error) {
	info := &authn.Claims{}
	if err := json.Unmarshal([]byte(s), info); err != nil {
		return nil, err
	}
	return info, nil
}
