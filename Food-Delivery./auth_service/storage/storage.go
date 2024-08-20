package storage

import (
	pb "delivery/auth_service/genproto"
)

type StorageI interface {
	User() UserI
}
type UserI interface {
	RegisterUser(req *pb.UserCreateReq) (*pb.Void, error)
	RegisterCourier(req *pb.UserCreateReq) (*pb.Void, error)
	Login(req *pb.LoginReq) (*pb.Token, error)
	RefreshToken(req *pb.Token) (*pb.Token, error)
	UpdateProfile(req *pb.UserUpdate) (*pb.Void, error)
	GetUserProfile(id *pb.ById) (*pb.UserCreateRes, error)
	ChangePassword(req *pb.PasswordChangeReq) (*pb.Void, error)
	ForgotPassword(req *pb.ForgotPasswordReq) (*pb.Void, error)
	ResetPassword(req *pb.PasswordResetReq) (*pb.Void, error)
	ConfirmEmail(req *pb.EmailConfirm) (*pb.Void, error)
	ResendCode(req *pb.ResendReq) (*pb.Void, error)
	CreateManager(req *pb.UserCreateReq) (*pb.Void, error)
	GetAllUsers(req *pb.UserFilter) (*pb.UserList, error)
}
