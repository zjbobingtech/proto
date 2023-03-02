// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: bbm/basic/basic.proto

package basic

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

// Api Endpoints for Basic service

func NewBasicEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Basic service

type BasicService interface {
	BasicVerifyRbac(ctx context.Context, in *BasicVerifyRbacRequest, opts ...client.CallOption) (*BasicResponse, error)
	BasicGlobalParameters(ctx context.Context, in *BasicGlobalParametersRequest, opts ...client.CallOption) (*BasicGlobalParametersResponse, error)
	BasicBehaviorTrace(ctx context.Context, in *BasicBehaviorTraceRequest, opts ...client.CallOption) (*BasicResponse, error)
	BasicUserInfo(ctx context.Context, in *BasicUserInfoRequest, opts ...client.CallOption) (*BasicUserInfoResponse, error)
}

type basicService struct {
	c    client.Client
	name string
}

func NewBasicService(name string, c client.Client) BasicService {
	return &basicService{
		c:    c,
		name: name,
	}
}

func (c *basicService) BasicVerifyRbac(ctx context.Context, in *BasicVerifyRbacRequest, opts ...client.CallOption) (*BasicResponse, error) {
	req := c.c.NewRequest(c.name, "Basic.BasicVerifyRbac", in)
	out := new(BasicResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicService) BasicGlobalParameters(ctx context.Context, in *BasicGlobalParametersRequest, opts ...client.CallOption) (*BasicGlobalParametersResponse, error) {
	req := c.c.NewRequest(c.name, "Basic.BasicGlobalParameters", in)
	out := new(BasicGlobalParametersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicService) BasicBehaviorTrace(ctx context.Context, in *BasicBehaviorTraceRequest, opts ...client.CallOption) (*BasicResponse, error) {
	req := c.c.NewRequest(c.name, "Basic.BasicBehaviorTrace", in)
	out := new(BasicResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicService) BasicUserInfo(ctx context.Context, in *BasicUserInfoRequest, opts ...client.CallOption) (*BasicUserInfoResponse, error) {
	req := c.c.NewRequest(c.name, "Basic.BasicUserInfo", in)
	out := new(BasicUserInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Basic service

type BasicHandler interface {
	BasicVerifyRbac(context.Context, *BasicVerifyRbacRequest, *BasicResponse) error
	BasicGlobalParameters(context.Context, *BasicGlobalParametersRequest, *BasicGlobalParametersResponse) error
	BasicBehaviorTrace(context.Context, *BasicBehaviorTraceRequest, *BasicResponse) error
	BasicUserInfo(context.Context, *BasicUserInfoRequest, *BasicUserInfoResponse) error
}

func RegisterBasicHandler(s server.Server, hdlr BasicHandler, opts ...server.HandlerOption) error {
	type basic interface {
		BasicVerifyRbac(ctx context.Context, in *BasicVerifyRbacRequest, out *BasicResponse) error
		BasicGlobalParameters(ctx context.Context, in *BasicGlobalParametersRequest, out *BasicGlobalParametersResponse) error
		BasicBehaviorTrace(ctx context.Context, in *BasicBehaviorTraceRequest, out *BasicResponse) error
		BasicUserInfo(ctx context.Context, in *BasicUserInfoRequest, out *BasicUserInfoResponse) error
	}
	type Basic struct {
		basic
	}
	h := &basicHandler{hdlr}
	return s.Handle(s.NewHandler(&Basic{h}, opts...))
}

type basicHandler struct {
	BasicHandler
}

func (h *basicHandler) BasicVerifyRbac(ctx context.Context, in *BasicVerifyRbacRequest, out *BasicResponse) error {
	return h.BasicHandler.BasicVerifyRbac(ctx, in, out)
}

func (h *basicHandler) BasicGlobalParameters(ctx context.Context, in *BasicGlobalParametersRequest, out *BasicGlobalParametersResponse) error {
	return h.BasicHandler.BasicGlobalParameters(ctx, in, out)
}

func (h *basicHandler) BasicBehaviorTrace(ctx context.Context, in *BasicBehaviorTraceRequest, out *BasicResponse) error {
	return h.BasicHandler.BasicBehaviorTrace(ctx, in, out)
}

func (h *basicHandler) BasicUserInfo(ctx context.Context, in *BasicUserInfoRequest, out *BasicUserInfoResponse) error {
	return h.BasicHandler.BasicUserInfo(ctx, in, out)
}
