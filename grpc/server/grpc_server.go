package main

import (
	"context"
	"net"
	"google.golang.org/grpc"
	calculator "github.com/youtube/grpcjsonrpc/grpc/proto/proto" // Replace with your actual package path
)

// server is a struct that implements the calculator service
type server struct {
	calculator.UnimplementedCalculatorServer
}

// Add is a function that handles the addition operation
func (s *server) Add(ctx context.Context, in *calculator.AddRequest) (*calculator.AddResponse, error) {
	return &calculator.AddResponse{Result: in.X + in.Y}, nil // Return the sum of X and Y
}

func main() {
	lis, _ := net.Listen("tcp", ":12345") // Listen on TCP port 12345
	s := grpc.NewServer() // Create a new gRPC server
	calculator.RegisterCalculatorServer(s, &server{}) // Register the server with gRPC
	s.Serve(lis) // Start serving
}
