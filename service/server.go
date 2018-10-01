package service

import (
	"context"

	"github.com/x64puzzle/spotted-auth/dao"
	"github.com/x64puzzle/spotted-common/util"
	pb "github.com/x64puzzle/spotted-proto/auth"
)

// Server implements spotted-proto/auth.Service
type Server struct{}

// Register implements spotted-proto/auth.Service.Register
func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	userDao := &dao.User{}

	// Register user in db
	resp, err := userDao.Register(req)
	if err != nil {
		return nil, err
	}

	// Generate jwt token for user
	jwt := util.NewToken()

	token, err := jwt.Generate(map[string]string{
		"username": req.Username,
		"email":    req.Email,
		"id":       req.ID,
	})
	if err != nil {
		return nil, err
	}

	resp.Token = token

	// Create login session
	session := &dao.Session{}

	if err := session.Create(req.Email, resp.Token); err != nil {
		return nil, err
	}

	return resp, nil
}

// Login implements spotted-proto/auth.Service.Login
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	userDao := &dao.User{}

	resp, err := userDao.Login(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Logout implements spotted-proto/auth.Service.Logout
func (s *Server) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	userDao := &dao.User{}

	resp, err := userDao.Logout(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// PasswordReset implements spotted-proto/auth.Service.PasswordReset
func (s *Server) PasswordReset(ctx context.Context, req *pb.PasswordResetRequest) (*pb.PasswordResetResponse, error) {
	userDao := &dao.User{}

	resp, err := userDao.PasswordReset(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
