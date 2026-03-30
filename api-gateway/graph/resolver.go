package graph

import (
	pb "task-service/proto"
	userpb "user-service/proto"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	TaskClient pb.TaskServiceClient
	UserClient userpb.UserServiceClient
}
