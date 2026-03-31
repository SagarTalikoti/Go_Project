package main

import (
	"log"
	"net/http"
	"os"

	"api-gateway/graph"
	orderpb "order-service/proto"
	trackorderpb "track-order-service/proto"
	userpb "user-service/proto"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	orderConn, err := grpc.NewClient("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to order service: %v", err)
	}
	defer orderConn.Close()
	orderClient := orderpb.NewOrderServiceClient(orderConn)

	trackOrderConn, err := grpc.NewClient("localhost:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to track order service: %v", err)
	}
	defer trackOrderConn.Close()
	trackOrderClient := trackorderpb.NewTrackOrderServiceClient(trackOrderConn)

	userConn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to user service: %v", err)
	}
	defer userConn.Close()
	userClient := userpb.NewUserServiceClient(userConn)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		OrderClient:      orderClient,
		TrackOrderClient: trackOrderClient,
		UserClient:       userClient,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
