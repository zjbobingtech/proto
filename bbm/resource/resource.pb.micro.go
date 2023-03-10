// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: bbm/resource/resource.proto

package resource

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

// Api Endpoints for Resource service

func NewResourceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Resource service

type ResourceService interface {
	ResourceDetail(ctx context.Context, in *ResourceDetailRequest, opts ...client.CallOption) (*ResourceDetailResponse, error)
}

type resourceService struct {
	c    client.Client
	name string
}

func NewResourceService(name string, c client.Client) ResourceService {
	return &resourceService{
		c:    c,
		name: name,
	}
}

func (c *resourceService) ResourceDetail(ctx context.Context, in *ResourceDetailRequest, opts ...client.CallOption) (*ResourceDetailResponse, error) {
	req := c.c.NewRequest(c.name, "Resource.ResourceDetail", in)
	out := new(ResourceDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Resource service

type ResourceHandler interface {
	ResourceDetail(context.Context, *ResourceDetailRequest, *ResourceDetailResponse) error
}

func RegisterResourceHandler(s server.Server, hdlr ResourceHandler, opts ...server.HandlerOption) error {
	type resource interface {
		ResourceDetail(ctx context.Context, in *ResourceDetailRequest, out *ResourceDetailResponse) error
	}
	type Resource struct {
		resource
	}
	h := &resourceHandler{hdlr}
	return s.Handle(s.NewHandler(&Resource{h}, opts...))
}

type resourceHandler struct {
	ResourceHandler
}

func (h *resourceHandler) ResourceDetail(ctx context.Context, in *ResourceDetailRequest, out *ResourceDetailResponse) error {
	return h.ResourceHandler.ResourceDetail(ctx, in, out)
}
