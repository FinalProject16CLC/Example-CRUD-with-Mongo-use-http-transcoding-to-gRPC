package user_services

import "go.mongodb.org/mongo-driver/mongo"

type UserServiceServer struct {
	UserCollection *mongo.Collection
}
