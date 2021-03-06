// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/albato.proto

package albato

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Albato service

func NewAlbatoEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Albato service

type AlbatoService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Albato_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Albato_PingPongService, error)
	CustomUserRequest(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error)
}

type albatoService struct {
	c    client.Client
	name string
}

func NewAlbatoService(name string, c client.Client) AlbatoService {
	return &albatoService{
		c:    c,
		name: name,
	}
}

func (c *albatoService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Albato.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *albatoService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Albato_StreamService, error) {
	req := c.c.NewRequest(c.name, "Albato.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &albatoServiceStream{stream}, nil
}

type Albato_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type albatoServiceStream struct {
	stream client.Stream
}

func (x *albatoServiceStream) Close() error {
	return x.stream.Close()
}

func (x *albatoServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *albatoServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *albatoServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *albatoServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *albatoService) PingPong(ctx context.Context, opts ...client.CallOption) (Albato_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Albato.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &albatoServicePingPong{stream}, nil
}

type Albato_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type albatoServicePingPong struct {
	stream client.Stream
}

func (x *albatoServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *albatoServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *albatoServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *albatoServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *albatoServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *albatoServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *albatoService) CustomUserRequest(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "Albato.CustomUserRequest", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Albato service

type AlbatoHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Albato_StreamStream) error
	PingPong(context.Context, Albato_PingPongStream) error
	CustomUserRequest(context.Context, *UserRequest, *UserResponse) error
}

func RegisterAlbatoHandler(s server.Server, hdlr AlbatoHandler, opts ...server.HandlerOption) error {
	type albato interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
		CustomUserRequest(ctx context.Context, in *UserRequest, out *UserResponse) error
	}
	type Albato struct {
		albato
	}
	h := &albatoHandler{hdlr}
	return s.Handle(s.NewHandler(&Albato{h}, opts...))
}

type albatoHandler struct {
	AlbatoHandler
}

func (h *albatoHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.AlbatoHandler.Call(ctx, in, out)
}

func (h *albatoHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.AlbatoHandler.Stream(ctx, m, &albatoStreamStream{stream})
}

type Albato_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type albatoStreamStream struct {
	stream server.Stream
}

func (x *albatoStreamStream) Close() error {
	return x.stream.Close()
}

func (x *albatoStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *albatoStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *albatoStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *albatoStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *albatoHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.AlbatoHandler.PingPong(ctx, &albatoPingPongStream{stream})
}

type Albato_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type albatoPingPongStream struct {
	stream server.Stream
}

func (x *albatoPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *albatoPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *albatoPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *albatoPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *albatoPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *albatoPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *albatoHandler) CustomUserRequest(ctx context.Context, in *UserRequest, out *UserResponse) error {
	return h.AlbatoHandler.CustomUserRequest(ctx, in, out)
}
