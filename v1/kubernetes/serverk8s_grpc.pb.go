// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package kubernetes

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

// KubernetesClient is the client API for Kubernetes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KubernetesClient interface {
	CreateNamespace(ctx context.Context, in *Namespace, opts ...grpc.CallOption) (*Result, error)
	// rpc ReadNamespace(string) returns (Void);
	DeleteNamespace(ctx context.Context, in *Namespace, opts ...grpc.CallOption) (*Result, error)
	ListNamespace(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ListNamespace, error)
	CreateService(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Service, error)
	// rpc ReadService(Void) returns (Void);
	// rpc UpdateService(Void) returns (Void);
	DeleteService(ctx context.Context, in *DeleteService, opts ...grpc.CallOption) (*Result, error)
	ListService(ctx context.Context, in *Namespace, opts ...grpc.CallOption) (*ListService, error)
	CreateDeployment(ctx context.Context, in *Deployment, opts ...grpc.CallOption) (*Result, error)
	// rpc ReadDeployment(Void) returns (Void);
	// rpc UpdateDeployment(Void) returns (Void);
	DeleteDeployment(ctx context.Context, in *DeleteDeployment, opts ...grpc.CallOption) (*Result, error)
	ListDeployment(ctx context.Context, in *Namespace, opts ...grpc.CallOption) (*ListDeployment, error)
	CreatePersistentVolumeClaim(ctx context.Context, in *Pvc, opts ...grpc.CallOption) (*Result, error)
	// rpc ReadPersistentVolumeClaim(Void) returns (Void);
	// rpc UpdatePersistentVolumeClaim(Void) returns (Void);
	DeletePersistentVolumeClaim(ctx context.Context, in *DeletePvc, opts ...grpc.CallOption) (*Result, error)
	ListPersistentVolumeClaim(ctx context.Context, in *Namespace, opts ...grpc.CallOption) (*ListPvc, error)
}

type kubernetesClient struct {
	cc grpc.ClientConnInterface
}

func NewKubernetesClient(cc grpc.ClientConnInterface) KubernetesClient {
	return &kubernetesClient{cc}
}

func (c *kubernetesClient) CreateNamespace(ctx context.Context, in *Namespace, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/v1.Kubernetes/CreateNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) DeleteNamespace(ctx context.Context, in *Namespace, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/v1.Kubernetes/DeleteNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) ListNamespace(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ListNamespace, error) {
	out := new(ListNamespace)
	err := c.cc.Invoke(ctx, "/v1.Kubernetes/ListNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) CreateService(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, "/v1.Kubernetes/CreateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) DeleteService(ctx context.Context, in *DeleteService, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/v1.Kubernetes/DeleteService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) ListService(ctx context.Context, in *Namespace, opts ...grpc.CallOption) (*ListService, error) {
	out := new(ListService)
	err := c.cc.Invoke(ctx, "/v1.Kubernetes/ListService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) CreateDeployment(ctx context.Context, in *Deployment, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/v1.Kubernetes/CreateDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) DeleteDeployment(ctx context.Context, in *DeleteDeployment, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/v1.Kubernetes/DeleteDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) ListDeployment(ctx context.Context, in *Namespace, opts ...grpc.CallOption) (*ListDeployment, error) {
	out := new(ListDeployment)
	err := c.cc.Invoke(ctx, "/v1.Kubernetes/ListDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) CreatePersistentVolumeClaim(ctx context.Context, in *Pvc, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/v1.Kubernetes/CreatePersistentVolumeClaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) DeletePersistentVolumeClaim(ctx context.Context, in *DeletePvc, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/v1.Kubernetes/DeletePersistentVolumeClaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) ListPersistentVolumeClaim(ctx context.Context, in *Namespace, opts ...grpc.CallOption) (*ListPvc, error) {
	out := new(ListPvc)
	err := c.cc.Invoke(ctx, "/v1.Kubernetes/ListPersistentVolumeClaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubernetesServer is the server API for Kubernetes service.
// All implementations must embed UnimplementedKubernetesServer
// for forward compatibility
type KubernetesServer interface {
	CreateNamespace(context.Context, *Namespace) (*Result, error)
	// rpc ReadNamespace(string) returns (Void);
	DeleteNamespace(context.Context, *Namespace) (*Result, error)
	ListNamespace(context.Context, *Void) (*ListNamespace, error)
	CreateService(context.Context, *Service) (*Service, error)
	// rpc ReadService(Void) returns (Void);
	// rpc UpdateService(Void) returns (Void);
	DeleteService(context.Context, *DeleteService) (*Result, error)
	ListService(context.Context, *Namespace) (*ListService, error)
	CreateDeployment(context.Context, *Deployment) (*Result, error)
	// rpc ReadDeployment(Void) returns (Void);
	// rpc UpdateDeployment(Void) returns (Void);
	DeleteDeployment(context.Context, *DeleteDeployment) (*Result, error)
	ListDeployment(context.Context, *Namespace) (*ListDeployment, error)
	CreatePersistentVolumeClaim(context.Context, *Pvc) (*Result, error)
	// rpc ReadPersistentVolumeClaim(Void) returns (Void);
	// rpc UpdatePersistentVolumeClaim(Void) returns (Void);
	DeletePersistentVolumeClaim(context.Context, *DeletePvc) (*Result, error)
	ListPersistentVolumeClaim(context.Context, *Namespace) (*ListPvc, error)
	mustEmbedUnimplementedKubernetesServer()
}

// UnimplementedKubernetesServer must be embedded to have forward compatible implementations.
type UnimplementedKubernetesServer struct {
}

func (UnimplementedKubernetesServer) CreateNamespace(context.Context, *Namespace) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNamespace not implemented")
}
func (UnimplementedKubernetesServer) DeleteNamespace(context.Context, *Namespace) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNamespace not implemented")
}
func (UnimplementedKubernetesServer) ListNamespace(context.Context, *Void) (*ListNamespace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNamespace not implemented")
}
func (UnimplementedKubernetesServer) CreateService(context.Context, *Service) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}
func (UnimplementedKubernetesServer) DeleteService(context.Context, *DeleteService) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}
func (UnimplementedKubernetesServer) ListService(context.Context, *Namespace) (*ListService, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListService not implemented")
}
func (UnimplementedKubernetesServer) CreateDeployment(context.Context, *Deployment) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeployment not implemented")
}
func (UnimplementedKubernetesServer) DeleteDeployment(context.Context, *DeleteDeployment) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDeployment not implemented")
}
func (UnimplementedKubernetesServer) ListDeployment(context.Context, *Namespace) (*ListDeployment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDeployment not implemented")
}
func (UnimplementedKubernetesServer) CreatePersistentVolumeClaim(context.Context, *Pvc) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePersistentVolumeClaim not implemented")
}
func (UnimplementedKubernetesServer) DeletePersistentVolumeClaim(context.Context, *DeletePvc) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePersistentVolumeClaim not implemented")
}
func (UnimplementedKubernetesServer) ListPersistentVolumeClaim(context.Context, *Namespace) (*ListPvc, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPersistentVolumeClaim not implemented")
}
func (UnimplementedKubernetesServer) mustEmbedUnimplementedKubernetesServer() {}

// UnsafeKubernetesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KubernetesServer will
// result in compilation errors.
type UnsafeKubernetesServer interface {
	mustEmbedUnimplementedKubernetesServer()
}

func RegisterKubernetesServer(s grpc.ServiceRegistrar, srv KubernetesServer) {
	s.RegisterService(&Kubernetes_ServiceDesc, srv)
}

func _Kubernetes_CreateNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Namespace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).CreateNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Kubernetes/CreateNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).CreateNamespace(ctx, req.(*Namespace))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_DeleteNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Namespace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).DeleteNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Kubernetes/DeleteNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).DeleteNamespace(ctx, req.(*Namespace))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_ListNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).ListNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Kubernetes/ListNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).ListNamespace(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_CreateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).CreateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Kubernetes/CreateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).CreateService(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Kubernetes/DeleteService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).DeleteService(ctx, req.(*DeleteService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_ListService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Namespace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).ListService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Kubernetes/ListService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).ListService(ctx, req.(*Namespace))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_CreateDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Deployment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).CreateDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Kubernetes/CreateDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).CreateDeployment(ctx, req.(*Deployment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_DeleteDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeployment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).DeleteDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Kubernetes/DeleteDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).DeleteDeployment(ctx, req.(*DeleteDeployment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_ListDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Namespace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).ListDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Kubernetes/ListDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).ListDeployment(ctx, req.(*Namespace))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_CreatePersistentVolumeClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pvc)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).CreatePersistentVolumeClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Kubernetes/CreatePersistentVolumeClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).CreatePersistentVolumeClaim(ctx, req.(*Pvc))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_DeletePersistentVolumeClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePvc)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).DeletePersistentVolumeClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Kubernetes/DeletePersistentVolumeClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).DeletePersistentVolumeClaim(ctx, req.(*DeletePvc))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_ListPersistentVolumeClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Namespace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).ListPersistentVolumeClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Kubernetes/ListPersistentVolumeClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).ListPersistentVolumeClaim(ctx, req.(*Namespace))
	}
	return interceptor(ctx, in, info, handler)
}

// Kubernetes_ServiceDesc is the grpc.ServiceDesc for Kubernetes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Kubernetes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Kubernetes",
	HandlerType: (*KubernetesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNamespace",
			Handler:    _Kubernetes_CreateNamespace_Handler,
		},
		{
			MethodName: "DeleteNamespace",
			Handler:    _Kubernetes_DeleteNamespace_Handler,
		},
		{
			MethodName: "ListNamespace",
			Handler:    _Kubernetes_ListNamespace_Handler,
		},
		{
			MethodName: "CreateService",
			Handler:    _Kubernetes_CreateService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _Kubernetes_DeleteService_Handler,
		},
		{
			MethodName: "ListService",
			Handler:    _Kubernetes_ListService_Handler,
		},
		{
			MethodName: "CreateDeployment",
			Handler:    _Kubernetes_CreateDeployment_Handler,
		},
		{
			MethodName: "DeleteDeployment",
			Handler:    _Kubernetes_DeleteDeployment_Handler,
		},
		{
			MethodName: "ListDeployment",
			Handler:    _Kubernetes_ListDeployment_Handler,
		},
		{
			MethodName: "CreatePersistentVolumeClaim",
			Handler:    _Kubernetes_CreatePersistentVolumeClaim_Handler,
		},
		{
			MethodName: "DeletePersistentVolumeClaim",
			Handler:    _Kubernetes_DeletePersistentVolumeClaim_Handler,
		},
		{
			MethodName: "ListPersistentVolumeClaim",
			Handler:    _Kubernetes_ListPersistentVolumeClaim_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serverk8s.proto",
}
