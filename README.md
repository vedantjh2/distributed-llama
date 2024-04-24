This repo helps you run LlaMa3-8B-instruct locally and lets the server handle multiple requests concurrently using go routines


To run server, open a new terminal window and type:

cd path/to/distributed-llama/server
go build # If you built it, or use `go run server.go` directly
./server

To run client, open a new terminal window and type:

cd path/to/distributed-llama/client
go build # If you built it, or use `go run client.go` directly
./client
