// @/sub/main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	redis "github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const urlRedis = "localhost:6379"
const urlMongo = "mongodb://localhost:27017/"

type User struct {
	Name         string `json:"name"`
	Location     string `json:"location"`
	Age          int    `json:"age"`
	Vaccine_type string `json:"vaccine_type"`
	N_dose       int    `json:"n_dose"`
}

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr:     urlRedis,
	Password: "", // no password set
	DB:       0,  // use default DB
})

func main() {
	subscriber := redisClient.Subscribe(ctx, "send-user-data")

	user := User{}

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal([]byte(msg.Payload), &user); err != nil {
			panic(err)
		}

		client, context := Connection()
		usuario := PbUser{Name: user.Name, Age: int32(user.Age), VaccineType: user.Vaccine_type, Location: user.Location, NDose: int32(user.N_dose)}
		InsertMongo(client, context, usuario)

		fmt.Println("Received message from " + msg.Channel + " channel.")
		fmt.Printf("%+v\n", user)
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

func InsertMongo(client *mongo.Client, ctx context.Context, data PbUser) bool {
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

type PbUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age         int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	VaccineType string `protobuf:"bytes,3,opt,name=vaccine_type,json=vaccineType,proto3" json:"vaccine_type,omitempty"`
	Location    string `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	NDose       int32  `protobuf:"varint,5,opt,name=n_dose,json=nDose,proto3" json:"n_dose,omitempty"`
}
