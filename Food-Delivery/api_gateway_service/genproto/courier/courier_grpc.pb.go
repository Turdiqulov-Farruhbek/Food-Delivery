// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: courier.proto

package courier

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
	CourierService_CreateCourier_FullMethodName = "/ecommerce.CourierService/CreateCourier"
	CourierService_GetCourier_FullMethodName    = "/ecommerce.CourierService/GetCourier"
	CourierService_UpdateCourier_FullMethodName = "/ecommerce.CourierService/UpdateCourier"
	CourierService_DeleteCourier_FullMethodName = "/ecommerce.CourierService/DeleteCourier"
	CourierService_ListCouriers_FullMethodName  = "/ecommerce.CourierService/ListCouriers"
)

// CourierServiceClient is the client API for CourierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage couriers
type CourierServiceClient interface {
	// Kuryer yaratish
	CreateCourier(ctx context.Context, in *CreateCourierRequest, opts ...grpc.CallOption) (*CourierResponse, error)
	// Kuryer ma'lumotlarini olish
	GetCourier(ctx context.Context, in *CourierRequest, opts ...grpc.CallOption) (*CourierResponse, error)
	// Kuryerni yangilash
	UpdateCourier(ctx context.Context, in *UpdateCourierRequest, opts ...grpc.CallOption) (*CourierResponse, error)
	// Kuryerni o'chirish
	DeleteCourier(ctx context.Context, in *CourierRequest, opts ...grpc.CallOption) (*CourierResponse, error)
	// Barcha kuryerlar ro'yxatini olish
	ListCouriers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CourierListResponse, error)
}

type courierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCourierServiceClient(cc grpc.ClientConnInterface) CourierServiceClient {
	return &courierServiceClient{cc}
}

func (c *courierServiceClient) CreateCourier(ctx context.Context, in *CreateCourierRequest, opts ...grpc.CallOption) (*CourierResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourierResponse)
	err := c.cc.Invoke(ctx, CourierService_CreateCourier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierServiceClient) GetCourier(ctx context.Context, in *CourierRequest, opts ...grpc.CallOption) (*CourierResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourierResponse)
	err := c.cc.Invoke(ctx, CourierService_GetCourier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierServiceClient) UpdateCourier(ctx context.Context, in *UpdateCourierRequest, opts ...grpc.CallOption) (*CourierResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourierResponse)
	err := c.cc.Invoke(ctx, CourierService_UpdateCourier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierServiceClient) DeleteCourier(ctx context.Context, in *CourierRequest, opts ...grpc.CallOption) (*CourierResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourierResponse)
	err := c.cc.Invoke(ctx, CourierService_DeleteCourier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierServiceClient) ListCouriers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CourierListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourierListResponse)
	err := c.cc.Invoke(ctx, CourierService_ListCouriers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourierServiceServer is the server API for CourierService service.
// All implementations must embed UnimplementedCourierServiceServer
// for forward compatibility
//
// Service to manage couriers
type CourierServiceServer interface {
	// Kuryer yaratish
	CreateCourier(context.Context, *CreateCourierRequest) (*CourierResponse, error)
	// Kuryer ma'lumotlarini olish
	GetCourier(context.Context, *CourierRequest) (*CourierResponse, error)
	// Kuryerni yangilash
	UpdateCourier(context.Context, *UpdateCourierRequest) (*CourierResponse, error)
	// Kuryerni o'chirish
	DeleteCourier(context.Context, *CourierRequest) (*CourierResponse, error)
	// Barcha kuryerlar ro'yxatini olish
	ListCouriers(context.Context, *Empty) (*CourierListResponse, error)
	mustEmbedUnimplementedCourierServiceServer()
}

// UnimplementedCourierServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCourierServiceServer struct {
}

func (UnimplementedCourierServiceServer) CreateCourier(context.Context, *CreateCourierRequest) (*CourierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourier not implemented")
}
func (UnimplementedCourierServiceServer) GetCourier(context.Context, *CourierRequest) (*CourierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourier not implemented")
}
func (UnimplementedCourierServiceServer) UpdateCourier(context.Context, *UpdateCourierRequest) (*CourierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCourier not implemented")
}
func (UnimplementedCourierServiceServer) DeleteCourier(context.Context, *CourierRequest) (*CourierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCourier not implemented")
}
func (UnimplementedCourierServiceServer) ListCouriers(context.Context, *Empty) (*CourierListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCouriers not implemented")
}
func (UnimplementedCourierServiceServer) mustEmbedUnimplementedCourierServiceServer() {}

// UnsafeCourierServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourierServiceServer will
// result in compilation errors.
type UnsafeCourierServiceServer interface {
	mustEmbedUnimplementedCourierServiceServer()
}

func RegisterCourierServiceServer(s grpc.ServiceRegistrar, srv CourierServiceServer) {
	s.RegisterService(&CourierService_ServiceDesc, srv)
}

func _CourierService_CreateCourier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServiceServer).CreateCourier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourierService_CreateCourier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServiceServer).CreateCourier(ctx, req.(*CreateCourierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierService_GetCourier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServiceServer).GetCourier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourierService_GetCourier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServiceServer).GetCourier(ctx, req.(*CourierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierService_UpdateCourier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCourierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServiceServer).UpdateCourier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourierService_UpdateCourier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServiceServer).UpdateCourier(ctx, req.(*UpdateCourierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierService_DeleteCourier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServiceServer).DeleteCourier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourierService_DeleteCourier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServiceServer).DeleteCourier(ctx, req.(*CourierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierService_ListCouriers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServiceServer).ListCouriers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourierService_ListCouriers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServiceServer).ListCouriers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CourierService_ServiceDesc is the grpc.ServiceDesc for CourierService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CourierService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.CourierService",
	HandlerType: (*CourierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCourier",
			Handler:    _CourierService_CreateCourier_Handler,
		},
		{
			MethodName: "GetCourier",
			Handler:    _CourierService_GetCourier_Handler,
		},
		{
			MethodName: "UpdateCourier",
			Handler:    _CourierService_UpdateCourier_Handler,
		},
		{
			MethodName: "DeleteCourier",
			Handler:    _CourierService_DeleteCourier_Handler,
		},
		{
			MethodName: "ListCouriers",
			Handler:    _CourierService_ListCouriers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "courier.proto",
}