package dao

import (
	"database/sql"

	"github.com/x64integer/go-common/util"
	"github.com/x64integer/spotted-auth/storage"
	pb "github.com/x64integer/spotted-proto/auth"
	pbu "github.com/x64integer/spotted-proto/user"
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

	_, err = storage.DB.Query("SELECT create_user($1, $2, $3, $4);", req.ID, req.Username, req.Email, pwd)
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

	if err := storage.DB.QueryRow("SELECT * FROM get_by_email($1);", email).Scan(&acc.ID, &acc.Username, &acc.Email, &acc.Password, &acc.CreatedAt); err != nil {
		return nil, err
	}

	return acc, nil
}

// CreateResetToken for user account
func (u *User) CreateResetToken(req *pb.ResetTokenRequest) (*pb.ResetTokenResponse, error) {
	token := util.RandomStr(64)

	existingToken, err := u.getResetToken(req.Email)
	if err != nil {
		return nil, err
	}

	if existingToken == "" {
		_, err := storage.DB.Query("SELECT create_reset_token($1, $2);", req.Email, token)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := storage.DB.Query("SELECT update_reset_token($1, $2);", req.Email, token)
		if err != nil {
			return nil, err
		}
	}

	resp := &pb.ResetTokenResponse{}
	resp.Token = token

	return resp, nil
}

// DeleteResetToken for user account
func (u *User) DeleteResetToken(req *pb.ResetTokenRequest) (*pb.DeleteResetTokenResponse, error) {
	_, err := storage.DB.Query("SELECT delete_reset_token($1);", req.Email)
	if err != nil {
		return nil, err
	}

	resp := &pb.DeleteResetTokenResponse{}
	resp.Success = true

	return resp, nil
}

// getResetToken by email
func (u *User) getResetToken(email string) (string, error) {
	var token sql.NullString

	if err := storage.DB.QueryRow("SELECT * FROM get_reset_token($1);", email).Scan(&token); err != nil {
		return "", err
	}

	return token.String, nil
}
