// +build int

package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/x64puzzle/spotted-common/config"
	l "github.com/x64puzzle/spotted-common/log"
	pb "github.com/x64puzzle/spotted-proto/auth"
	"google.golang.org/grpc"
)

var env = config.NewAuth()

func TestRegister(t *testing.T) {
	t.SkipNow()

	conn, err := grpc.Dial(":"+env.Port, grpc.WithInsecure())
	if err != nil {
		l.Log.Info("Failed to dial grpc: ", err)
	}
	defer conn.Close()

	client := pb.NewAuthClient(conn)

	resp, err := client.Register(context.Background(), &pb.RegisterRequest{
		Username: "semir",
		Email:    "semir@mail.com",
		Password: "pwd123",
	})
	if err != nil {
		l.Log.Error("Failed to call Register: ", err)
	}

	assert.NotNil(t, resp, "Response should not be nil")

	l.Log.Info("Register resp: ", resp)
}

func TestLogin(t *testing.T) {
	conn, err := grpc.Dial(":"+env.Port, grpc.WithInsecure())
	if err != nil {
		l.Log.Info("Failed to dial grpc: ", err)
	}
	defer conn.Close()

	client := pb.NewAuthClient(conn)

	resp, err := client.Login(context.Background(), &pb.LoginRequest{
		Email:    "semir@mail.com",
		Password: "pwd123",
	})
	if err != nil {
		l.Log.Error("Failed to call Login: ", err)
	}

	assert.NotNil(t, resp, "Response should not be nil")

	l.Log.Info("Login resp: ", resp)
}

func TestLogout(t *testing.T) {
	conn, err := grpc.Dial(":"+env.Port, grpc.WithInsecure())
	if err != nil {
		l.Log.Info("Failed to dial grpc: ", err)
	}
	defer conn.Close()

	client := pb.NewAuthClient(conn)

	resp, err := client.Logout(context.Background(), &pb.LogoutRequest{
		Email: "semir@mail.com",
	})
	if err != nil {
		l.Log.Error("Failed to call Logout: ", err)
	}

	assert.NotNil(t, resp, "Response should not be nil")

	l.Log.Info("Logout resp: ", resp)
}

func TestCreateResetToken(t *testing.T) {
	conn, err := grpc.Dial(":"+env.Port, grpc.WithInsecure())
	if err != nil {
		l.Log.Info("Failed to dial grpc: ", err)
	}
	defer conn.Close()

	client := pb.NewAuthClient(conn)

	resp, err := client.CreateResetToken(context.Background(), &pb.ResetTokenRequest{
		Email: "semir@mail.com",
	})
	if err != nil {
		l.Log.Error("Failed to call CreateResetToken: ", err)
	}

	assert.NotNil(t, resp, "Response should not be nil")
	assert.NotEmpty(t, resp.Token, "Token not generated")

	l.Log.Info("CreateResetToken resp: ", resp)
}

func TestDeleteResetToken(t *testing.T) {
	conn, err := grpc.Dial(":"+env.Port, grpc.WithInsecure())
	if err != nil {
		l.Log.Info("Failed to dial grpc: ", err)
	}
	defer conn.Close()

	client := pb.NewAuthClient(conn)

	resp, err := client.DeleteResetToken(context.Background(), &pb.ResetTokenRequest{
		Email: "semir@mail.com",
	})
	if err != nil {
		l.Log.Error("Failed to call DeleteResetToken: ", err)
	}

	assert.NotNil(t, resp, "Response should not be nil")
	assert.True(t, resp.Success, "Token not deleted")

	l.Log.Info("DeleteResetToken resp: ", resp)
}
