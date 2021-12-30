package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	pb "server/management"

	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-redis/redis/v8"
)

const urlRedis = "34.135.96.5:6379"
const urlMongo = "mongodb://34.135.96.5:27017/"

const (
	port = ":80"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagmentServer
}

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr:     urlRedis,
	Password: "rojas", // no password set
	DB:       0,       // use default DB
})

func main() {
	Export()
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	insertUser := pb.User{Name: in.Name, Age: in.Age, VaccineType: in.VaccineType, Location: in.Location, NDose: in.NDose}
	client, context := Connection()
	InsertMongo(client, context, insertUser)
	InsertRedis(insertUser)
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
	client, err := mongo.NewClient(options.Client().ApplyURI(urlMongo))
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

	defer client.Disconnect(ctx)
	return trigger
}

func InsertRedis(user pb.User) bool {
	trigger := false

	payload, err := json.Marshal(user)
	if err != nil {
		trigger = true
		panic(err)
	}

	// Posting through the chanel
	if err := redisClient.Publish(ctx, "send-user-data", payload).Err(); err != nil {
		trigger = true
		panic(err)
	}

	// Verifying age range and writing to redis - database
	if user.Age >= 0 && user.Age <= 11 {
		redisClient.Incr(ctx, "range0_11")
	} else if user.Age >= 12 && user.Age <= 18 {
		redisClient.Incr(ctx, "range12_18")
	} else if user.Age >= 19 && user.Age <= 26 {
		redisClient.Incr(ctx, "range19_26")
	} else if user.Age >= 27 && user.Age <= 59 {
		redisClient.Incr(ctx, "range27_59")
	} else {
		redisClient.Incr(ctx, "range60_end")
	}

	// Save name to report the last five vaccinated
	redisClient.LPush(ctx, "users", user.Name)
	return trigger
}
