// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/front_user/front_user.proto

package front_user

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for FrontUser service

func NewFrontUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FrontUser service

type FrontUserService interface {
	FrontUserRegister(ctx context.Context, in *FrontUserRequest, opts ...client.CallOption) (*FrontUserResponse, error)
	FrontUserSendEmail(ctx context.Context, in *FrontUserMailRequest, opts ...client.CallOption) (*FrontUserResponse, error)
	FrontUserLogin(ctx context.Context, in *FrontUserRequest, opts ...client.CallOption) (*FrontUserResponse, error)
}

type frontUserService struct {
	c    client.Client
	name string
}

func NewFrontUserService(name string, c client.Client) FrontUserService {
	return &frontUserService{
		c:    c,
		name: name,
	}
}

func (c *frontUserService) FrontUserRegister(ctx context.Context, in *FrontUserRequest, opts ...client.CallOption) (*FrontUserResponse, error) {
	req := c.c.NewRequest(c.name, "FrontUser.FrontUserRegister", in)
	out := new(FrontUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontUserService) FrontUserSendEmail(ctx context.Context, in *FrontUserMailRequest, opts ...client.CallOption) (*FrontUserResponse, error) {
	req := c.c.NewRequest(c.name, "FrontUser.FrontUserSendEmail", in)
	out := new(FrontUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontUserService) FrontUserLogin(ctx context.Context, in *FrontUserRequest, opts ...client.CallOption) (*FrontUserResponse, error) {
	req := c.c.NewRequest(c.name, "FrontUser.FrontUserLogin", in)
	out := new(FrontUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FrontUser service

type FrontUserHandler interface {
	FrontUserRegister(context.Context, *FrontUserRequest, *FrontUserResponse) error
	FrontUserSendEmail(context.Context, *FrontUserMailRequest, *FrontUserResponse) error
	FrontUserLogin(context.Context, *FrontUserRequest, *FrontUserResponse) error
}

func RegisterFrontUserHandler(s server.Server, hdlr FrontUserHandler, opts ...server.HandlerOption) error {
	type frontUser interface {
		FrontUserRegister(ctx context.Context, in *FrontUserRequest, out *FrontUserResponse) error
		FrontUserSendEmail(ctx context.Context, in *FrontUserMailRequest, out *FrontUserResponse) error
		FrontUserLogin(ctx context.Context, in *FrontUserRequest, out *FrontUserResponse) error
	}
	type FrontUser struct {
		frontUser
	}
	h := &frontUserHandler{hdlr}
	return s.Handle(s.NewHandler(&FrontUser{h}, opts...))
}

type frontUserHandler struct {
	FrontUserHandler
}

func (h *frontUserHandler) FrontUserRegister(ctx context.Context, in *FrontUserRequest, out *FrontUserResponse) error {
	return h.FrontUserHandler.FrontUserRegister(ctx, in, out)
}

func (h *frontUserHandler) FrontUserSendEmail(ctx context.Context, in *FrontUserMailRequest, out *FrontUserResponse) error {
	return h.FrontUserHandler.FrontUserSendEmail(ctx, in, out)
}

func (h *frontUserHandler) FrontUserLogin(ctx context.Context, in *FrontUserRequest, out *FrontUserResponse) error {
	return h.FrontUserHandler.FrontUserLogin(ctx, in, out)
}
