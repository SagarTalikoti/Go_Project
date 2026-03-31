package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "track-order-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedTrackOrderServiceServer
	trackings []*pb.TrackOrder
}

func (s *server) UpdateStatus(ctx context.Context, req *pb.UpdateStatusRequest) (*pb.TrackOrderResponse, error) {
	for _, t := range s.trackings {
		if t.OrderId == req.OrderId {
			t.ShippingStatus = req.ShippingStatus
			return &pb.TrackOrderResponse{TrackOrder: t}, nil
		}
	}
	id := fmt.Sprintf("track-%d", len(s.trackings)+1)
	newItem := &pb.TrackOrder{
		Id:             id,
		OrderId:        req.OrderId,
		ShippingStatus: req.ShippingStatus,
	}
	s.trackings = append(s.trackings, newItem)
	return &pb.TrackOrderResponse{TrackOrder: newItem}, nil
}

func (s *server) GetTrackingInfo(ctx context.Context, req *pb.GetTrackingInfoRequest) (*pb.TrackOrderResponse, error) {
	for _, t := range s.trackings {
		if t.OrderId == req.OrderId {
			return &pb.TrackOrderResponse{TrackOrder: t}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Tracking info not found")
}

func main() {
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTrackOrderServiceServer(s, &server{})

	log.Printf("Track order service listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
