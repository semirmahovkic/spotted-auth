package service

import (
	"context"
	"errors"

	"github.com/x64integer/go-common/util"
	"github.com/x64integer/spotted-auth/dao"
	pb "github.com/x64integer/spotted-proto/auth"
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
	jwt := &util.Token{}

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
	// Find user by email and validate password
	userDao := &dao.User{}

	acc, err := userDao.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if valid := util.ValidPassword(acc.Password, req.Password); !valid {
		return nil, errors.New("invalid password")
	}

	// Generate jwt token
	jwt := &util.Token{}

	token, err := jwt.Generate(map[string]string{
		"username": acc.Username,
		"email":    acc.Email,
		"id":       acc.ID,
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.LoginResponse{}

	resp.Token = token

	// Create login session
	session := &dao.Session{}

	if err := session.Create(req.Email, resp.Token); err != nil {
		return nil, err
	}

	return resp, nil
}

// Logout implements spotted-proto/auth.Service.Logout
func (s *Server) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	session := &dao.Session{}

	if err := session.Destroy(req.Email); err != nil {
		return nil, err
	}

	resp := &pb.LogoutResponse{}
	resp.Success = true

	return resp, nil
}

// CreateResetToken implements spotted-proto/auth.Service.CreateResetToken
func (s *Server) CreateResetToken(ctx context.Context, req *pb.ResetTokenRequest) (*pb.ResetTokenResponse, error) {
	userDao := &dao.User{}

	resp, err := userDao.CreateResetToken(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteResetToken implements spotted-proto/auth.Service.DeleteResetToken
func (s *Server) DeleteResetToken(ctx context.Context, req *pb.ResetTokenRequest) (*pb.DeleteResetTokenResponse, error) {
	userDao := &dao.User{}

	resp, err := userDao.DeleteResetToken(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
