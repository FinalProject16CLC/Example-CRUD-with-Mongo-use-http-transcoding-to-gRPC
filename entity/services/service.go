package entity_services

import "go.mongodb.org/mongo-driver/mongo"

type EntityServiceServer struct {
	EntityCollection *mongo.Collection
}
