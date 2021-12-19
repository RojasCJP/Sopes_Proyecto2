package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "servidor/gRPC/management"
	"servidor/structs"

	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagmentServer
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	insertUser := pb.User{Name: in.Name, Age: in.Age, VaccineType: in.VaccineType, Location: in.Location, NDose: in.NDose}
	client, context := Connection()
	InsertMongo(client, context, insertUser)
	return &insertUser, nil
}

func Export() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagmentServer(s, &UserManagementServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func Connection() (*mongo.Client, context.Context) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	/*
		List databases
	*/
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	return client, ctx
}

func InsertMongo(client *mongo.Client, ctx context.Context, data pb.User) bool {
	trigger := false
	client, ctx = Connection()

	collection := client.Database("SopesProyecto2").Collection("datos")

	/*
	  Insert documents
	*/
	var docs []interface{}
	docs = append(docs, bson.D{{"name", data.Name}, {"location", data.Location}, {"age", data.Age}, {"vaccine_type", data.VaccineType}, {"n_dose", data.NDose}})

	res, insertErr := collection.InsertMany(ctx, docs)
	if insertErr != nil {
		trigger = true
		log.Fatal("insert error", insertErr)
	}
	fmt.Println(res)
	/*
		Iterate a cursor and print it
	*/
	cur, currErr := collection.Find(ctx, bson.D{})

	if currErr != nil {
		trigger = true
		log.Fatal("find error", currErr)
	}
	defer cur.Close(ctx)

	var posts []structs.Prueba
	if err := cur.All(ctx, &posts); err != nil {
		trigger = true
		log.Fatal("cur all", err)
	}

	defer client.Disconnect(ctx)
	return trigger
}
