// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: adminAlert.proto

package notification

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	AdminAlertService_CreateAdminAlert_FullMethodName = "/ecommerce.AdminAlertService/CreateAdminAlert"
	AdminAlertService_GetAdminAlert_FullMethodName    = "/ecommerce.AdminAlertService/GetAdminAlert"
	AdminAlertService_DeleteAdminAlert_FullMethodName = "/ecommerce.AdminAlertService/DeleteAdminAlert"
	AdminAlertService_ListAdminAlerts_FullMethodName  = "/ecommerce.AdminAlertService/ListAdminAlerts"
)

// AdminAlertServiceClient is the client API for AdminAlertService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage admin alerts
type AdminAlertServiceClient interface {
	// Admin ogohlantirishini yaratish
	CreateAdminAlert(ctx context.Context, in *CreateAdminAlertRequest, opts ...grpc.CallOption) (*AdminAlertResponse, error)
	// Admin ogohlantirishini olish
	GetAdminAlert(ctx context.Context, in *AdminAlertRequest, opts ...grpc.CallOption) (*AdminAlertResponse, error)
	// Admin ogohlantirishini o'chirish
	DeleteAdminAlert(ctx context.Context, in *AdminAlertRequest, opts ...grpc.CallOption) (*AdminAlertResponse, error)
	// Barcha admin ogohlantirishlarini olish
	ListAdminAlerts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AdminAlertListResponse, error)
}

type adminAlertServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminAlertServiceClient(cc grpc.ClientConnInterface) AdminAlertServiceClient {
	return &adminAlertServiceClient{cc}
}

func (c *adminAlertServiceClient) CreateAdminAlert(ctx context.Context, in *CreateAdminAlertRequest, opts ...grpc.CallOption) (*AdminAlertResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminAlertResponse)
	err := c.cc.Invoke(ctx, AdminAlertService_CreateAdminAlert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAlertServiceClient) GetAdminAlert(ctx context.Context, in *AdminAlertRequest, opts ...grpc.CallOption) (*AdminAlertResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminAlertResponse)
	err := c.cc.Invoke(ctx, AdminAlertService_GetAdminAlert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAlertServiceClient) DeleteAdminAlert(ctx context.Context, in *AdminAlertRequest, opts ...grpc.CallOption) (*AdminAlertResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminAlertResponse)
	err := c.cc.Invoke(ctx, AdminAlertService_DeleteAdminAlert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAlertServiceClient) ListAdminAlerts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AdminAlertListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminAlertListResponse)
	err := c.cc.Invoke(ctx, AdminAlertService_ListAdminAlerts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminAlertServiceServer is the server API for AdminAlertService service.
// All implementations must embed UnimplementedAdminAlertServiceServer
// for forward compatibility
//
// Service to manage admin alerts
type AdminAlertServiceServer interface {
	// Admin ogohlantirishini yaratish
	CreateAdminAlert(context.Context, *CreateAdminAlertRequest) (*AdminAlertResponse, error)
	// Admin ogohlantirishini olish
	GetAdminAlert(context.Context, *AdminAlertRequest) (*AdminAlertResponse, error)
	// Admin ogohlantirishini o'chirish
	DeleteAdminAlert(context.Context, *AdminAlertRequest) (*AdminAlertResponse, error)
	// Barcha admin ogohlantirishlarini olish
	ListAdminAlerts(context.Context, *Empty) (*AdminAlertListResponse, error)
	mustEmbedUnimplementedAdminAlertServiceServer()
}

// UnimplementedAdminAlertServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminAlertServiceServer struct {
}

func (UnimplementedAdminAlertServiceServer) CreateAdminAlert(context.Context, *CreateAdminAlertRequest) (*AdminAlertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAdminAlert not implemented")
}
func (UnimplementedAdminAlertServiceServer) GetAdminAlert(context.Context, *AdminAlertRequest) (*AdminAlertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdminAlert not implemented")
}
func (UnimplementedAdminAlertServiceServer) DeleteAdminAlert(context.Context, *AdminAlertRequest) (*AdminAlertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAdminAlert not implemented")
}
func (UnimplementedAdminAlertServiceServer) ListAdminAlerts(context.Context, *Empty) (*AdminAlertListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAdminAlerts not implemented")
}
func (UnimplementedAdminAlertServiceServer) mustEmbedUnimplementedAdminAlertServiceServer() {}

// UnsafeAdminAlertServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminAlertServiceServer will
// result in compilation errors.
type UnsafeAdminAlertServiceServer interface {
	mustEmbedUnimplementedAdminAlertServiceServer()
}

func RegisterAdminAlertServiceServer(s grpc.ServiceRegistrar, srv AdminAlertServiceServer) {
	s.RegisterService(&AdminAlertService_ServiceDesc, srv)
}

func _AdminAlertService_CreateAdminAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdminAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAlertServiceServer).CreateAdminAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminAlertService_CreateAdminAlert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAlertServiceServer).CreateAdminAlert(ctx, req.(*CreateAdminAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAlertService_GetAdminAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAlertServiceServer).GetAdminAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminAlertService_GetAdminAlert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAlertServiceServer).GetAdminAlert(ctx, req.(*AdminAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAlertService_DeleteAdminAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAlertServiceServer).DeleteAdminAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminAlertService_DeleteAdminAlert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAlertServiceServer).DeleteAdminAlert(ctx, req.(*AdminAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAlertService_ListAdminAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAlertServiceServer).ListAdminAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminAlertService_ListAdminAlerts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAlertServiceServer).ListAdminAlerts(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminAlertService_ServiceDesc is the grpc.ServiceDesc for AdminAlertService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminAlertService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.AdminAlertService",
	HandlerType: (*AdminAlertServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAdminAlert",
			Handler:    _AdminAlertService_CreateAdminAlert_Handler,
		},
		{
			MethodName: "GetAdminAlert",
			Handler:    _AdminAlertService_GetAdminAlert_Handler,
		},
		{
			MethodName: "DeleteAdminAlert",
			Handler:    _AdminAlertService_DeleteAdminAlert_Handler,
		},
		{
			MethodName: "ListAdminAlerts",
			Handler:    _AdminAlertService_ListAdminAlerts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "adminAlert.proto",
}
