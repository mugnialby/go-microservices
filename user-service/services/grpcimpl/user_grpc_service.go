package grpcserviceimpl

import (
	"context"
	"log"

	proto "github.com/mugnialby/go-microservices/user-service/proto/users"
	"github.com/mugnialby/go-microservices/user-service/services"
)

type userGrpcService struct {
	proto.UnimplementedUserServiceServer
	userService services.UserService
}

func NewUserServiceServer(userService services.UserService) proto.UserServiceServer {
	return &userGrpcService{
		userService: userService,
	}
}

func (userGrpcService *userGrpcService) FindByUsername(
	context context.Context,
	findByUsernameRequest *proto.FindByUsernameRequest,
) (*proto.FindByUsernameResponse, error) {
	user, err := userGrpcService.userService.FindByUsername(findByUsernameRequest.Username)
	if err != nil {
		log.Println(err.Error())
		return &proto.FindByUsernameResponse{}, nil
	}

	findByUsernameResponse := proto.FindByUsernameResponse{
		Email:    user.Email,
		Password: user.Password,
	}

	return &findByUsernameResponse, err
}
