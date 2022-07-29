// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sdkpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SDKClient is the client API for SDK service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SDKClient interface {
	Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error)
	WaitForTrigger(ctx context.Context, in *TriggerDescription, opts ...grpc.CallOption) (*TriggerResponse, error)
	ExecuteState(ctx context.Context, in *StateDescription, opts ...grpc.CallOption) (*StateOutput, error)
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error)
}

type sDKClient struct {
	cc grpc.ClientConnInterface
}

func NewSDKClient(cc grpc.ClientConnInterface) SDKClient {
	return &sDKClient{cc}
}

func (c *sDKClient) Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error) {
	out := new(CallResponse)
	err := c.cc.Invoke(ctx, "/sdkpb.SDK/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) WaitForTrigger(ctx context.Context, in *TriggerDescription, opts ...grpc.CallOption) (*TriggerResponse, error) {
	out := new(TriggerResponse)
	err := c.cc.Invoke(ctx, "/sdkpb.SDK/WaitForTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) ExecuteState(ctx context.Context, in *StateDescription, opts ...grpc.CallOption) (*StateOutput, error) {
	out := new(StateOutput)
	err := c.cc.Invoke(ctx, "/sdkpb.SDK/ExecuteState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error) {
	out := new(RunResponse)
	err := c.cc.Invoke(ctx, "/sdkpb.SDK/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SDKServer is the server API for SDK service.
// All implementations must embed UnimplementedSDKServer
// for forward compatibility
type SDKServer interface {
	Call(context.Context, *CallRequest) (*CallResponse, error)
	WaitForTrigger(context.Context, *TriggerDescription) (*TriggerResponse, error)
	ExecuteState(context.Context, *StateDescription) (*StateOutput, error)
	Run(context.Context, *RunRequest) (*RunResponse, error)
	mustEmbedUnimplementedSDKServer()
}

// UnimplementedSDKServer must be embedded to have forward compatible implementations.
type UnimplementedSDKServer struct {
}

func (UnimplementedSDKServer) Call(context.Context, *CallRequest) (*CallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedSDKServer) WaitForTrigger(context.Context, *TriggerDescription) (*TriggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitForTrigger not implemented")
}
func (UnimplementedSDKServer) ExecuteState(context.Context, *StateDescription) (*StateOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteState not implemented")
}
func (UnimplementedSDKServer) Run(context.Context, *RunRequest) (*RunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (UnimplementedSDKServer) mustEmbedUnimplementedSDKServer() {}

// UnsafeSDKServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SDKServer will
// result in compilation errors.
type UnsafeSDKServer interface {
	mustEmbedUnimplementedSDKServer()
}

func RegisterSDKServer(s grpc.ServiceRegistrar, srv SDKServer) {
	s.RegisterService(&SDK_ServiceDesc, srv)
}

func _SDK_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdkpb.SDK/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).Call(ctx, req.(*CallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_WaitForTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerDescription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).WaitForTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdkpb.SDK/WaitForTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).WaitForTrigger(ctx, req.(*TriggerDescription))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_ExecuteState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateDescription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).ExecuteState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdkpb.SDK/ExecuteState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).ExecuteState(ctx, req.(*StateDescription))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdkpb.SDK/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).Run(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SDK_ServiceDesc is the grpc.ServiceDesc for SDK service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SDK_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sdkpb.SDK",
	HandlerType: (*SDKServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _SDK_Call_Handler,
		},
		{
			MethodName: "WaitForTrigger",
			Handler:    _SDK_WaitForTrigger_Handler,
		},
		{
			MethodName: "ExecuteState",
			Handler:    _SDK_ExecuteState_Handler,
		},
		{
			MethodName: "Run",
			Handler:    _SDK_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sdkpb/sdk.proto",
}
