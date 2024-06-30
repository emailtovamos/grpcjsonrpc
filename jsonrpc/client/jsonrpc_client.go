package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Args struct holds the arguments to be sent to the server
type Args struct {
	X, Y int
}

func main() {
	args := &Args{X: 1, Y: 2} // Set the arguments
	data, _ := json.Marshal(args) // Marshal the arguments into JSON format
	// Send an HTTP POST request to the server
	response, _ := http.Post("http://localhost:1234/rpc", "application/json", bytes.NewBufferString(`{"method":"JSONRPCDemoService.Add","params":[`+string(data)+`],"id":1}`))
	var result map[string]interface{} // Map to store the JSON response
	json.NewDecoder(response.Body).Decode(&result) // Decode the response into the map
	fmt.Println(result["result"]) // Print the result from the map
}
