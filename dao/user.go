package dao

import "github.com/x64puzzle/spotted-proto/auth"

// User Data Access Object
type User struct{}

// Register user account
func (u *User) Register(req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	return nil, nil
}

// Login user account
func (u *User) Login(req *auth.LoginRequest) (*auth.LoginResponse, error) {
	return nil, nil
}

// Logout user account
func (u *User) Logout(req *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	return nil, nil
}

// PasswordReset for user account
func (u *User) PasswordReset(req *auth.PasswordResetRequest) (*auth.PasswordResetResponse, error) {
	return nil, nil
}
