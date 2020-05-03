package main

import (
	"checksum/api"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

// main start a gRPC server and waits for connection
func main() {

	port := flag.Int("port", 4040, "port the server should listen on")

	flag.Parse()

	// Listen on the chosen port

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to start listener: %v", err)
	}
	log.Printf("Listening on port %v\n", *port)

	// Create a server instance from the gRPC API
	s := api.Server{}
	grpcServer := grpc.NewServer()
	api.RegisterChecksumServer(grpcServer, &s) // Start the gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
