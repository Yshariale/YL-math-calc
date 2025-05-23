// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: proto/task/task.proto

package task

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
	TaskManagementService_TaskGet_FullMethodName  = "/task.TaskManagementService/TaskGet"
	TaskManagementService_TaskPost_FullMethodName = "/task.TaskManagementService/TaskPost"
)

// TaskManagementServiceClient is the client API for TaskManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskManagementServiceClient interface {
	TaskGet(ctx context.Context, in *TaskGetRequest, opts ...grpc.CallOption) (*TaskGetResponse, error)
	TaskPost(ctx context.Context, in *TaskPostRequest, opts ...grpc.CallOption) (*TaskPostResponse, error)
}

type taskManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskManagementServiceClient(cc grpc.ClientConnInterface) TaskManagementServiceClient {
	return &taskManagementServiceClient{cc}
}

func (c *taskManagementServiceClient) TaskGet(ctx context.Context, in *TaskGetRequest, opts ...grpc.CallOption) (*TaskGetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskGetResponse)
	err := c.cc.Invoke(ctx, TaskManagementService_TaskGet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagementServiceClient) TaskPost(ctx context.Context, in *TaskPostRequest, opts ...grpc.CallOption) (*TaskPostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskPostResponse)
	err := c.cc.Invoke(ctx, TaskManagementService_TaskPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskManagementServiceServer is the server API for TaskManagementService service.
// All implementations must embed UnimplementedTaskManagementServiceServer
// for forward compatibility.
type TaskManagementServiceServer interface {
	TaskGet(context.Context, *TaskGetRequest) (*TaskGetResponse, error)
	TaskPost(context.Context, *TaskPostRequest) (*TaskPostResponse, error)
	mustEmbedUnimplementedTaskManagementServiceServer()
}

// UnimplementedTaskManagementServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTaskManagementServiceServer struct{}

func (UnimplementedTaskManagementServiceServer) TaskGet(context.Context, *TaskGetRequest) (*TaskGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskGet not implemented")
}
func (UnimplementedTaskManagementServiceServer) TaskPost(context.Context, *TaskPostRequest) (*TaskPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskPost not implemented")
}
func (UnimplementedTaskManagementServiceServer) mustEmbedUnimplementedTaskManagementServiceServer() {}
func (UnimplementedTaskManagementServiceServer) testEmbeddedByValue()                               {}

// UnsafeTaskManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskManagementServiceServer will
// result in compilation errors.
type UnsafeTaskManagementServiceServer interface {
	mustEmbedUnimplementedTaskManagementServiceServer()
}

func RegisterTaskManagementServiceServer(s grpc.ServiceRegistrar, srv TaskManagementServiceServer) {
	// If the following call pancis, it indicates UnimplementedTaskManagementServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TaskManagementService_ServiceDesc, srv)
}

func _TaskManagementService_TaskGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagementServiceServer).TaskGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskManagementService_TaskGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagementServiceServer).TaskGet(ctx, req.(*TaskGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManagementService_TaskPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagementServiceServer).TaskPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskManagementService_TaskPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagementServiceServer).TaskPost(ctx, req.(*TaskPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskManagementService_ServiceDesc is the grpc.ServiceDesc for TaskManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "task.TaskManagementService",
	HandlerType: (*TaskManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TaskGet",
			Handler:    _TaskManagementService_TaskGet_Handler,
		},
		{
			MethodName: "TaskPost",
			Handler:    _TaskManagementService_TaskPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/task/task.proto",
}
