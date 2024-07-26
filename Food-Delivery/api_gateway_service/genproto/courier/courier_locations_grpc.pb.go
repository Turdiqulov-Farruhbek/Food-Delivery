// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: courier_locations.proto

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
	CourierLocationService_CreateCourierLocation_FullMethodName  = "/courierlocations.CourierLocationService/CreateCourierLocation"
	CourierLocationService_GetCourierLocation_FullMethodName     = "/courierlocations.CourierLocationService/GetCourierLocation"
	CourierLocationService_UpdateCourierLocation_FullMethodName  = "/courierlocations.CourierLocationService/UpdateCourierLocation"
	CourierLocationService_DeleteCourierLocation_FullMethodName  = "/courierlocations.CourierLocationService/DeleteCourierLocation"
	CourierLocationService_GetAllCourierLocations_FullMethodName = "/courierlocations.CourierLocationService/GetAllCourierLocations"
)

// CourierLocationServiceClient is the client API for CourierLocationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service definition for courier location management
type CourierLocationServiceClient interface {
	CreateCourierLocation(ctx context.Context, in *CreateCourierLocationRequest, opts ...grpc.CallOption) (*CourierLocationResponse, error)
	GetCourierLocation(ctx context.Context, in *GetCourierLocationRequest, opts ...grpc.CallOption) (*CourierLocationResponse, error)
	UpdateCourierLocation(ctx context.Context, in *UpdateCourierLocationRequest, opts ...grpc.CallOption) (*CourierLocationResponse, error)
	DeleteCourierLocation(ctx context.Context, in *DeleteCourierLocationRequest, opts ...grpc.CallOption) (*DeleteCourierLocationResponse, error)
	GetAllCourierLocations(ctx context.Context, in *GetAllCourierLocationsRequest, opts ...grpc.CallOption) (*GetAllCourierLocationsResponse, error)
}

type courierLocationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCourierLocationServiceClient(cc grpc.ClientConnInterface) CourierLocationServiceClient {
	return &courierLocationServiceClient{cc}
}

func (c *courierLocationServiceClient) CreateCourierLocation(ctx context.Context, in *CreateCourierLocationRequest, opts ...grpc.CallOption) (*CourierLocationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourierLocationResponse)
	err := c.cc.Invoke(ctx, CourierLocationService_CreateCourierLocation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierLocationServiceClient) GetCourierLocation(ctx context.Context, in *GetCourierLocationRequest, opts ...grpc.CallOption) (*CourierLocationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourierLocationResponse)
	err := c.cc.Invoke(ctx, CourierLocationService_GetCourierLocation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierLocationServiceClient) UpdateCourierLocation(ctx context.Context, in *UpdateCourierLocationRequest, opts ...grpc.CallOption) (*CourierLocationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourierLocationResponse)
	err := c.cc.Invoke(ctx, CourierLocationService_UpdateCourierLocation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierLocationServiceClient) DeleteCourierLocation(ctx context.Context, in *DeleteCourierLocationRequest, opts ...grpc.CallOption) (*DeleteCourierLocationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCourierLocationResponse)
	err := c.cc.Invoke(ctx, CourierLocationService_DeleteCourierLocation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierLocationServiceClient) GetAllCourierLocations(ctx context.Context, in *GetAllCourierLocationsRequest, opts ...grpc.CallOption) (*GetAllCourierLocationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllCourierLocationsResponse)
	err := c.cc.Invoke(ctx, CourierLocationService_GetAllCourierLocations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourierLocationServiceServer is the server API for CourierLocationService service.
// All implementations must embed UnimplementedCourierLocationServiceServer
// for forward compatibility
//
// Service definition for courier location management
type CourierLocationServiceServer interface {
	CreateCourierLocation(context.Context, *CreateCourierLocationRequest) (*CourierLocationResponse, error)
	GetCourierLocation(context.Context, *GetCourierLocationRequest) (*CourierLocationResponse, error)
	UpdateCourierLocation(context.Context, *UpdateCourierLocationRequest) (*CourierLocationResponse, error)
	DeleteCourierLocation(context.Context, *DeleteCourierLocationRequest) (*DeleteCourierLocationResponse, error)
	GetAllCourierLocations(context.Context, *GetAllCourierLocationsRequest) (*GetAllCourierLocationsResponse, error)
	mustEmbedUnimplementedCourierLocationServiceServer()
}

// UnimplementedCourierLocationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCourierLocationServiceServer struct {
}

func (UnimplementedCourierLocationServiceServer) CreateCourierLocation(context.Context, *CreateCourierLocationRequest) (*CourierLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourierLocation not implemented")
}
func (UnimplementedCourierLocationServiceServer) GetCourierLocation(context.Context, *GetCourierLocationRequest) (*CourierLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourierLocation not implemented")
}
func (UnimplementedCourierLocationServiceServer) UpdateCourierLocation(context.Context, *UpdateCourierLocationRequest) (*CourierLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCourierLocation not implemented")
}
func (UnimplementedCourierLocationServiceServer) DeleteCourierLocation(context.Context, *DeleteCourierLocationRequest) (*DeleteCourierLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCourierLocation not implemented")
}
func (UnimplementedCourierLocationServiceServer) GetAllCourierLocations(context.Context, *GetAllCourierLocationsRequest) (*GetAllCourierLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCourierLocations not implemented")
}
func (UnimplementedCourierLocationServiceServer) mustEmbedUnimplementedCourierLocationServiceServer() {
}

// UnsafeCourierLocationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourierLocationServiceServer will
// result in compilation errors.
type UnsafeCourierLocationServiceServer interface {
	mustEmbedUnimplementedCourierLocationServiceServer()
}

func RegisterCourierLocationServiceServer(s grpc.ServiceRegistrar, srv CourierLocationServiceServer) {
	s.RegisterService(&CourierLocationService_ServiceDesc, srv)
}

func _CourierLocationService_CreateCourierLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourierLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierLocationServiceServer).CreateCourierLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourierLocationService_CreateCourierLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierLocationServiceServer).CreateCourierLocation(ctx, req.(*CreateCourierLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierLocationService_GetCourierLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourierLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierLocationServiceServer).GetCourierLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourierLocationService_GetCourierLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierLocationServiceServer).GetCourierLocation(ctx, req.(*GetCourierLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierLocationService_UpdateCourierLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCourierLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierLocationServiceServer).UpdateCourierLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourierLocationService_UpdateCourierLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierLocationServiceServer).UpdateCourierLocation(ctx, req.(*UpdateCourierLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierLocationService_DeleteCourierLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCourierLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierLocationServiceServer).DeleteCourierLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourierLocationService_DeleteCourierLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierLocationServiceServer).DeleteCourierLocation(ctx, req.(*DeleteCourierLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierLocationService_GetAllCourierLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCourierLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierLocationServiceServer).GetAllCourierLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourierLocationService_GetAllCourierLocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierLocationServiceServer).GetAllCourierLocations(ctx, req.(*GetAllCourierLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CourierLocationService_ServiceDesc is the grpc.ServiceDesc for CourierLocationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CourierLocationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "courierlocations.CourierLocationService",
	HandlerType: (*CourierLocationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCourierLocation",
			Handler:    _CourierLocationService_CreateCourierLocation_Handler,
		},
		{
			MethodName: "GetCourierLocation",
			Handler:    _CourierLocationService_GetCourierLocation_Handler,
		},
		{
			MethodName: "UpdateCourierLocation",
			Handler:    _CourierLocationService_UpdateCourierLocation_Handler,
		},
		{
			MethodName: "DeleteCourierLocation",
			Handler:    _CourierLocationService_DeleteCourierLocation_Handler,
		},
		{
			MethodName: "GetAllCourierLocations",
			Handler:    _CourierLocationService_GetAllCourierLocations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "courier_locations.proto",
}
