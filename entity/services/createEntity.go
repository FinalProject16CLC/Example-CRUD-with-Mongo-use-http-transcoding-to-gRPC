package entity_services

import (
	"context"
	"fmt"

	entity_pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/entity/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateEntity is a gRPC function to create an entity in MongoDB
func (s *EntityServiceServer) CreateEntity(ctx context.Context, req *entity_pb.CreateEntityReq) (*entity_pb.CreateEntityRes, error) {
	// Essentially doing req.Entity to access the struct with a nil check
	entity := req.GetEntity()
	// Now we have to convert this into a EtityItem type to convert into BSON
	data := &entity_models.EntityItem{
		// ID:    Empty, so it gets omitted and MongoDB generates a unique Object ID upon insertion.
		Name:        entity.GetName(),
		Description: entity.GetDescription(),
		URL:         entity.GetUrl(),
	}

	// Insert the data into the database, result contains the newly generated Object ID for the new document
	result, err := s.EntityCollection.InsertOne(ctx, data)
	// check for potential errors
	if err != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	// add the id to entity, first cast the "generic type" (go doesn't have real generics yet) to an Object ID.
	oid := result.InsertedID.(primitive.ObjectID)
	// Convert the object id to it's string counterpart
	entity.Id = oid.Hex()
	// return the entity in a CreateEntityRes type
	return &entity_pb.CreateEntityRes{Entity: entity}, nil
}