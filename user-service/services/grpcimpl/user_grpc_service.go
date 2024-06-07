package grpcserviceimpl

import (
	"context"

	proto "github.com/mugnialby/go-microservices/user-service/proto/user"
	"github.com/mugnialby/go-microservices/user-service/repositories"
)

type userServiceServer struct {
	proto.UnimplementedUserServiceServer
	userRepository repositories.UserRepository
}

func NewUserServiceServer(userRepository repositories.UserRepository) proto.UserServiceServer {
	return &userServiceServer{
		userRepository: userRepository,
	}
}

func (userServiceServer *userServiceServer) FindByUsername(context.Context, *proto.FindByUsernameRequest) (*proto.FindByUsernameResponse, error) {
	return nil, nil
}
