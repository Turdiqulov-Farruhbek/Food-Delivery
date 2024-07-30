// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: orders.proto

package product

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
	OrderProductService_CreateOrder_FullMethodName = "/ecommerce.OrderProductService/CreateOrder"
	OrderProductService_GetOrder_FullMethodName    = "/ecommerce.OrderProductService/GetOrder"
	OrderProductService_UpdateOrder_FullMethodName = "/ecommerce.OrderProductService/UpdateOrder"
	OrderProductService_DeleteOrder_FullMethodName = "/ecommerce.OrderProductService/DeleteOrder"
	OrderProductService_ListOrders_FullMethodName  = "/ecommerce.OrderProductService/ListOrders"
)

// OrderProductServiceClient is the client API for OrderProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage orders
type OrderProductServiceClient interface {
	// Buyurtma yaratish
	CreateOrder(ctx context.Context, in *CreateOrderProductRequest, opts ...grpc.CallOption) (*OrderProductResponse, error)
	// Buyurtma ma'lumotlarini olish
	GetOrder(ctx context.Context, in *OrderProductRequest, opts ...grpc.CallOption) (*OrderProductResponse, error)
	// Buyurtmani yangilash
	UpdateOrder(ctx context.Context, in *UpdateOrderProductRequest, opts ...grpc.CallOption) (*OrderProductResponse, error)
	// Buyurtmani o'chirish
	DeleteOrder(ctx context.Context, in *OrderProductRequest, opts ...grpc.CallOption) (*OrderProductResponse, error)
	// Buyurtmalar ro'yxatini olish
	ListOrders(ctx context.Context, in *OrderProductListRequest, opts ...grpc.CallOption) (*OrderProductListResponse, error)
}

type orderProductServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderProductServiceClient(cc grpc.ClientConnInterface) OrderProductServiceClient {
	return &orderProductServiceClient{cc}
}

func (c *orderProductServiceClient) CreateOrder(ctx context.Context, in *CreateOrderProductRequest, opts ...grpc.CallOption) (*OrderProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderProductResponse)
	err := c.cc.Invoke(ctx, OrderProductService_CreateOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderProductServiceClient) GetOrder(ctx context.Context, in *OrderProductRequest, opts ...grpc.CallOption) (*OrderProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderProductResponse)
	err := c.cc.Invoke(ctx, OrderProductService_GetOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderProductServiceClient) UpdateOrder(ctx context.Context, in *UpdateOrderProductRequest, opts ...grpc.CallOption) (*OrderProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderProductResponse)
	err := c.cc.Invoke(ctx, OrderProductService_UpdateOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderProductServiceClient) DeleteOrder(ctx context.Context, in *OrderProductRequest, opts ...grpc.CallOption) (*OrderProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderProductResponse)
	err := c.cc.Invoke(ctx, OrderProductService_DeleteOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderProductServiceClient) ListOrders(ctx context.Context, in *OrderProductListRequest, opts ...grpc.CallOption) (*OrderProductListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderProductListResponse)
	err := c.cc.Invoke(ctx, OrderProductService_ListOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderProductServiceServer is the server API for OrderProductService service.
// All implementations must embed UnimplementedOrderProductServiceServer
// for forward compatibility
//
// Service to manage orders
type OrderProductServiceServer interface {
	// Buyurtma yaratish
	CreateOrder(context.Context, *CreateOrderProductRequest) (*OrderProductResponse, error)
	// Buyurtma ma'lumotlarini olish
	GetOrder(context.Context, *OrderProductRequest) (*OrderProductResponse, error)
	// Buyurtmani yangilash
	UpdateOrder(context.Context, *UpdateOrderProductRequest) (*OrderProductResponse, error)
	// Buyurtmani o'chirish
	DeleteOrder(context.Context, *OrderProductRequest) (*OrderProductResponse, error)
	// Buyurtmalar ro'yxatini olish
	ListOrders(context.Context, *OrderProductListRequest) (*OrderProductListResponse, error)
	mustEmbedUnimplementedOrderProductServiceServer()
}

// UnimplementedOrderProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderProductServiceServer struct {
}

func (UnimplementedOrderProductServiceServer) CreateOrder(context.Context, *CreateOrderProductRequest) (*OrderProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderProductServiceServer) GetOrder(context.Context, *OrderProductRequest) (*OrderProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderProductServiceServer) UpdateOrder(context.Context, *UpdateOrderProductRequest) (*OrderProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedOrderProductServiceServer) DeleteOrder(context.Context, *OrderProductRequest) (*OrderProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedOrderProductServiceServer) ListOrders(context.Context, *OrderProductListRequest) (*OrderProductListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}
func (UnimplementedOrderProductServiceServer) mustEmbedUnimplementedOrderProductServiceServer() {}

// UnsafeOrderProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderProductServiceServer will
// result in compilation errors.
type UnsafeOrderProductServiceServer interface {
	mustEmbedUnimplementedOrderProductServiceServer()
}

func RegisterOrderProductServiceServer(s grpc.ServiceRegistrar, srv OrderProductServiceServer) {
	s.RegisterService(&OrderProductService_ServiceDesc, srv)
}

func _OrderProductService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderProductServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderProductService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderProductServiceServer).CreateOrder(ctx, req.(*CreateOrderProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderProductService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderProductServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderProductService_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderProductServiceServer).GetOrder(ctx, req.(*OrderProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderProductService_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderProductServiceServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderProductService_UpdateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderProductServiceServer).UpdateOrder(ctx, req.(*UpdateOrderProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderProductService_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderProductServiceServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderProductService_DeleteOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderProductServiceServer).DeleteOrder(ctx, req.(*OrderProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderProductService_ListOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderProductServiceServer).ListOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderProductService_ListOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderProductServiceServer).ListOrders(ctx, req.(*OrderProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderProductService_ServiceDesc is the grpc.ServiceDesc for OrderProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.OrderProductService",
	HandlerType: (*OrderProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderProductService_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _OrderProductService_GetOrder_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _OrderProductService_UpdateOrder_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _OrderProductService_DeleteOrder_Handler,
		},
		{
			MethodName: "ListOrders",
			Handler:    _OrderProductService_ListOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders.proto",
}
