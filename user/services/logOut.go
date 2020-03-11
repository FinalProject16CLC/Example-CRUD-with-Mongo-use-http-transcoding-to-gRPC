package user_services

import (
	"context"
	"fmt"
	"log"

	user_pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos/user"
	user_models "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/user/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

func (s *UserServiceServer) LogOut(ctx context.Context, req *user_pb.LogOutReq) (*user_pb.LogOutRes, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetUserId().GetId())

	result := s.UserCollection.FindOne(ctx, bson.M{"_id": oid})
	user := &user_models.UserItem{}
	if err = result.Decode(&user); err != nil {
		log.Println("Error when log out", err)
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(
				codes.Unauthenticated,
				fmt.Sprintf("Not found User {error: {%v}}", err),
			)
		}
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	return &user_pb.LogOutRes{Success: true}, nil
}
