// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: bbm/bomber/bomber.proto

package bomber

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

// Api Endpoints for Bomber service

func NewBomberEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Bomber service

type BomberService interface {
	BomberScriptDetail(ctx context.Context, in *BomberScriptDetailRequest, opts ...client.CallOption) (*BomberScriptDetailResponse, error)
}

type bomberService struct {
	c    client.Client
	name string
}

func NewBomberService(name string, c client.Client) BomberService {
	return &bomberService{
		c:    c,
		name: name,
	}
}

func (c *bomberService) BomberScriptDetail(ctx context.Context, in *BomberScriptDetailRequest, opts ...client.CallOption) (*BomberScriptDetailResponse, error) {
	req := c.c.NewRequest(c.name, "Bomber.BomberScriptDetail", in)
	out := new(BomberScriptDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Bomber service

type BomberHandler interface {
	BomberScriptDetail(context.Context, *BomberScriptDetailRequest, *BomberScriptDetailResponse) error
}

func RegisterBomberHandler(s server.Server, hdlr BomberHandler, opts ...server.HandlerOption) error {
	type bomber interface {
		BomberScriptDetail(ctx context.Context, in *BomberScriptDetailRequest, out *BomberScriptDetailResponse) error
	}
	type Bomber struct {
		bomber
	}
	h := &bomberHandler{hdlr}
	return s.Handle(s.NewHandler(&Bomber{h}, opts...))
}

type bomberHandler struct {
	BomberHandler
}

func (h *bomberHandler) BomberScriptDetail(ctx context.Context, in *BomberScriptDetailRequest, out *BomberScriptDetailResponse) error {
	return h.BomberHandler.BomberScriptDetail(ctx, in, out)
}
