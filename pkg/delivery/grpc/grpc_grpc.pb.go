// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: grpc.proto

package grpcpb

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

// AntiBrutforceClient is the client API for AntiBrutforce service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AntiBrutforceClient interface {
	// Search returns a Google search result for the query.
	TryAuth(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
	ClearBucket(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
	AddToWhileList(ctx context.Context, in *SubNet, opts ...grpc.CallOption) (*Result, error)
	DelFromWhileList(ctx context.Context, in *SubNet, opts ...grpc.CallOption) (*Result, error)
	AddToBlackList(ctx context.Context, in *SubNet, opts ...grpc.CallOption) (*Result, error)
	DelFromBlackList(ctx context.Context, in *SubNet, opts ...grpc.CallOption) (*Result, error)
}

type antiBrutforceClient struct {
	cc grpc.ClientConnInterface
}

func NewAntiBrutforceClient(cc grpc.ClientConnInterface) AntiBrutforceClient {
	return &antiBrutforceClient{cc}
}

func (c *antiBrutforceClient) TryAuth(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ab_grpc.anti_brutforce/tryAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antiBrutforceClient) ClearBucket(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ab_ggrpcrpc.anti_brutforce/clearBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antiBrutforceClient) AddToWhileList(ctx context.Context, in *SubNet, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ab_grpc.anti_brutforce/AddToWhileList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antiBrutforceClient) DelFromWhileList(ctx context.Context, in *SubNet, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ab_grpc.anti_brutforce/DelFromWhileList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antiBrutforceClient) AddToBlackList(ctx context.Context, in *SubNet, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ab_grpc.anti_brutforce/AddToBlackList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antiBrutforceClient) DelFromBlackList(ctx context.Context, in *SubNet, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ab_grpc.anti_brutforce/DelFromBlackList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AntiBrutforceServer is the server API for AntiBrutforce service.
// All implementations must embed UnimplementedAntiBrutforceServer
// for forward compatibility
type AntiBrutforceServer interface {
	// Search returns a Google search result for the query.
	TryAuth(context.Context, *Request) (*Result, error)
	ClearBucket(context.Context, *Request) (*Result, error)
	AddToWhileList(context.Context, *SubNet) (*Result, error)
	DelFromWhileList(context.Context, *SubNet) (*Result, error)
	AddToBlackList(context.Context, *SubNet) (*Result, error)
	DelFromBlackList(context.Context, *SubNet) (*Result, error)
	mustEmbedUnimplementedAntiBrutforceServer()
}

// UnimplementedAntiBrutforceServer must be embedded to have forward compatible implementations.
type UnimplementedAntiBrutforceServer struct {
}

func (UnimplementedAntiBrutforceServer) TryAuth(context.Context, *Request) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TryAuth not implemented")
}
func (UnimplementedAntiBrutforceServer) ClearBucket(context.Context, *Request) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearBucket not implemented")
}
func (UnimplementedAntiBrutforceServer) AddToWhileList(context.Context, *SubNet) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToWhileList not implemented")
}
func (UnimplementedAntiBrutforceServer) DelFromWhileList(context.Context, *SubNet) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelFromWhileList not implemented")
}
func (UnimplementedAntiBrutforceServer) AddToBlackList(context.Context, *SubNet) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToBlackList not implemented")
}
func (UnimplementedAntiBrutforceServer) DelFromBlackList(context.Context, *SubNet) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelFromBlackList not implemented")
}
func (UnimplementedAntiBrutforceServer) mustEmbedUnimplementedAntiBrutforceServer() {}

// UnsafeAntiBrutforceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AntiBrutforceServer will
// result in compilation errors.
type UnsafeAntiBrutforceServer interface {
	mustEmbedUnimplementedAntiBrutforceServer()
}

func RegisterAntiBrutforceServer(s grpc.ServiceRegistrar, srv AntiBrutforceServer) {
	s.RegisterService(&AntiBrutforce_ServiceDesc, srv)
}

func _AntiBrutforce_TryAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBrutforceServer).TryAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ab_grpc.anti_brutforce/tryAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBrutforceServer).TryAuth(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntiBrutforce_ClearBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBrutforceServer).ClearBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ab_grpc.anti_brutforce/clearBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBrutforceServer).ClearBucket(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntiBrutforce_AddToWhileList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubNet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBrutforceServer).AddToWhileList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ab_grpc.anti_brutforce/AddToWhileList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBrutforceServer).AddToWhileList(ctx, req.(*SubNet))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntiBrutforce_DelFromWhileList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubNet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBrutforceServer).DelFromWhileList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ab_grpc.anti_brutforce/DelFromWhileList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBrutforceServer).DelFromWhileList(ctx, req.(*SubNet))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntiBrutforce_AddToBlackList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubNet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBrutforceServer).AddToBlackList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ab_grpc.anti_brutforce/AddToBlackList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBrutforceServer).AddToBlackList(ctx, req.(*SubNet))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntiBrutforce_DelFromBlackList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubNet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBrutforceServer).DelFromBlackList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ab_grpc.anti_brutforce/DelFromBlackList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBrutforceServer).DelFromBlackList(ctx, req.(*SubNet))
	}
	return interceptor(ctx, in, info, handler)
}

// AntiBrutforce_ServiceDesc is the grpc.ServiceDesc for AntiBrutforce service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AntiBrutforce_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ab_grpc.anti_brutforce",
	HandlerType: (*AntiBrutforceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "tryAuth",
			Handler:    _AntiBrutforce_TryAuth_Handler,
		},
		{
			MethodName: "clearBucket",
			Handler:    _AntiBrutforce_ClearBucket_Handler,
		},
		{
			MethodName: "AddToWhileList",
			Handler:    _AntiBrutforce_AddToWhileList_Handler,
		},
		{
			MethodName: "DelFromWhileList",
			Handler:    _AntiBrutforce_DelFromWhileList_Handler,
		},
		{
			MethodName: "AddToBlackList",
			Handler:    _AntiBrutforce_AddToBlackList_Handler,
		},
		{
			MethodName: "DelFromBlackList",
			Handler:    _AntiBrutforce_DelFromBlackList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}
