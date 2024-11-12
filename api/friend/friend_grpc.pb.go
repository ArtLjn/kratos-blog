// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: friend.proto

package friend

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

const (
	Friend_CreateFriend_FullMethodName = "/friend.Friend/CreateFriend"
	Friend_UpdateFriend_FullMethodName = "/friend.Friend/UpdateFriend"
	Friend_DeleteFriend_FullMethodName = "/friend.Friend/DeleteFriend"
	Friend_GetFriend_FullMethodName    = "/friend.Friend/GetFriend"
	Friend_ListFriend_FullMethodName   = "/friend.Friend/ListFriend"
)

// FriendClient is the client API for Friend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FriendClient interface {
	CreateFriend(ctx context.Context, in *CreateFriendRequest, opts ...grpc.CallOption) (*CreateFriendReply, error)
	UpdateFriend(ctx context.Context, in *UpdateFriendRequest, opts ...grpc.CallOption) (*UpdateFriendReply, error)
	DeleteFriend(ctx context.Context, in *DeleteFriendRequest, opts ...grpc.CallOption) (*DeleteFriendReply, error)
	GetFriend(ctx context.Context, in *GetFriendRequest, opts ...grpc.CallOption) (*GetFriendReply, error)
	ListFriend(ctx context.Context, in *ListFriendRequest, opts ...grpc.CallOption) (*ListFriendReply, error)
}

type friendClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendClient(cc grpc.ClientConnInterface) FriendClient {
	return &friendClient{cc}
}

func (c *friendClient) CreateFriend(ctx context.Context, in *CreateFriendRequest, opts ...grpc.CallOption) (*CreateFriendReply, error) {
	out := new(CreateFriendReply)
	err := c.cc.Invoke(ctx, Friend_CreateFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) UpdateFriend(ctx context.Context, in *UpdateFriendRequest, opts ...grpc.CallOption) (*UpdateFriendReply, error) {
	out := new(UpdateFriendReply)
	err := c.cc.Invoke(ctx, Friend_UpdateFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) DeleteFriend(ctx context.Context, in *DeleteFriendRequest, opts ...grpc.CallOption) (*DeleteFriendReply, error) {
	out := new(DeleteFriendReply)
	err := c.cc.Invoke(ctx, Friend_DeleteFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) GetFriend(ctx context.Context, in *GetFriendRequest, opts ...grpc.CallOption) (*GetFriendReply, error) {
	out := new(GetFriendReply)
	err := c.cc.Invoke(ctx, Friend_GetFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) ListFriend(ctx context.Context, in *ListFriendRequest, opts ...grpc.CallOption) (*ListFriendReply, error) {
	out := new(ListFriendReply)
	err := c.cc.Invoke(ctx, Friend_ListFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendServer is the server API for Friend service.
// All implementations must embed UnimplementedFriendServer
// for forward compatibility
type FriendServer interface {
	CreateFriend(context.Context, *CreateFriendRequest) (*CreateFriendReply, error)
	UpdateFriend(context.Context, *UpdateFriendRequest) (*UpdateFriendReply, error)
	DeleteFriend(context.Context, *DeleteFriendRequest) (*DeleteFriendReply, error)
	GetFriend(context.Context, *GetFriendRequest) (*GetFriendReply, error)
	ListFriend(context.Context, *ListFriendRequest) (*ListFriendReply, error)
	mustEmbedUnimplementedFriendServer()
}

// UnimplementedFriendServer must be embedded to have forward compatible implementations.
type UnimplementedFriendServer struct {
}

func (UnimplementedFriendServer) CreateFriend(context.Context, *CreateFriendRequest) (*CreateFriendReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFriend not implemented")
}
func (UnimplementedFriendServer) UpdateFriend(context.Context, *UpdateFriendRequest) (*UpdateFriendReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFriend not implemented")
}
func (UnimplementedFriendServer) DeleteFriend(context.Context, *DeleteFriendRequest) (*DeleteFriendReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFriend not implemented")
}
func (UnimplementedFriendServer) GetFriend(context.Context, *GetFriendRequest) (*GetFriendReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriend not implemented")
}
func (UnimplementedFriendServer) ListFriend(context.Context, *ListFriendRequest) (*ListFriendReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFriend not implemented")
}
func (UnimplementedFriendServer) mustEmbedUnimplementedFriendServer() {}

// UnsafeFriendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendServer will
// result in compilation errors.
type UnsafeFriendServer interface {
	mustEmbedUnimplementedFriendServer()
}

func RegisterFriendServer(s grpc.ServiceRegistrar, srv FriendServer) {
	s.RegisterService(&Friend_ServiceDesc, srv)
}

func _Friend_CreateFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).CreateFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_CreateFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).CreateFriend(ctx, req.(*CreateFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_UpdateFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).UpdateFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_UpdateFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).UpdateFriend(ctx, req.(*UpdateFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_DeleteFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).DeleteFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_DeleteFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).DeleteFriend(ctx, req.(*DeleteFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_GetFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).GetFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_GetFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).GetFriend(ctx, req.(*GetFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_ListFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).ListFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_ListFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).ListFriend(ctx, req.(*ListFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Friend_ServiceDesc is the grpc.ServiceDesc for Friend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Friend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "friend.Friend",
	HandlerType: (*FriendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFriend",
			Handler:    _Friend_CreateFriend_Handler,
		},
		{
			MethodName: "UpdateFriend",
			Handler:    _Friend_UpdateFriend_Handler,
		},
		{
			MethodName: "DeleteFriend",
			Handler:    _Friend_DeleteFriend_Handler,
		},
		{
			MethodName: "GetFriend",
			Handler:    _Friend_GetFriend_Handler,
		},
		{
			MethodName: "ListFriend",
			Handler:    _Friend_ListFriend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "friend.proto",
}