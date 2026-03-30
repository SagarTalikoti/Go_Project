package main

import (
	"context"
	"errors"
	"log"
	"net"
	"sync"
	"crypto/rand"
	"encoding/hex"

	pb "task-service/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTaskServiceServer
	mu    sync.RWMutex
	tasks map[string]*pb.Task
}

func generateID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (s *server) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := &pb.Task{
		Id:          generateID(),
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
	}
	s.tasks[task.Id] = task

	return &pb.TaskResponse{Task: task}, nil
}

func (s *server) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.TaskResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[req.Id]
	if !exists {
		return nil, errors.New("task not found")
	}

	return &pb.TaskResponse{Task: task}, nil
}

func (s *server) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var tasks []*pb.Task
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}

	return &pb.ListTasksResponse{Tasks: tasks}, nil
}

func (s *server) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.TaskResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[req.Id]
	if !exists {
		return nil, errors.New("task not found")
	}

	task.Title = req.Title
	task.Description = req.Description
	task.Completed = req.Completed
	s.tasks[req.Id] = task

	return &pb.TaskResponse{Task: task}, nil
}

func (s *server) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[req.Id]; !exists {
		return &pb.DeleteTaskResponse{Success: false}, errors.New("task not found")
	}

	delete(s.tasks, req.Id)
	return &pb.DeleteTaskResponse{Success: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTaskServiceServer(s, &server{
		tasks: make(map[string]*pb.Task),
	})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
