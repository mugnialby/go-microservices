package server

import (
	"log"
	"net"

	"github.com/mugnialby/go-microservices/user-service/proto/users"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func StartGrpcServer(
	userServiceServer users.UserServiceServer,
) {
	grpcServer := grpc.NewServer()

	users.RegisterUserServiceServer(grpcServer, userServiceServer)

	listener, err := net.Listen("tcp", viper.GetString("services.users.port"))
	if err != nil {
		logrus.Fatalf("Failed to listen : %v", err)
	}

	log.Printf("Starting gRPC server on port %v", viper.GetString("services.users.port"))
	if err := grpcServer.Serve(listener); err != nil {
		logrus.Fatalf("Failed to serve : %v", err)
	}
}
