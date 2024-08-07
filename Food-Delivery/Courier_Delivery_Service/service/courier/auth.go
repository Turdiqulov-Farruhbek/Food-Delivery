package courier

import (
	"context"
	"courier_delivery/config/logger"
	"courier_delivery/genproto/user"
	stg "courier_delivery/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServiceServer struktura
type AuthServiceServer struct {
	UserSvc stg.StorageCourierInterface
	user.UnimplementedAuthServiceServer
	log logger.Logger
}

// NewAuthService yangi UserServiceServer yaratish uchun funksiya
func NewAuthService(userSvc stg.StorageCourierInterface, log logger.Logger) *AuthServiceServer {
	return &AuthServiceServer{UserSvc: userSvc, log: log}
}

func (s *AuthServiceServer) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	res, err := s.UserSvc.Auth().UserRegister(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("failed to register user: %v", err)

		return nil, err
	}
	return res, nil
}

func (s *AuthServiceServer) UserLogin(ctx context.Context, req *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	res, err := s.UserSvc.Auth().UserLogin(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to register user: %v", err)
	}
	return res, nil
}

// CreateUser RPC
func (s *AuthServiceServer) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.UserResponse, error) {
	res, err := s.UserSvc.Auth().CreateUser(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}
	return res, nil
}

// GetUser RPC
func (s *AuthServiceServer) GetUser(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error) {
	res, err := s.UserSvc.Auth().GetUser(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	return res, nil
}

// UpdateUser RPC
func (s *AuthServiceServer) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UserResponse, error) {
	res, err := s.UserSvc.Auth().UpdateUser(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update user: %v", err)
	}
	return res, nil
}

// DeleteUser RPC
func (s *AuthServiceServer) DeleteUser(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error) {
	res, err := s.UserSvc.Auth().DeleteUser(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete user: %v", err)
	}
	return res, nil
}

// GetAllUsers RPC
func (s *AuthServiceServer) GetAllUsers(ctx context.Context, req *user.GetAllUsersRequest) (*user.GetAllUsersResponse, error) {
	res, err := s.UserSvc.Auth().GetAllUsers(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get all users: %v", err)
	}
	return res, nil
}

func (s *AuthServiceServer) CourierRegister(ctx context.Context, req *user.CourierRegisterRequest) (*user.CourierRegisterResponse, error) {
	res, err := s.UserSvc.Auth().CourierRegister(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to register user: %v", err)
	}
	return res, nil
}

func (s *AuthServiceServer) CourierLogin(ctx context.Context, req *user.CourierLoginRequest) (*user.CourierLoginResponse, error) {
	res, err := s.UserSvc.Auth().CourierLogin(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to register user: %v", err)
	}
	return res, nil
}

// CreateCourier RPC
func (s *AuthServiceServer) CreateCourier(ctx context.Context, req *user.CreateCourierRequestAuth) (*user.CourierResponseAuth, error) {
	res, err := s.UserSvc.Auth().CreateCourier(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create courier: %v", err)
	}
	return res, nil
}

// GetCourier RPC
func (s *AuthServiceServer) GetCourier(ctx context.Context, req *user.CourierRequestAuth) (*user.CourierResponseAuth, error) {
	res, err := s.UserSvc.Auth().GetCourier(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get courier: %v", err)
	}
	return res, nil
}

// UpdateCourier RPC
func (s *AuthServiceServer) UpdateCourier(ctx context.Context, req *user.UpdateCourierRequestAuth) (*user.CourierResponseAuth, error) {
	res, err := s.UserSvc.Auth().UpdateCourier(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update courier: %v", err)
	}
	return res, nil
}

// DeleteCourier RPC
func (s *AuthServiceServer) DeleteCourier(ctx context.Context, req *user.CourierRequestAuth) (*user.CourierResponseAuth, error) {
	res, err := s.UserSvc.Auth().DeleteCourier(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete courier: %v", err)
	}
	return res, nil
}

// GetAllCouriers RPC
func (s *AuthServiceServer) GetAllCouriers(ctx context.Context, req *user.GetAllCouriersRequest) (*user.GetAllCouriersResponse, error) {
	res, err := s.UserSvc.Auth().GetAllCouriers(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get all couriers: %v", err)
	}
	return res, nil
}
