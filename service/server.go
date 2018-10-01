package service

import (
	"context"

	"github.com/x64puzzle/spotted-auth/dao"
	"github.com/x64puzzle/spotted-proto/auth"
)

// Server implements spotted-proto/auth.Service
type Server struct{}

// Register implements spotted-proto/auth.Service.Register
func (s *Server) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	userDao := &dao.User{}

	resp, err := userDao.Register(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Login implements spotted-proto/auth.Service.Login
func (s *Server) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	userDao := &dao.User{}

	resp, err := userDao.Login(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Logout implements spotted-proto/auth.Service.Logout
func (s *Server) Logout(ctx context.Context, req *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	userDao := &dao.User{}

	resp, err := userDao.Logout(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// PasswordReset implements spotted-proto/auth.Service.PasswordReset
func (s *Server) PasswordReset(ctx context.Context, req *auth.PasswordResetRequest) (*auth.PasswordResetResponse, error) {
	userDao := &dao.User{}

	resp, err := userDao.PasswordReset(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
