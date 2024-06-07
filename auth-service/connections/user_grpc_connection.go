package connections

import (
	"log"

	proto "github.com/mugnialby/go-microservices/auth-service/proto/users"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserConnection() proto.UserServiceClient {
	var clientConnection *grpc.ClientConn
	clientConnection, err := grpc.Dial(viper.GetString("services.users.port"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to users service : %v", err)
	}

	return proto.NewUserServiceClient(clientConnection)
}
