// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: cart_Items.proto

package genproto

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
	CartItemService_CreateCartItem_FullMethodName = "/ecommerce.CartItemService/CreateCartItem"
	CartItemService_GetCartItem_FullMethodName    = "/ecommerce.CartItemService/GetCartItem"
	CartItemService_UpdateCartItem_FullMethodName = "/ecommerce.CartItemService/UpdateCartItem"
	CartItemService_DeleteCartItem_FullMethodName = "/ecommerce.CartItemService/DeleteCartItem"
	CartItemService_ListCartItems_FullMethodName  = "/ecommerce.CartItemService/ListCartItems"
)

// CartItemServiceClient is the client API for CartItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage cart items
type CartItemServiceClient interface {
	// Savat mahsulotini yaratish
	CreateCartItem(ctx context.Context, in *CreateCartItemRequest, opts ...grpc.CallOption) (*CartItemResponse, error)
	// Savat mahsuloti ma'lumotlarini olish
	GetCartItem(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartItemResponse, error)
	// Savat mahsulotini yangilash
	UpdateCartItem(ctx context.Context, in *UpdateCartItemRequest, opts ...grpc.CallOption) (*CartItemResponse, error)
	// Savat mahsulotini o'chirish
	DeleteCartItem(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartItemResponse, error)
	// Savat mahsulotlari ro'yxatini olish
	ListCartItems(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartItemListResponse, error)
}

type cartItemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartItemServiceClient(cc grpc.ClientConnInterface) CartItemServiceClient {
	return &cartItemServiceClient{cc}
}

func (c *cartItemServiceClient) CreateCartItem(ctx context.Context, in *CreateCartItemRequest, opts ...grpc.CallOption) (*CartItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartItemResponse)
	err := c.cc.Invoke(ctx, CartItemService_CreateCartItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartItemServiceClient) GetCartItem(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartItemResponse)
	err := c.cc.Invoke(ctx, CartItemService_GetCartItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartItemServiceClient) UpdateCartItem(ctx context.Context, in *UpdateCartItemRequest, opts ...grpc.CallOption) (*CartItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartItemResponse)
	err := c.cc.Invoke(ctx, CartItemService_UpdateCartItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartItemServiceClient) DeleteCartItem(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartItemResponse)
	err := c.cc.Invoke(ctx, CartItemService_DeleteCartItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartItemServiceClient) ListCartItems(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartItemListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartItemListResponse)
	err := c.cc.Invoke(ctx, CartItemService_ListCartItems_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartItemServiceServer is the server API for CartItemService service.
// All implementations must embed UnimplementedCartItemServiceServer
// for forward compatibility
//
// Service to manage cart items
type CartItemServiceServer interface {
	// Savat mahsulotini yaratish
	CreateCartItem(context.Context, *CreateCartItemRequest) (*CartItemResponse, error)
	// Savat mahsuloti ma'lumotlarini olish
	GetCartItem(context.Context, *CartItemRequest) (*CartItemResponse, error)
	// Savat mahsulotini yangilash
	UpdateCartItem(context.Context, *UpdateCartItemRequest) (*CartItemResponse, error)
	// Savat mahsulotini o'chirish
	DeleteCartItem(context.Context, *CartItemRequest) (*CartItemResponse, error)
	// Savat mahsulotlari ro'yxatini olish
	ListCartItems(context.Context, *CartItemRequest) (*CartItemListResponse, error)
	mustEmbedUnimplementedCartItemServiceServer()
}

// UnimplementedCartItemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCartItemServiceServer struct {
}

func (UnimplementedCartItemServiceServer) CreateCartItem(context.Context, *CreateCartItemRequest) (*CartItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCartItem not implemented")
}
func (UnimplementedCartItemServiceServer) GetCartItem(context.Context, *CartItemRequest) (*CartItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCartItem not implemented")
}
func (UnimplementedCartItemServiceServer) UpdateCartItem(context.Context, *UpdateCartItemRequest) (*CartItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCartItem not implemented")
}
func (UnimplementedCartItemServiceServer) DeleteCartItem(context.Context, *CartItemRequest) (*CartItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCartItem not implemented")
}
func (UnimplementedCartItemServiceServer) ListCartItems(context.Context, *CartItemRequest) (*CartItemListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCartItems not implemented")
}
func (UnimplementedCartItemServiceServer) mustEmbedUnimplementedCartItemServiceServer() {}

// UnsafeCartItemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartItemServiceServer will
// result in compilation errors.
type UnsafeCartItemServiceServer interface {
	mustEmbedUnimplementedCartItemServiceServer()
}

func RegisterCartItemServiceServer(s grpc.ServiceRegistrar, srv CartItemServiceServer) {
	s.RegisterService(&CartItemService_ServiceDesc, srv)
}

func _CartItemService_CreateCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartItemServiceServer).CreateCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartItemService_CreateCartItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartItemServiceServer).CreateCartItem(ctx, req.(*CreateCartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartItemService_GetCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartItemServiceServer).GetCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartItemService_GetCartItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartItemServiceServer).GetCartItem(ctx, req.(*CartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartItemService_UpdateCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartItemServiceServer).UpdateCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartItemService_UpdateCartItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartItemServiceServer).UpdateCartItem(ctx, req.(*UpdateCartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartItemService_DeleteCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartItemServiceServer).DeleteCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartItemService_DeleteCartItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartItemServiceServer).DeleteCartItem(ctx, req.(*CartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartItemService_ListCartItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartItemServiceServer).ListCartItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartItemService_ListCartItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartItemServiceServer).ListCartItems(ctx, req.(*CartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CartItemService_ServiceDesc is the grpc.ServiceDesc for CartItemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CartItemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.CartItemService",
	HandlerType: (*CartItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCartItem",
			Handler:    _CartItemService_CreateCartItem_Handler,
		},
		{
			MethodName: "GetCartItem",
			Handler:    _CartItemService_GetCartItem_Handler,
		},
		{
			MethodName: "UpdateCartItem",
			Handler:    _CartItemService_UpdateCartItem_Handler,
		},
		{
			MethodName: "DeleteCartItem",
			Handler:    _CartItemService_DeleteCartItem_Handler,
		},
		{
			MethodName: "ListCartItems",
			Handler:    _CartItemService_ListCartItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cart_Items.proto",
}