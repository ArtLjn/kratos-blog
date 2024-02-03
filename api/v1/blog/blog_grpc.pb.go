// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: v1/blog/blog.proto

package blog

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
	Blog_CreateBlog_FullMethodName             = "/api.v1.Blog/CreateBlog"
	Blog_UpdateBlog_FullMethodName             = "/api.v1.Blog/UpdateBlog"
	Blog_UpdateIndividualFields_FullMethodName = "/api.v1.Blog/UpdateIndividualFields"
	Blog_DeleteBlog_FullMethodName             = "/api.v1.Blog/DeleteBlog"
	Blog_GetBlogByTag_FullMethodName           = "/api.v1.Blog/GetBlogByTag"
	Blog_ListBlog_FullMethodName               = "/api.v1.Blog/ListBlog"
	Blog_GetBlogByID_FullMethodName            = "/api.v1.Blog/GetBlogByID"
	Blog_GetBlogByTitle_FullMethodName         = "/api.v1.Blog/GetBlogByTitle"
	Blog_UpdateOnly_FullMethodName             = "/api.v1.Blog/UpdateOnly"
	Blog_CacheBlog_FullMethodName              = "/api.v1.Blog/CacheBlog"
	Blog_GetCacheBlog_FullMethodName           = "/api.v1.Blog/GetCacheBlog"
)

// BlogClient is the client API for Blog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogClient interface {
	CreateBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogReply, error)
	UpdateBlog(ctx context.Context, in *UpdateBlogRequest, opts ...grpc.CallOption) (*UpdateBlogReply, error)
	UpdateIndividualFields(ctx context.Context, in *UpdateIndividualFieldsRequest, opts ...grpc.CallOption) (*UpdateIndividualFieldsReply, error)
	DeleteBlog(ctx context.Context, in *DeleteBlogRequest, opts ...grpc.CallOption) (*DeleteBlogReply, error)
	GetBlogByTag(ctx context.Context, in *GetBlogRequest, opts ...grpc.CallOption) (*GetBlogReply, error)
	ListBlog(ctx context.Context, in *ListBlogRequest, opts ...grpc.CallOption) (*ListBlogReply, error)
	GetBlogByID(ctx context.Context, in *GetBlogIDRequest, opts ...grpc.CallOption) (*GetBlogIDReply, error)
	GetBlogByTitle(ctx context.Context, in *GetBlogByTitleRequest, opts ...grpc.CallOption) (*GetBlogByTitleReply, error)
	UpdateOnly(ctx context.Context, in *UpdateOnlyRequest, opts ...grpc.CallOption) (*UpdateOnlyReply, error)
	CacheBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogReply, error)
	GetCacheBlog(ctx context.Context, in *ListBlogRequest, opts ...grpc.CallOption) (*ListCacheReply, error)
}

type blogClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogClient(cc grpc.ClientConnInterface) BlogClient {
	return &blogClient{cc}
}

func (c *blogClient) CreateBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogReply, error) {
	out := new(CreateBlogReply)
	err := c.cc.Invoke(ctx, Blog_CreateBlog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) UpdateBlog(ctx context.Context, in *UpdateBlogRequest, opts ...grpc.CallOption) (*UpdateBlogReply, error) {
	out := new(UpdateBlogReply)
	err := c.cc.Invoke(ctx, Blog_UpdateBlog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) UpdateIndividualFields(ctx context.Context, in *UpdateIndividualFieldsRequest, opts ...grpc.CallOption) (*UpdateIndividualFieldsReply, error) {
	out := new(UpdateIndividualFieldsReply)
	err := c.cc.Invoke(ctx, Blog_UpdateIndividualFields_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) DeleteBlog(ctx context.Context, in *DeleteBlogRequest, opts ...grpc.CallOption) (*DeleteBlogReply, error) {
	out := new(DeleteBlogReply)
	err := c.cc.Invoke(ctx, Blog_DeleteBlog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) GetBlogByTag(ctx context.Context, in *GetBlogRequest, opts ...grpc.CallOption) (*GetBlogReply, error) {
	out := new(GetBlogReply)
	err := c.cc.Invoke(ctx, Blog_GetBlogByTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) ListBlog(ctx context.Context, in *ListBlogRequest, opts ...grpc.CallOption) (*ListBlogReply, error) {
	out := new(ListBlogReply)
	err := c.cc.Invoke(ctx, Blog_ListBlog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) GetBlogByID(ctx context.Context, in *GetBlogIDRequest, opts ...grpc.CallOption) (*GetBlogIDReply, error) {
	out := new(GetBlogIDReply)
	err := c.cc.Invoke(ctx, Blog_GetBlogByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) GetBlogByTitle(ctx context.Context, in *GetBlogByTitleRequest, opts ...grpc.CallOption) (*GetBlogByTitleReply, error) {
	out := new(GetBlogByTitleReply)
	err := c.cc.Invoke(ctx, Blog_GetBlogByTitle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) UpdateOnly(ctx context.Context, in *UpdateOnlyRequest, opts ...grpc.CallOption) (*UpdateOnlyReply, error) {
	out := new(UpdateOnlyReply)
	err := c.cc.Invoke(ctx, Blog_UpdateOnly_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) CacheBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogReply, error) {
	out := new(CreateBlogReply)
	err := c.cc.Invoke(ctx, Blog_CacheBlog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) GetCacheBlog(ctx context.Context, in *ListBlogRequest, opts ...grpc.CallOption) (*ListCacheReply, error) {
	out := new(ListCacheReply)
	err := c.cc.Invoke(ctx, Blog_GetCacheBlog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServer is the server API for Blog service.
// All implementations must embed UnimplementedBlogServer
// for forward compatibility
type BlogServer interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*CreateBlogReply, error)
	UpdateBlog(context.Context, *UpdateBlogRequest) (*UpdateBlogReply, error)
	UpdateIndividualFields(context.Context, *UpdateIndividualFieldsRequest) (*UpdateIndividualFieldsReply, error)
	DeleteBlog(context.Context, *DeleteBlogRequest) (*DeleteBlogReply, error)
	GetBlogByTag(context.Context, *GetBlogRequest) (*GetBlogReply, error)
	ListBlog(context.Context, *ListBlogRequest) (*ListBlogReply, error)
	GetBlogByID(context.Context, *GetBlogIDRequest) (*GetBlogIDReply, error)
	GetBlogByTitle(context.Context, *GetBlogByTitleRequest) (*GetBlogByTitleReply, error)
	UpdateOnly(context.Context, *UpdateOnlyRequest) (*UpdateOnlyReply, error)
	CacheBlog(context.Context, *CreateBlogRequest) (*CreateBlogReply, error)
	GetCacheBlog(context.Context, *ListBlogRequest) (*ListCacheReply, error)
	mustEmbedUnimplementedBlogServer()
}

// UnimplementedBlogServer must be embedded to have forward compatible implementations.
type UnimplementedBlogServer struct {
}

func (UnimplementedBlogServer) CreateBlog(context.Context, *CreateBlogRequest) (*CreateBlogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlog not implemented")
}
func (UnimplementedBlogServer) UpdateBlog(context.Context, *UpdateBlogRequest) (*UpdateBlogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlog not implemented")
}
func (UnimplementedBlogServer) UpdateIndividualFields(context.Context, *UpdateIndividualFieldsRequest) (*UpdateIndividualFieldsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIndividualFields not implemented")
}
func (UnimplementedBlogServer) DeleteBlog(context.Context, *DeleteBlogRequest) (*DeleteBlogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlog not implemented")
}
func (UnimplementedBlogServer) GetBlogByTag(context.Context, *GetBlogRequest) (*GetBlogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlogByTag not implemented")
}
func (UnimplementedBlogServer) ListBlog(context.Context, *ListBlogRequest) (*ListBlogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBlog not implemented")
}
func (UnimplementedBlogServer) GetBlogByID(context.Context, *GetBlogIDRequest) (*GetBlogIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlogByID not implemented")
}
func (UnimplementedBlogServer) GetBlogByTitle(context.Context, *GetBlogByTitleRequest) (*GetBlogByTitleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlogByTitle not implemented")
}
func (UnimplementedBlogServer) UpdateOnly(context.Context, *UpdateOnlyRequest) (*UpdateOnlyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOnly not implemented")
}
func (UnimplementedBlogServer) CacheBlog(context.Context, *CreateBlogRequest) (*CreateBlogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CacheBlog not implemented")
}
func (UnimplementedBlogServer) GetCacheBlog(context.Context, *ListBlogRequest) (*ListCacheReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCacheBlog not implemented")
}
func (UnimplementedBlogServer) mustEmbedUnimplementedBlogServer() {}

// UnsafeBlogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogServer will
// result in compilation errors.
type UnsafeBlogServer interface {
	mustEmbedUnimplementedBlogServer()
}

func RegisterBlogServer(s grpc.ServiceRegistrar, srv BlogServer) {
	s.RegisterService(&Blog_ServiceDesc, srv)
}

func _Blog_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blog_CreateBlog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).CreateBlog(ctx, req.(*CreateBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_UpdateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).UpdateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blog_UpdateBlog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).UpdateBlog(ctx, req.(*UpdateBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_UpdateIndividualFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIndividualFieldsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).UpdateIndividualFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blog_UpdateIndividualFields_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).UpdateIndividualFields(ctx, req.(*UpdateIndividualFieldsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_DeleteBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).DeleteBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blog_DeleteBlog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).DeleteBlog(ctx, req.(*DeleteBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_GetBlogByTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).GetBlogByTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blog_GetBlogByTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).GetBlogByTag(ctx, req.(*GetBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_ListBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).ListBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blog_ListBlog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).ListBlog(ctx, req.(*ListBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_GetBlogByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlogIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).GetBlogByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blog_GetBlogByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).GetBlogByID(ctx, req.(*GetBlogIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_GetBlogByTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlogByTitleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).GetBlogByTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blog_GetBlogByTitle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).GetBlogByTitle(ctx, req.(*GetBlogByTitleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_UpdateOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).UpdateOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blog_UpdateOnly_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).UpdateOnly(ctx, req.(*UpdateOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_CacheBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).CacheBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blog_CacheBlog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).CacheBlog(ctx, req.(*CreateBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_GetCacheBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).GetCacheBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blog_GetCacheBlog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).GetCacheBlog(ctx, req.(*ListBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Blog_ServiceDesc is the grpc.ServiceDesc for Blog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Blog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.Blog",
	HandlerType: (*BlogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _Blog_CreateBlog_Handler,
		},
		{
			MethodName: "UpdateBlog",
			Handler:    _Blog_UpdateBlog_Handler,
		},
		{
			MethodName: "UpdateIndividualFields",
			Handler:    _Blog_UpdateIndividualFields_Handler,
		},
		{
			MethodName: "DeleteBlog",
			Handler:    _Blog_DeleteBlog_Handler,
		},
		{
			MethodName: "GetBlogByTag",
			Handler:    _Blog_GetBlogByTag_Handler,
		},
		{
			MethodName: "ListBlog",
			Handler:    _Blog_ListBlog_Handler,
		},
		{
			MethodName: "GetBlogByID",
			Handler:    _Blog_GetBlogByID_Handler,
		},
		{
			MethodName: "GetBlogByTitle",
			Handler:    _Blog_GetBlogByTitle_Handler,
		},
		{
			MethodName: "UpdateOnly",
			Handler:    _Blog_UpdateOnly_Handler,
		},
		{
			MethodName: "CacheBlog",
			Handler:    _Blog_CacheBlog_Handler,
		},
		{
			MethodName: "GetCacheBlog",
			Handler:    _Blog_GetCacheBlog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/blog/blog.proto",
}
