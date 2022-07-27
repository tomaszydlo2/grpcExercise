package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpcExercise/internal/db"
	"grpcExercise/internal/serverdb"
	"grpcExercise/internal/users"
	"log"
	"net"
)

var (
	port = flag.Int("port", 8001, "The server port")
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	users.RegisterUsersServer(grpcServer, &serverdb.Server{Database: &db.Db{}})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
