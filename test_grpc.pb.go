// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// TokenClient is the client API for Token service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenClient interface {
	CreateRefreshToken(ctx context.Context, in *UserLoginMessage, opts ...grpc.CallOption) (*RefreshToken, error)
	// Header Check! Token!
	UpdatehRefreshToken(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*RefreshToken, error)
	DeleteRefreshToken(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*LoginTokenMessage, error)
	MakeAccessToken(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*AccessToken, error)
}

type tokenClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenClient(cc grpc.ClientConnInterface) TokenClient {
	return &tokenClient{cc}
}

func (c *tokenClient) CreateRefreshToken(ctx context.Context, in *UserLoginMessage, opts ...grpc.CallOption) (*RefreshToken, error) {
	out := new(RefreshToken)
	err := c.cc.Invoke(ctx, "/v1.Token/CreateRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenClient) UpdatehRefreshToken(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*RefreshToken, error) {
	out := new(RefreshToken)
	err := c.cc.Invoke(ctx, "/v1.Token/UpdatehRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenClient) DeleteRefreshToken(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*LoginTokenMessage, error) {
	out := new(LoginTokenMessage)
	err := c.cc.Invoke(ctx, "/v1.Token/DeleteRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenClient) MakeAccessToken(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*AccessToken, error) {
	out := new(AccessToken)
	err := c.cc.Invoke(ctx, "/v1.Token/MakeAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenServer is the server API for Token service.
// All implementations must embed UnimplementedTokenServer
// for forward compatibility
type TokenServer interface {
	CreateRefreshToken(context.Context, *UserLoginMessage) (*RefreshToken, error)
	// Header Check! Token!
	UpdatehRefreshToken(context.Context, *Uuid) (*RefreshToken, error)
	DeleteRefreshToken(context.Context, *Uuid) (*LoginTokenMessage, error)
	MakeAccessToken(context.Context, *Uuid) (*AccessToken, error)
	mustEmbedUnimplementedTokenServer()
}

// UnimplementedTokenServer must be embedded to have forward compatible implementations.
type UnimplementedTokenServer struct {
}

func (UnimplementedTokenServer) CreateRefreshToken(context.Context, *UserLoginMessage) (*RefreshToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRefreshToken not implemented")
}
func (UnimplementedTokenServer) UpdatehRefreshToken(context.Context, *Uuid) (*RefreshToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatehRefreshToken not implemented")
}
func (UnimplementedTokenServer) DeleteRefreshToken(context.Context, *Uuid) (*LoginTokenMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRefreshToken not implemented")
}
func (UnimplementedTokenServer) MakeAccessToken(context.Context, *Uuid) (*AccessToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeAccessToken not implemented")
}
func (UnimplementedTokenServer) mustEmbedUnimplementedTokenServer() {}

// UnsafeTokenServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenServer will
// result in compilation errors.
type UnsafeTokenServer interface {
	mustEmbedUnimplementedTokenServer()
}

func RegisterTokenServer(s grpc.ServiceRegistrar, srv TokenServer) {
	s.RegisterService(&Token_ServiceDesc, srv)
}

func _Token_CreateRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServer).CreateRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Token/CreateRefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServer).CreateRefreshToken(ctx, req.(*UserLoginMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Token_UpdatehRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Uuid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServer).UpdatehRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Token/UpdatehRefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServer).UpdatehRefreshToken(ctx, req.(*Uuid))
	}
	return interceptor(ctx, in, info, handler)
}

func _Token_DeleteRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Uuid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServer).DeleteRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Token/DeleteRefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServer).DeleteRefreshToken(ctx, req.(*Uuid))
	}
	return interceptor(ctx, in, info, handler)
}

func _Token_MakeAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Uuid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServer).MakeAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Token/MakeAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServer).MakeAccessToken(ctx, req.(*Uuid))
	}
	return interceptor(ctx, in, info, handler)
}

// Token_ServiceDesc is the grpc.ServiceDesc for Token service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Token_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Token",
	HandlerType: (*TokenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRefreshToken",
			Handler:    _Token_CreateRefreshToken_Handler,
		},
		{
			MethodName: "UpdatehRefreshToken",
			Handler:    _Token_UpdatehRefreshToken_Handler,
		},
		{
			MethodName: "DeleteRefreshToken",
			Handler:    _Token_DeleteRefreshToken_Handler,
		},
		{
			MethodName: "MakeAccessToken",
			Handler:    _Token_MakeAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
