package courier

import (
	"context"
	"courier_delivery/genproto/user"
	stg "courier_delivery/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserServiceServer struktura
type UserServiceServer struct {
	UserSvc stg.AuthServiceInterface
	user.UnimplementedAuthServiceServer
}

// NewUserServiceServer yangi UserServiceServer yaratish uchun funksiya
func NewUserServiceServer(userSvc stg.AuthServiceInterface) *UserServiceServer {
	return &UserServiceServer{UserSvc: userSvc}
}

func (s *UserServiceServer) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	res, err := s.UserSvc.UserRegister(ctx, req)
    if err!= nil {
        return nil, status.Errorf(codes.Internal, "failed to register user: %v", err)
    }
    return res, nil
}

func (s *UserServiceServer) UserLogin(ctx context.Context, req *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	res, err := s.UserSvc.UserLogin(ctx, req)
    if err!= nil {
        return nil, status.Errorf(codes.Internal, "failed to register user: %v", err)
    }
    return res, nil
}

// CreateUser RPC
func (s *UserServiceServer) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.UserResponse, error) {
	res, err := s.UserSvc.CreateUser(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}
	return res, nil
}

// GetUser RPC
func (s *UserServiceServer) GetUser(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error) {
	res, err := s.UserSvc.GetUser(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	return res, nil
}

// UpdateUser RPC
func (s *UserServiceServer) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UserResponse, error) {
	res, err := s.UserSvc.UpdateUser(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update user: %v", err)
	}
	return res, nil
}

// DeleteUser RPC
func (s *UserServiceServer) DeleteUser(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error) {
	res, err := s.UserSvc.DeleteUser(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete user: %v", err)
	}
	return res, nil
}

// GetAllUsers RPC
func (s *UserServiceServer) GetAllUsers(ctx context.Context, req *user.GetAllUsersRequest) (*user.GetAllUsersResponse, error) {
	res, err := s.UserSvc.GetAllUsers(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get all users: %v", err)
	}
	return res, nil
}

func (s *UserServiceServer) CourierRegister(ctx context.Context, req *user.CourierRegisterRequest) (*user.CourierRegisterResponse, error) {
	res, err := s.UserSvc.CourierRegister(ctx, req)
    if err!= nil {
        return nil, status.Errorf(codes.Internal, "failed to register user: %v", err)
    }
    return res, nil
}


func (s *UserServiceServer) CourierLogin(ctx context.Context, req *user.CourierLoginRequest) (*user.CourierLoginResponse, error) {
	res, err := s.UserSvc.CourierLogin(ctx, req)
    if err!= nil {
        return nil, status.Errorf(codes.Internal, "failed to register user: %v", err)
    }
    return res, nil
}

// CreateCourier RPC
func (s *UserServiceServer) CreateCourier(ctx context.Context, req *user.CreateCourierRequestAuth) (*user.CourierResponseAuth, error) {
	res, err := s.UserSvc.CreateCourier(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create courier: %v", err)
	}
	return res, nil
}

// GetCourier RPC
func (s *UserServiceServer) GetCourier(ctx context.Context, req *user.CourierRequestAuth) (*user.CourierResponseAuth, error) {
	res, err := s.UserSvc.GetCourier(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get courier: %v", err)
	}
	return res, nil
}

// UpdateCourier RPC
func (s *UserServiceServer) UpdateCourier(ctx context.Context, req *user.UpdateCourierRequestAuth) (*user.CourierResponseAuth, error) {
	res, err := s.UserSvc.UpdateCourier(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update courier: %v", err)
	}
	return res, nil
}

// DeleteCourier RPC
func (s *UserServiceServer) DeleteCourier(ctx context.Context, req *user.CourierRequestAuth) (*user.CourierResponseAuth, error) {
	res, err := s.UserSvc.DeleteCourier(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete courier: %v", err)
	}
	return res, nil
}

// GetAllCouriers RPC
func (s *UserServiceServer) GetAllCouriers(ctx context.Context, req *user.GetAllCouriersRequest) (*user.GetAllCouriersResponse, error) {
	res, err := s.UserSvc.GetAllCouriers(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get all couriers: %v", err)
	}
	return res, nil
}
