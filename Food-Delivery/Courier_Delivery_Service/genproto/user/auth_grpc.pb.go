// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: auth.proto

package user

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
	AuthService_UserLogin_FullMethodName       = "/ecommerce.AuthService/UserLogin"
	AuthService_UserRegister_FullMethodName    = "/ecommerce.AuthService/UserRegister"
	AuthService_CourierLogin_FullMethodName    = "/ecommerce.AuthService/CourierLogin"
	AuthService_CourierRegister_FullMethodName = "/ecommerce.AuthService/CourierRegister"
	AuthService_CreateUser_FullMethodName      = "/ecommerce.AuthService/CreateUser"
	AuthService_GetUser_FullMethodName         = "/ecommerce.AuthService/GetUser"
	AuthService_GetAllUsers_FullMethodName     = "/ecommerce.AuthService/GetAllUsers"
	AuthService_UpdateUser_FullMethodName      = "/ecommerce.AuthService/UpdateUser"
	AuthService_DeleteUser_FullMethodName      = "/ecommerce.AuthService/DeleteUser"
	AuthService_CreateCourier_FullMethodName   = "/ecommerce.AuthService/CreateCourier"
	AuthService_GetCourier_FullMethodName      = "/ecommerce.AuthService/GetCourier"
	AuthService_GetAllCouriers_FullMethodName  = "/ecommerce.AuthService/GetAllCouriers"
	AuthService_UpdateCourier_FullMethodName   = "/ecommerce.AuthService/UpdateCourier"
	AuthService_DeleteCourier_FullMethodName   = "/ecommerce.AuthService/DeleteCourier"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage user and courier authentication and CRUD operations
type AuthServiceClient interface {
	// Foydalanuvchi uchun autentifikatsiyani amalga oshirish
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	// Foydalanuvchi uchun ro'yxatdan o'tish
	UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error)
	// Kuryer uchun autentifikatsiyani amalga oshirish
	CourierLogin(ctx context.Context, in *CourierLoginRequest, opts ...grpc.CallOption) (*CourierLoginResponse, error)
	// Kuryer uchun ro'yxatdan o'tish
	CourierRegister(ctx context.Context, in *CourierRegisterRequest, opts ...grpc.CallOption) (*CourierRegisterResponse, error)
	// Foydalanuvchi yaratish
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	// Foydalanuvchi ma'lumotlarini olish
	GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	// Foydalanuvchilar ro'yxatini olish
	GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error)
	// Foydalanuvchi ma'lumotlarini yangilash
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	// Foydalanuvchi o'chirish
	DeleteUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	// Kuryer yaratish
	CreateCourier(ctx context.Context, in *CreateCourierRequestAuth, opts ...grpc.CallOption) (*CourierResponseAuth, error)
	// Kuryer ma'lumotlarini olish
	GetCourier(ctx context.Context, in *CourierRequestAuth, opts ...grpc.CallOption) (*CourierResponseAuth, error)
	// Kuryerlar ro'yxatini olish
	GetAllCouriers(ctx context.Context, in *GetAllCouriersRequest, opts ...grpc.CallOption) (*GetAllCouriersResponse, error)
	// Kuryer ma'lumotlarini yangilash
	UpdateCourier(ctx context.Context, in *UpdateCourierRequestAuth, opts ...grpc.CallOption) (*CourierResponseAuth, error)
	// Kuryer o'chirish
	DeleteCourier(ctx context.Context, in *CourierRequestAuth, opts ...grpc.CallOption) (*CourierResponseAuth, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, AuthService_UserLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserRegisterResponse)
	err := c.cc.Invoke(ctx, AuthService_UserRegister_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CourierLogin(ctx context.Context, in *CourierLoginRequest, opts ...grpc.CallOption) (*CourierLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourierLoginResponse)
	err := c.cc.Invoke(ctx, AuthService_CourierLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CourierRegister(ctx context.Context, in *CourierRegisterRequest, opts ...grpc.CallOption) (*CourierRegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourierRegisterResponse)
	err := c.cc.Invoke(ctx, AuthService_CourierRegister_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, AuthService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, AuthService_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllUsersResponse)
	err := c.cc.Invoke(ctx, AuthService_GetAllUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, AuthService_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeleteUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, AuthService_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CreateCourier(ctx context.Context, in *CreateCourierRequestAuth, opts ...grpc.CallOption) (*CourierResponseAuth, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourierResponseAuth)
	err := c.cc.Invoke(ctx, AuthService_CreateCourier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetCourier(ctx context.Context, in *CourierRequestAuth, opts ...grpc.CallOption) (*CourierResponseAuth, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourierResponseAuth)
	err := c.cc.Invoke(ctx, AuthService_GetCourier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetAllCouriers(ctx context.Context, in *GetAllCouriersRequest, opts ...grpc.CallOption) (*GetAllCouriersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllCouriersResponse)
	err := c.cc.Invoke(ctx, AuthService_GetAllCouriers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateCourier(ctx context.Context, in *UpdateCourierRequestAuth, opts ...grpc.CallOption) (*CourierResponseAuth, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourierResponseAuth)
	err := c.cc.Invoke(ctx, AuthService_UpdateCourier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeleteCourier(ctx context.Context, in *CourierRequestAuth, opts ...grpc.CallOption) (*CourierResponseAuth, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourierResponseAuth)
	err := c.cc.Invoke(ctx, AuthService_DeleteCourier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
//
// Service to manage user and courier authentication and CRUD operations
type AuthServiceServer interface {
	// Foydalanuvchi uchun autentifikatsiyani amalga oshirish
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	// Foydalanuvchi uchun ro'yxatdan o'tish
	UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error)
	// Kuryer uchun autentifikatsiyani amalga oshirish
	CourierLogin(context.Context, *CourierLoginRequest) (*CourierLoginResponse, error)
	// Kuryer uchun ro'yxatdan o'tish
	CourierRegister(context.Context, *CourierRegisterRequest) (*CourierRegisterResponse, error)
	// Foydalanuvchi yaratish
	CreateUser(context.Context, *CreateUserRequest) (*UserResponse, error)
	// Foydalanuvchi ma'lumotlarini olish
	GetUser(context.Context, *UserRequest) (*UserResponse, error)
	// Foydalanuvchilar ro'yxatini olish
	GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error)
	// Foydalanuvchi ma'lumotlarini yangilash
	UpdateUser(context.Context, *UpdateUserRequest) (*UserResponse, error)
	// Foydalanuvchi o'chirish
	DeleteUser(context.Context, *UserRequest) (*UserResponse, error)
	// Kuryer yaratish
	CreateCourier(context.Context, *CreateCourierRequestAuth) (*CourierResponseAuth, error)
	// Kuryer ma'lumotlarini olish
	GetCourier(context.Context, *CourierRequestAuth) (*CourierResponseAuth, error)
	// Kuryerlar ro'yxatini olish
	GetAllCouriers(context.Context, *GetAllCouriersRequest) (*GetAllCouriersResponse, error)
	// Kuryer ma'lumotlarini yangilash
	UpdateCourier(context.Context, *UpdateCourierRequestAuth) (*CourierResponseAuth, error)
	// Kuryer o'chirish
	DeleteCourier(context.Context, *CourierRequestAuth) (*CourierResponseAuth, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedAuthServiceServer) UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedAuthServiceServer) CourierLogin(context.Context, *CourierLoginRequest) (*CourierLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CourierLogin not implemented")
}
func (UnimplementedAuthServiceServer) CourierRegister(context.Context, *CourierRegisterRequest) (*CourierRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CourierRegister not implemented")
}
func (UnimplementedAuthServiceServer) CreateUser(context.Context, *CreateUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedAuthServiceServer) GetUser(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedAuthServiceServer) GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedAuthServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedAuthServiceServer) DeleteUser(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedAuthServiceServer) CreateCourier(context.Context, *CreateCourierRequestAuth) (*CourierResponseAuth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourier not implemented")
}
func (UnimplementedAuthServiceServer) GetCourier(context.Context, *CourierRequestAuth) (*CourierResponseAuth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourier not implemented")
}
func (UnimplementedAuthServiceServer) GetAllCouriers(context.Context, *GetAllCouriersRequest) (*GetAllCouriersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCouriers not implemented")
}
func (UnimplementedAuthServiceServer) UpdateCourier(context.Context, *UpdateCourierRequestAuth) (*CourierResponseAuth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCourier not implemented")
}
func (UnimplementedAuthServiceServer) DeleteCourier(context.Context, *CourierRequestAuth) (*CourierResponseAuth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCourier not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UserRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserRegister(ctx, req.(*UserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CourierLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourierLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CourierLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CourierLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CourierLogin(ctx, req.(*CourierLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CourierRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourierRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CourierRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CourierRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CourierRegister(ctx, req.(*CourierRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetAllUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetAllUsers(ctx, req.(*GetAllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeleteUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CreateCourier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourierRequestAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CreateCourier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CreateCourier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CreateCourier(ctx, req.(*CreateCourierRequestAuth))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetCourier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourierRequestAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetCourier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetCourier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetCourier(ctx, req.(*CourierRequestAuth))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetAllCouriers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCouriersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetAllCouriers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetAllCouriers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetAllCouriers(ctx, req.(*GetAllCouriersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateCourier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCourierRequestAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateCourier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UpdateCourier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateCourier(ctx, req.(*UpdateCourierRequestAuth))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeleteCourier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourierRequestAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeleteCourier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_DeleteCourier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeleteCourier(ctx, req.(*CourierRequestAuth))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _AuthService_UserLogin_Handler,
		},
		{
			MethodName: "UserRegister",
			Handler:    _AuthService_UserRegister_Handler,
		},
		{
			MethodName: "CourierLogin",
			Handler:    _AuthService_CourierLogin_Handler,
		},
		{
			MethodName: "CourierRegister",
			Handler:    _AuthService_CourierRegister_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _AuthService_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _AuthService_GetUser_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _AuthService_GetAllUsers_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _AuthService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _AuthService_DeleteUser_Handler,
		},
		{
			MethodName: "CreateCourier",
			Handler:    _AuthService_CreateCourier_Handler,
		},
		{
			MethodName: "GetCourier",
			Handler:    _AuthService_GetCourier_Handler,
		},
		{
			MethodName: "GetAllCouriers",
			Handler:    _AuthService_GetAllCouriers_Handler,
		},
		{
			MethodName: "UpdateCourier",
			Handler:    _AuthService_UpdateCourier_Handler,
		},
		{
			MethodName: "DeleteCourier",
			Handler:    _AuthService_DeleteCourier_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
