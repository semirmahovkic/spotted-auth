package dao

import (
	l "github.com/x64puzzle/spotted-common/log"
	"github.com/x64puzzle/spotted-common/util"
	pb "github.com/x64puzzle/spotted-proto/auth"
)

// User Data Access Object
type User struct{}

// Register user account
func (u *User) Register(req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	l.Log.Info("Register req: ", req)

	uuid := util.UUID()

	req.ID = uuid

	// TODO: Insert into db

	resp := &pb.RegisterResponse{}

	return resp, nil
}

// Login user account
func (u *User) Login(req *pb.LoginRequest) (*pb.LoginResponse, error) {
	l.Log.Info("Login req: ", req)

	return &pb.LoginResponse{}, nil
}

// Logout user account
func (u *User) Logout(req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	l.Log.Info("Logout req: ", req)

	return &pb.LogoutResponse{}, nil
}

// PasswordReset for user account
func (u *User) PasswordReset(req *pb.PasswordResetRequest) (*pb.PasswordResetResponse, error) {
	l.Log.Info("PasswordReset req: ", req)

	return &pb.PasswordResetResponse{}, nil
}
