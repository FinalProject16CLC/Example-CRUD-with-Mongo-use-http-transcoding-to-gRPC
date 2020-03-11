package user_services

import (
	"context"
	"fmt"

	user_pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos/user"
	user_models "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/user/models"
	"github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/user/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

func (s *UserServiceServer) SignUp(ctx context.Context, req *user_pb.SignUpReq) (*user_pb.SignUpRes, error) {
	user := req.GetUser()
	result := s.UserCollection.FindOne(ctx, bson.M{"user_name": user.GetUserName()})
	data := &user_models.UserItem{}
	err := result.Decode(&data)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal error: %v", err),
			)
		}
	}

	if data.UserName != "" {
		return nil, status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf("This email already existed."),
		)
	}

	data.UserName = user.GetUserName()
	data.Password, err = utils.HashPassword(user.GetPassword())

	if err != nil {
		return nil, status.Errorf(
			codes.Internal, fmt.Sprintf("Internal error: %v", err),
		)
	}

	// Save user
	resultInsert, err := s.UserCollection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal, fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid := resultInsert.InsertedID.(primitive.ObjectID)
	id := oid.Hex()

	return &user_pb.SignUpRes{Id: id, Success: true}, nil
}
