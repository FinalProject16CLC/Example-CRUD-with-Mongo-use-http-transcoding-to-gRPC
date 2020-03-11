package user_models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type UserItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserName string             `bson:"user_name"`
	Password string             `bson:"pasword"`
}

func NewUserCollection(db *mongo.Client) (userCollection *mongo.Collection) {
	// Create indexs
	mod := mongo.IndexModel{
		Keys: bson.M{
			"user_name": -1, // index in ascending order
		},
		// create UniqueIndex option
		Options: options.Index().SetUnique(true),
	}
	userCollection = db.Database("mydb").Collection("user")
	userCollection.Indexes().CreateOne(context.Background(), mod)
	return
}
