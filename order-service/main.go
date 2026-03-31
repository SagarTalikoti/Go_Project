package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "order-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedOrderServiceServer
	orders []*pb.Order
}

func (s *server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	id := fmt.Sprintf("order-%d", len(s.orders)+1)
	order := &pb.Order{
		Id:         id,
		UserId:     req.UserId,
		ItemName:   req.ItemName,
		Quantity:   req.Quantity,
		TotalPrice: req.TotalPrice,
		Status:     "CREATED",
	}
	s.orders = append(s.orders, order)
	return &pb.OrderResponse{Order: order}, nil
}

func (s *server) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest) (*pb.OrderResponse, error) {
	for _, order := range s.orders {
		if order.Id == req.Id {
			order.Status = req.Status
			return &pb.OrderResponse{Order: order}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Order not found")
}

func (s *server) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	for _, order := range s.orders {
		if order.Id == req.Id {
			return &pb.OrderResponse{Order: order}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Order not found")
}

func (s *server) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	return &pb.ListOrdersResponse{Orders: s.orders}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, &server{})

	log.Printf("Order service listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
