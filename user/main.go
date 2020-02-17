package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/golang/glog"
	"github.com/joho/godotenv"
	user_pb "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/protos/user"
	user_models "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/user/models"
	user_services "github.com/trinhdaiphuc/Example-CRUD-with-Mongo-use-http-transcoding-to-gRPC/user/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	flag.Parse()
	defer glog.Flush()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port :50052...")

	// Start our listener, 50051 is the default gRPC port
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Unable to listen on port :50052: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	srv := &user_services.UserServiceServer{}
	user_pb.RegisterUserServiceServer(s, srv)

	// Initialize MongoDb client
	fmt.Println("Connecting to MongoDB...")

	mongoctx := context.Background()
	db, err := mongo.Connect(mongoctx, options.Client().ApplyURI(os.Getenv("DB_HOST")))
	fmt.Println("DB_HOST ", os.Getenv("DB_HOST"))

	if err != nil {
		log.Fatal(err)
		glog.Fatal(err)
	}

	err = db.Ping(mongoctx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to Mongodb")
	}
	srv.UserCollection = user_models.NewUserCollection(db)
	// Start the server in a child routine
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
			glog.Fatal(err)
		}
	}()
	fmt.Println("Server succesfully started on port :50051")
	// Right way to stop the server using a SHUTDOWN HOOK
	// Create a channel to receive OS signals
	c := make(chan os.Signal)

	// Relay os.Interrupt to our channel (os.Interrupt = CTRL+C)
	// Ignore other incoming signals
	signal.Notify(c, os.Interrupt)

	// Block main routine until a signal is received
	// As long as user doesn't press CTRL+C a message is not passed and our main routine keeps running
	<-c

	// After receiving CTRL+C Properly stop the server
	fmt.Println("\nStopping the server...")
	s.Stop()
	listener.Close()
	fmt.Println("Closing MongoDB connection")
	db.Disconnect(mongoctx)
	fmt.Println("Done.")
}
