package graph

import (
	orderpb "order-service/proto"
	trackorderpb "track-order-service/proto"
	userpb "user-service/proto"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	OrderClient      orderpb.OrderServiceClient
	TrackOrderClient trackorderpb.TrackOrderServiceClient
	UserClient       userpb.UserServiceClient
}
