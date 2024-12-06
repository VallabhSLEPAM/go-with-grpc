// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: hello/hello.proto

package hello

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	HelloService_SayHello_FullMethodName          = "/hello.HelloService/SayHello"
	HelloService_HelloServerStream_FullMethodName = "/hello.HelloService/HelloServerStream"
	HelloService_HelloClientStream_FullMethodName = "/hello.HelloService/HelloClientStream"
	HelloService_HelloContinuous_FullMethodName   = "/hello.HelloService/HelloContinuous"
)

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
	// Unary Server
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	// Server stream: Send greetings 10 times with 1 second interval
	HelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[HelloResponse], error)
	// Client stream: Send 10 request and gets one whole response
	HelloClientStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[HelloRequest, HelloResponse], error)
	// Bidirectional stream: Send 10 request and gets responsesfor each
	HelloContinuous(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[HelloRequest, HelloResponse], error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, HelloService_SayHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) HelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[HelloResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[0], HelloService_HelloServerStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[HelloRequest, HelloResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type HelloService_HelloServerStreamClient = grpc.ServerStreamingClient[HelloResponse]

func (c *helloServiceClient) HelloClientStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[HelloRequest, HelloResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[1], HelloService_HelloClientStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[HelloRequest, HelloResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type HelloService_HelloClientStreamClient = grpc.ClientStreamingClient[HelloRequest, HelloResponse]

func (c *helloServiceClient) HelloContinuous(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[HelloRequest, HelloResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[2], HelloService_HelloContinuous_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[HelloRequest, HelloResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type HelloService_HelloContinuousClient = grpc.BidiStreamingClient[HelloRequest, HelloResponse]

// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility.
type HelloServiceServer interface {
	// Unary Server
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	// Server stream: Send greetings 10 times with 1 second interval
	HelloServerStream(*HelloRequest, grpc.ServerStreamingServer[HelloResponse]) error
	// Client stream: Send 10 request and gets one whole response
	HelloClientStream(grpc.ClientStreamingServer[HelloRequest, HelloResponse]) error
	// Bidirectional stream: Send 10 request and gets responsesfor each
	HelloContinuous(grpc.BidiStreamingServer[HelloRequest, HelloResponse]) error
	mustEmbedUnimplementedHelloServiceServer()
}

// UnimplementedHelloServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHelloServiceServer struct{}

func (UnimplementedHelloServiceServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedHelloServiceServer) HelloServerStream(*HelloRequest, grpc.ServerStreamingServer[HelloResponse]) error {
	return status.Errorf(codes.Unimplemented, "method HelloServerStream not implemented")
}
func (UnimplementedHelloServiceServer) HelloClientStream(grpc.ClientStreamingServer[HelloRequest, HelloResponse]) error {
	return status.Errorf(codes.Unimplemented, "method HelloClientStream not implemented")
}
func (UnimplementedHelloServiceServer) HelloContinuous(grpc.BidiStreamingServer[HelloRequest, HelloResponse]) error {
	return status.Errorf(codes.Unimplemented, "method HelloContinuous not implemented")
}
func (UnimplementedHelloServiceServer) mustEmbedUnimplementedHelloServiceServer() {}
func (UnimplementedHelloServiceServer) testEmbeddedByValue()                      {}

// UnsafeHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServiceServer will
// result in compilation errors.
type UnsafeHelloServiceServer interface {
	mustEmbedUnimplementedHelloServiceServer()
}

func RegisterHelloServiceServer(s grpc.ServiceRegistrar, srv HelloServiceServer) {
	// If the following call pancis, it indicates UnimplementedHelloServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HelloService_ServiceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloService_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_HelloServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServiceServer).HelloServerStream(m, &grpc.GenericServerStream[HelloRequest, HelloResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type HelloService_HelloServerStreamServer = grpc.ServerStreamingServer[HelloResponse]

func _HelloService_HelloClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).HelloClientStream(&grpc.GenericServerStream[HelloRequest, HelloResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type HelloService_HelloClientStreamServer = grpc.ClientStreamingServer[HelloRequest, HelloResponse]

func _HelloService_HelloContinuous_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).HelloContinuous(&grpc.GenericServerStream[HelloRequest, HelloResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type HelloService_HelloContinuousServer = grpc.BidiStreamingServer[HelloRequest, HelloResponse]

// HelloService_ServiceDesc is the grpc.ServiceDesc for HelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HelloServerStream",
			Handler:       _HelloService_HelloServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HelloClientStream",
			Handler:       _HelloService_HelloClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "HelloContinuous",
			Handler:       _HelloService_HelloContinuous_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hello/hello.proto",
}
