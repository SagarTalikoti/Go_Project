package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"net"
	"sync"

	pb "user-service/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
	mu    sync.RWMutex
	users map[string]*pb.User
}

func generateID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := &pb.User{
		Id:       generateID(),
		Username: req.Username,
		Email:    req.Email,
	}
	s.users[user.Id] = user

	return &pb.UserResponse{User: user}, nil
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[req.Id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return &pb.UserResponse{User: user}, nil
}

func (s *server) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var users []*pb.User
	for _, user := range s.users {
		users = append(users, user)
	}

	return &pb.ListUsersResponse{Users: users}, nil
}

func (s *server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[req.Id]
	if !exists {
		return nil, errors.New("user not found")
	}

	user.Username = req.Username
	user.Email = req.Email
	s.users[req.Id] = user

	return &pb.UserResponse{User: user}, nil
}

func (s *server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[req.Id]; !exists {
		return &pb.DeleteUserResponse{Success: false}, errors.New("user not found")
	}

	delete(s.users, req.Id)
	return &pb.DeleteUserResponse{Success: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{
		users: make(map[string]*pb.User),
	})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
