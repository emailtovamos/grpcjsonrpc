package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	calculator "github.com/youtube/grpcjsonrpc/grpc/proto/proto" // Replace with your actual package path
)

func main() {
	conn, _ := grpc.Dial("localhost:12345", grpc.WithInsecure()) // Connect to the gRPC server
	defer conn.Close() // Ensure the connection is closed when done
	c := calculator.NewCalculatorClient(conn) // Create a new Calculator client

	r, _ := c.Add(context.Background(), &calculator.AddRequest{X: 1, Y: 2}) // Send an Add request
	fmt.Printf("gRPC client received: %d\n", r.Result) // Print the result received from the server
}
