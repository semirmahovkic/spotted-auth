package service

import (
	"net"

	"github.com/x64puzzle/spotted-common/config"
	l "github.com/x64puzzle/spotted-common/log"
	pb "github.com/x64puzzle/spotted-proto/auth"
	"google.golang.org/grpc"
)

var env = config.NewAuth()

// ListenGRPC server
func ListenGRPC() {
	listener, err := net.Listen("tcp", ":"+env.Port)
	if err != nil {
		l.Log.Fatal("Failed to create tcp listener: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, &Server{})

	if err := grpcServer.Serve(listener); err != nil {
		l.Log.Fatal("Failed to serve grpc server: ", err)
	}
}
