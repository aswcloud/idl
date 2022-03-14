// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package servercomm

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
	CreateRegisterToken(ctx context.Context, in *Void, opts ...grpc.CallOption) (*TokenMessage, error)
	CreateRefreshToken(ctx context.Context, in *CreateRefreshTokenMessage, opts ...grpc.CallOption) (*TokenMessage, error)
	// Needs, Authorization
	ReadRefreshToken(ctx context.Context, in *Void, opts ...grpc.CallOption) (*RefreshTokenList, error)
	// Needs, Authorization
	// Uuid : Refreshtoken Uuid
	UpdatehRefreshToken(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*TokenMessage, error)
	// Needs, Authorization
	// Uuid : Refreshtoken Uuid
	DeleteRefreshToken(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*DeleteRefreshTokenMessage, error)
	// Needs, Authorization
	// Uuid : Refreshtoken Uuid
	CreateAccessToken(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*TokenMessage, error)
}

type tokenClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenClient(cc grpc.ClientConnInterface) TokenClient {
	return &tokenClient{cc}
}

func (c *tokenClient) CreateRegisterToken(ctx context.Context, in *Void, opts ...grpc.CallOption) (*TokenMessage, error) {
	out := new(TokenMessage)
	err := c.cc.Invoke(ctx, "/v1.Token/CreateRegisterToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenClient) CreateRefreshToken(ctx context.Context, in *CreateRefreshTokenMessage, opts ...grpc.CallOption) (*TokenMessage, error) {
	out := new(TokenMessage)
	err := c.cc.Invoke(ctx, "/v1.Token/CreateRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenClient) ReadRefreshToken(ctx context.Context, in *Void, opts ...grpc.CallOption) (*RefreshTokenList, error) {
	out := new(RefreshTokenList)
	err := c.cc.Invoke(ctx, "/v1.Token/ReadRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenClient) UpdatehRefreshToken(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*TokenMessage, error) {
	out := new(TokenMessage)
	err := c.cc.Invoke(ctx, "/v1.Token/UpdatehRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenClient) DeleteRefreshToken(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*DeleteRefreshTokenMessage, error) {
	out := new(DeleteRefreshTokenMessage)
	err := c.cc.Invoke(ctx, "/v1.Token/DeleteRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenClient) CreateAccessToken(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*TokenMessage, error) {
	out := new(TokenMessage)
	err := c.cc.Invoke(ctx, "/v1.Token/CreateAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenServer is the server API for Token service.
// All implementations must embed UnimplementedTokenServer
// for forward compatibility
type TokenServer interface {
	CreateRegisterToken(context.Context, *Void) (*TokenMessage, error)
	CreateRefreshToken(context.Context, *CreateRefreshTokenMessage) (*TokenMessage, error)
	// Needs, Authorization
	ReadRefreshToken(context.Context, *Void) (*RefreshTokenList, error)
	// Needs, Authorization
	// Uuid : Refreshtoken Uuid
	UpdatehRefreshToken(context.Context, *Uuid) (*TokenMessage, error)
	// Needs, Authorization
	// Uuid : Refreshtoken Uuid
	DeleteRefreshToken(context.Context, *Uuid) (*DeleteRefreshTokenMessage, error)
	// Needs, Authorization
	// Uuid : Refreshtoken Uuid
	CreateAccessToken(context.Context, *Uuid) (*TokenMessage, error)
	mustEmbedUnimplementedTokenServer()
}

// UnimplementedTokenServer must be embedded to have forward compatible implementations.
type UnimplementedTokenServer struct {
}

func (UnimplementedTokenServer) CreateRegisterToken(context.Context, *Void) (*TokenMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRegisterToken not implemented")
}
func (UnimplementedTokenServer) CreateRefreshToken(context.Context, *CreateRefreshTokenMessage) (*TokenMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRefreshToken not implemented")
}
func (UnimplementedTokenServer) ReadRefreshToken(context.Context, *Void) (*RefreshTokenList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadRefreshToken not implemented")
}
func (UnimplementedTokenServer) UpdatehRefreshToken(context.Context, *Uuid) (*TokenMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatehRefreshToken not implemented")
}
func (UnimplementedTokenServer) DeleteRefreshToken(context.Context, *Uuid) (*DeleteRefreshTokenMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRefreshToken not implemented")
}
func (UnimplementedTokenServer) CreateAccessToken(context.Context, *Uuid) (*TokenMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccessToken not implemented")
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

func _Token_CreateRegisterToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServer).CreateRegisterToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Token/CreateRegisterToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServer).CreateRegisterToken(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Token_CreateRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRefreshTokenMessage)
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
		return srv.(TokenServer).CreateRefreshToken(ctx, req.(*CreateRefreshTokenMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Token_ReadRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServer).ReadRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Token/ReadRefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServer).ReadRefreshToken(ctx, req.(*Void))
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

func _Token_CreateAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Uuid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServer).CreateAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Token/CreateAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServer).CreateAccessToken(ctx, req.(*Uuid))
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
			MethodName: "CreateRegisterToken",
			Handler:    _Token_CreateRegisterToken_Handler,
		},
		{
			MethodName: "CreateRefreshToken",
			Handler:    _Token_CreateRefreshToken_Handler,
		},
		{
			MethodName: "ReadRefreshToken",
			Handler:    _Token_ReadRefreshToken_Handler,
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
			MethodName: "CreateAccessToken",
			Handler:    _Token_CreateAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loginToken.proto",
}
