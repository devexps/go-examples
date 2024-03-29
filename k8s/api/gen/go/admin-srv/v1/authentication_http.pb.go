// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.2
// - protoc             v3.19.4
// source: admin-srv/v1/authentication.proto

package v1

import (
	context "context"
	http "github.com/devexps/go-micro/v2/transport/http"
	binding "github.com/devexps/go-micro/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the micro package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAuthenticationServiceLogin = "/admin_srv.v1.AuthenticationService/Login"

type AuthenticationServiceHTTPServer interface {
	// Login Login api
	Login(context.Context, *LoginRequest) (*LoginReply, error)
}

func RegisterAuthenticationServiceHTTPServer(s *http.Server, srv AuthenticationServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/admin/v1/login", _AuthenticationService_Login0_HTTP_Handler(srv))
}

func _AuthenticationService_Login0_HTTP_Handler(srv AuthenticationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthenticationServiceLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

type AuthenticationServiceHTTPClient interface {
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
}

type AuthenticationServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewAuthenticationServiceHTTPClient(client *http.Client) AuthenticationServiceHTTPClient {
	return &AuthenticationServiceHTTPClientImpl{client}
}

func (c *AuthenticationServiceHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/admin/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthenticationServiceLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
