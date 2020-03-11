package user_services

import (
	"context"
	"fmt"
	"log"

	user_pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos/user"
	user_models "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/user/models"
	"github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/user/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

func (s *UserServiceServer) LogIn(ctx context.Context, req *user_pb.LogInReq) (*user_pb.LogInRes, error) {
	user := req.GetUser()

	result := s.UserCollection.FindOne(ctx, bson.M{"user_name": user.GetUserName()})
	data := &user_models.UserItem{}
	if err := result.Decode(&data); err != nil {
		log.Println("Error when log in", err)
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(
				codes.Unauthenticated,
				fmt.Sprintf("Invalid email or password."),
			)
		}
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	if ok := utils.CheckPasswordHash(user.Password, data.Password); !ok {
		return nil, status.Errorf(
			codes.Unauthenticated,
			fmt.Sprintf("Password is invalid."),
		)
	}

	id := data.ID.Hex()

	return &user_pb.LogInRes{Id: id, Success: true}, nil
}
