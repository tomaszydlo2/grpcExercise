package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpcExercise/internal/users"
	"log"
	"time"
)

var (
	addr = flag.String("addr", "localhost:8001", "the address to connect to")
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	uClient := users.NewUsersClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Testing the server:
	ClientCreateUser(uClient, ctx, &users.User{
		Id:       1,
		Username: "JohnyT",
		Name:     "John",
		Surname:  "Testinger",
	})
	ClientCreateUser(uClient, ctx, &users.User{
		Id:       2,
		Username: "AnthonyK",
		Name:     "Anton",
		Surname:  "Kowalski",
	})

	ClientReadUser(uClient, ctx, &users.Id{Id: 1})

	ClientReadAllUsers(uClient, ctx)

	ClientDeleteUser(uClient, ctx, &users.Id{Id: 1})

	ClientUpdateUser(uClient, ctx, &users.User{
		Id:       2,
		Username: "TomaszClient",
		Name:     "Tomasz",
		Surname:  "Testing",
	})

}

func ClientCreateUser(uClient users.UsersClient, ctx context.Context, user *users.User) {
	response, err := uClient.CreateUser(ctx, user)
	if err != nil {
		log.Fatalf("Error when creating user: %s", err)
	}
	log.Printf("Creating a user: %s", response.Message)
}

func ClientReadUser(uClient users.UsersClient, ctx context.Context, id *users.Id) {
	response, err := uClient.ReadUser(ctx, id)
	if err != nil {
		log.Fatalf("Error when reading user: %s", err)
	}
	log.Printf("Reading one user: %s", response.Message)
}

func ClientReadAllUsers(uClient users.UsersClient, ctx context.Context) {
	response, err := uClient.ReadUsers(ctx, &users.Empty{})
	if err != nil {
		log.Fatalf("Error when reading users: %s", err)
	}
	log.Printf("Reading all users: %s", response.Message)
}

func ClientDeleteUser(uClient users.UsersClient, ctx context.Context, id *users.Id) {
	response, err := uClient.DeleteUser(ctx, id)
	if err != nil {
		log.Fatalf("Error when deleting user: %s", err)
	}
	log.Printf("Deleting an user: %s", response.Message)
}

func ClientUpdateUser(uClient users.UsersClient, ctx context.Context, user *users.User) {
	response, err := uClient.UpdateUser(ctx, user)
	if err != nil {
		log.Fatalf("Error when updating user: %s", err)
	}
	log.Printf("Updating an user: %s", response.Message)
}
