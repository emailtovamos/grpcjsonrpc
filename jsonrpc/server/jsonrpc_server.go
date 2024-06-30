package main

import (
	"net/http"
	"github.com/gorilla/rpc" // Import the gorilla/rpc package for creating the RPC server
	"github.com/gorilla/rpc/json" // Import JSON codec for RPC
)

// Define a structure for the arguments expected by the function
type Args struct {
	X, Y int
}

// Define a service that will handle incoming RPC requests
type JSONRPCDemoService struct{}

// Add is a method of JSONRPCDemoService that handles addition requests
func (t *JSONRPCDemoService) Add(r *http.Request, args *Args, reply *int) error {
	*reply = args.X + args.Y // Perform addition and store the result in reply
	return nil // Return no error
}

func main() {
	s := rpc.NewServer() // Create a new RPC server
	s.RegisterCodec(json.NewCodec(), "application/json") // Register the JSON codec for the server
	s.RegisterService(new(JSONRPCDemoService), "") // Register the JSONRPCDemoService with the server
	http.Handle("/rpc", s) // Set up the HTTP route
	http.ListenAndServe(":1234", nil) // Start listening on port 1234
}
