package service

import (
	"net"

	"github.com/x64integer/spotted-common/config"
	l "github.com/x64integer/spotted-common/log"
	pb "github.com/x64integer/spotted-proto/auth"
	"google.golang.org/grpc"
)

var conf = config.NewAuth()

// ListenGRPC server
func ListenGRPC() {
	listener, err := net.Listen("tcp", ":"+conf.Port)
	if err != nil {
		l.Log.Fatal("failed to create tcp listener: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, &Server{})

	if err := grpcServer.Serve(listener); err != nil {
		l.Log.Fatal("failed to serve grpc server: ", err)
	}
}
