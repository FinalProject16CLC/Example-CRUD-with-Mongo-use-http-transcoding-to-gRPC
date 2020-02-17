package user_services

import (
	"context"

	user_pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos/user"
	user_models "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/user/models"
)

func (s *UserServiceServer) LogIn(ctx context.Context, req *user_pb.LogInReq) (*user_pb.LogInRes, error) {
	user := req.GetUser()

	data := &user_models.UserItem{
		UserName: user.GetUserName(),
		Password: user.GetPassword(),
	}

	result := s.UserCollection.FindOne(context.Background(), bson.M{"user_name": data.UserName})
	user = user_models.User{}
	if err := result.Decode(&user); err != nil {
		h.AppLog.Info("Error when sign in by email ", err)
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(
				codes.Unauthorized, 
				mt.Sprintf("Internal error: %v", err),
			)
		}
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	return &user_pb.LogInRes{User: user}, nil
}
