// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: platform/user_srv_passport.proto

package passport

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/tinyhole/user-srv-passport/idl/base"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Passport service

type PassportService interface {
	// 注册
	SignUp(ctx context.Context, in *SignUpReq, opts ...client.CallOption) (*SignUpRsp, error)
	// 微信登陆，暂时只做这个登陆
	WeChatSignIn(ctx context.Context, in *WeChatSignInReq, opts ...client.CallOption) (*WeChatSignInRsp, error)
}

type passportService struct {
	c    client.Client
	name string
}

func NewPassportService(name string, c client.Client) PassportService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "platform.user.srv.passport"
	}
	return &passportService{
		c:    c,
		name: name,
	}
}

func (c *passportService) SignUp(ctx context.Context, in *SignUpReq, opts ...client.CallOption) (*SignUpRsp, error) {
	req := c.c.NewRequest(c.name, "Passport.SignUp", in)
	out := new(SignUpRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passportService) WeChatSignIn(ctx context.Context, in *WeChatSignInReq, opts ...client.CallOption) (*WeChatSignInRsp, error) {
	req := c.c.NewRequest(c.name, "Passport.WeChatSignIn", in)
	out := new(WeChatSignInRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Passport service

type PassportHandler interface {
	// 注册
	SignUp(context.Context, *SignUpReq, *SignUpRsp) error
	// 微信登陆，暂时只做这个登陆
	WeChatSignIn(context.Context, *WeChatSignInReq, *WeChatSignInRsp) error
}

func RegisterPassportHandler(s server.Server, hdlr PassportHandler, opts ...server.HandlerOption) error {
	type passport interface {
		SignUp(ctx context.Context, in *SignUpReq, out *SignUpRsp) error
		WeChatSignIn(ctx context.Context, in *WeChatSignInReq, out *WeChatSignInRsp) error
	}
	type Passport struct {
		passport
	}
	h := &passportHandler{hdlr}
	return s.Handle(s.NewHandler(&Passport{h}, opts...))
}

type passportHandler struct {
	PassportHandler
}

func (h *passportHandler) SignUp(ctx context.Context, in *SignUpReq, out *SignUpRsp) error {
	return h.PassportHandler.SignUp(ctx, in, out)
}

func (h *passportHandler) WeChatSignIn(ctx context.Context, in *WeChatSignInReq, out *WeChatSignInRsp) error {
	return h.PassportHandler.WeChatSignIn(ctx, in, out)
}