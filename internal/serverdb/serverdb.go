package serverdb

import (
	"context"
	"grpcExercise/internal/db"
	"grpcExercise/internal/users"
)

type Server struct {
	users.UnimplementedUsersServer
	Database db.Storage
}

func (s *Server) CreateUser(ctx context.Context, in *users.User) (*users.UserResponse, error) {
	s.Database.CreateUser(in)
	return &users.UserResponse{Message: "Created user " + in.GetUsername()}, nil
}

func (s *Server) UpdateUser(ctx context.Context, in *users.User) (*users.UserResponse, error) {
	if s.Database.UpdateUser(in) {
		return &users.UserResponse{Message: "Successfully updated user " + in.GetUsername()}, nil
	}
	return &users.UserResponse{Message: "Couldn't find user of id " + string(in.GetId())}, nil
}

func (s *Server) DeleteUser(ctx context.Context, in *users.Id) (*users.UserResponse, error) {
	if s.Database.DeleteUser(in) {
		return &users.UserResponse{Message: "Successfully removed user " + string(in.GetId())}, nil
	}
	return &users.UserResponse{Message: "Couldn't find user of id " + string(in.GetId())}, nil
}

func (s *Server) ReadUser(ctx context.Context, in *users.Id) (*users.UserResponse, error) {
	return &users.UserResponse{Message: s.Database.ReadUser(in)}, nil
}

func (s *Server) ReadUsers(ctx context.Context, in *users.Empty) (*users.UserResponse, error) {
	return &users.UserResponse{Message: s.Database.ReadUsers()}, nil
}
