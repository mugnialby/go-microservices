package server

import (
	"log"
	"net"

	"github.com/mugnialby/go-microservices/user-service/proto/user"
	"google.golang.org/grpc"
)

func StartGrpcServer(
	userServiceServer user.UserServiceServer,
) {
	grpcServer := grpc.NewServer()

	user.RegisterUserServiceServer(grpcServer, userServiceServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	log.Println("Starting gRPC server on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
