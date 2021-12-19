package client

import (
	"context"
	"log"
	"time"

	pb "servidor/gRPC/management"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func Export(user pb.User) bool {
	trigger := false
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatal("did not connect: %v", err)
		trigger = true
	}
	defer conn.Close()
	c := pb.NewUserManagmentClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.CreateNewUser(ctx, &user)
	if err != nil {
		log.Fatalf("could not create user: %v", err)
		trigger = true
	}
	log.Printf(`User Details:
	NAME: %s
	AGE: %d`, r.GetName(), r.GetAge())
	return trigger
}
