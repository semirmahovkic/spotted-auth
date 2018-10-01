package dao

import (
	"github.com/x64puzzle/spotted-common/storage"
	"github.com/x64puzzle/spotted-common/util"
	pb "github.com/x64puzzle/spotted-proto/auth"
	pbu "github.com/x64puzzle/spotted-proto/user"
)

// User Data Access Object
type User struct{}

// Register user account
func (u *User) Register(req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	uuid := util.UUID()

	req.ID = uuid
	pwd, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	_, err = storage.PG.Query("SELECT create_user($1, $2, $3, $4);", req.ID, req.Username, req.Email, pwd)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{}, nil
}

// GetByEmail user account info
// NOTE: Use with caution, because password will be returned too!
// This func should be used for validation only!
func (u *User) GetByEmail(email string) (*pbu.Account, error) {
	acc := &pbu.Account{}

	if err := storage.PG.QueryRow("SELECT * FROM get_user_by_email($1);", email).Scan(&acc.ID, &acc.Username, &acc.Email, &acc.Password, &acc.CreatedAt); err != nil {
		return nil, err
	}

	return acc, nil
}

// PasswordReset for user account
func (u *User) PasswordReset(req *pb.PasswordResetRequest) (*pb.PasswordResetResponse, error) {
	return &pb.PasswordResetResponse{}, nil
}
