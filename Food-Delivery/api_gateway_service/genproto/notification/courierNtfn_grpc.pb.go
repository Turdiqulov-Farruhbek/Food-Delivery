// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: courierNtfn.proto

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
	CourierNotificationService_CreateCourierNotification_FullMethodName = "/ecommerce.CourierNotificationService/CreateCourierNotification"
	CourierNotificationService_GetCourierNotification_FullMethodName    = "/ecommerce.CourierNotificationService/GetCourierNotification"
	CourierNotificationService_DeleteCourierNotification_FullMethodName = "/ecommerce.CourierNotificationService/DeleteCourierNotification"
	CourierNotificationService_ListCourierNotifications_FullMethodName  = "/ecommerce.CourierNotificationService/ListCourierNotifications"
)

// CourierNotificationServiceClient is the client API for CourierNotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage courier notifications
type CourierNotificationServiceClient interface {
	// Kuryer bildirishnomasini yaratish
	CreateCourierNotification(ctx context.Context, in *CreateCourierNotificationRequest, opts ...grpc.CallOption) (*CourierNotificationResponse, error)
	// Kuryer bildirishnomasini olish
	GetCourierNotification(ctx context.Context, in *CourierNotificationRequest, opts ...grpc.CallOption) (*CourierNotificationResponse, error)
	// Kuryer bildirishnomasini o'chirish
	DeleteCourierNotification(ctx context.Context, in *CourierNotificationRequest, opts ...grpc.CallOption) (*CourierNotificationResponse, error)
	// Barcha kuryer bildirishnomalarini olish
	ListCourierNotifications(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CourierNotificationListResponse, error)
}

type courierNotificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCourierNotificationServiceClient(cc grpc.ClientConnInterface) CourierNotificationServiceClient {
	return &courierNotificationServiceClient{cc}
}

func (c *courierNotificationServiceClient) CreateCourierNotification(ctx context.Context, in *CreateCourierNotificationRequest, opts ...grpc.CallOption) (*CourierNotificationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourierNotificationResponse)
	err := c.cc.Invoke(ctx, CourierNotificationService_CreateCourierNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierNotificationServiceClient) GetCourierNotification(ctx context.Context, in *CourierNotificationRequest, opts ...grpc.CallOption) (*CourierNotificationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourierNotificationResponse)
	err := c.cc.Invoke(ctx, CourierNotificationService_GetCourierNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierNotificationServiceClient) DeleteCourierNotification(ctx context.Context, in *CourierNotificationRequest, opts ...grpc.CallOption) (*CourierNotificationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourierNotificationResponse)
	err := c.cc.Invoke(ctx, CourierNotificationService_DeleteCourierNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierNotificationServiceClient) ListCourierNotifications(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CourierNotificationListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourierNotificationListResponse)
	err := c.cc.Invoke(ctx, CourierNotificationService_ListCourierNotifications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourierNotificationServiceServer is the server API for CourierNotificationService service.
// All implementations must embed UnimplementedCourierNotificationServiceServer
// for forward compatibility
//
// Service to manage courier notifications
type CourierNotificationServiceServer interface {
	// Kuryer bildirishnomasini yaratish
	CreateCourierNotification(context.Context, *CreateCourierNotificationRequest) (*CourierNotificationResponse, error)
	// Kuryer bildirishnomasini olish
	GetCourierNotification(context.Context, *CourierNotificationRequest) (*CourierNotificationResponse, error)
	// Kuryer bildirishnomasini o'chirish
	DeleteCourierNotification(context.Context, *CourierNotificationRequest) (*CourierNotificationResponse, error)
	// Barcha kuryer bildirishnomalarini olish
	ListCourierNotifications(context.Context, *Empty) (*CourierNotificationListResponse, error)
	mustEmbedUnimplementedCourierNotificationServiceServer()
}

// UnimplementedCourierNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCourierNotificationServiceServer struct {
}

func (UnimplementedCourierNotificationServiceServer) CreateCourierNotification(context.Context, *CreateCourierNotificationRequest) (*CourierNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourierNotification not implemented")
}
func (UnimplementedCourierNotificationServiceServer) GetCourierNotification(context.Context, *CourierNotificationRequest) (*CourierNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourierNotification not implemented")
}
func (UnimplementedCourierNotificationServiceServer) DeleteCourierNotification(context.Context, *CourierNotificationRequest) (*CourierNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCourierNotification not implemented")
}
func (UnimplementedCourierNotificationServiceServer) ListCourierNotifications(context.Context, *Empty) (*CourierNotificationListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCourierNotifications not implemented")
}
func (UnimplementedCourierNotificationServiceServer) mustEmbedUnimplementedCourierNotificationServiceServer() {
}

// UnsafeCourierNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourierNotificationServiceServer will
// result in compilation errors.
type UnsafeCourierNotificationServiceServer interface {
	mustEmbedUnimplementedCourierNotificationServiceServer()
}

func RegisterCourierNotificationServiceServer(s grpc.ServiceRegistrar, srv CourierNotificationServiceServer) {
	s.RegisterService(&CourierNotificationService_ServiceDesc, srv)
}

func _CourierNotificationService_CreateCourierNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourierNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierNotificationServiceServer).CreateCourierNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourierNotificationService_CreateCourierNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierNotificationServiceServer).CreateCourierNotification(ctx, req.(*CreateCourierNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierNotificationService_GetCourierNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourierNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierNotificationServiceServer).GetCourierNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourierNotificationService_GetCourierNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierNotificationServiceServer).GetCourierNotification(ctx, req.(*CourierNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierNotificationService_DeleteCourierNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourierNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierNotificationServiceServer).DeleteCourierNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourierNotificationService_DeleteCourierNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierNotificationServiceServer).DeleteCourierNotification(ctx, req.(*CourierNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierNotificationService_ListCourierNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierNotificationServiceServer).ListCourierNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourierNotificationService_ListCourierNotifications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierNotificationServiceServer).ListCourierNotifications(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CourierNotificationService_ServiceDesc is the grpc.ServiceDesc for CourierNotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CourierNotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.CourierNotificationService",
	HandlerType: (*CourierNotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCourierNotification",
			Handler:    _CourierNotificationService_CreateCourierNotification_Handler,
		},
		{
			MethodName: "GetCourierNotification",
			Handler:    _CourierNotificationService_GetCourierNotification_Handler,
		},
		{
			MethodName: "DeleteCourierNotification",
			Handler:    _CourierNotificationService_DeleteCourierNotification_Handler,
		},
		{
			MethodName: "ListCourierNotifications",
			Handler:    _CourierNotificationService_ListCourierNotifications_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "courierNtfn.proto",
}