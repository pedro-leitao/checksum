package api

import (
	"checksum"
	"log"

	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// DammCompute serves a request for computing a checksum using the Damm algorithm
func (s *Server) DammCompute(ctx context.Context, in *Request) (*Response, error) {
	var dm checksum.Damm

	log.Printf("Received request %v", in)

	_, ns, err := dm.Compute(in.Payload)
	if err != nil {
		log.Printf("Failed to process: %v", err)
		return &Response{Payload: "", Valid: false, Error: err.Error()}, nil
	}
	return &Response{Payload: ns, Valid: true, Error: ""}, nil
}

// VerhoeffCompute serves a request for computing a checksum using the Verhoeff algorithm
func (s *Server) VerhoeffCompute(ctx context.Context, in *Request) (*Response, error) {
	var vh checksum.Verhoeff

	log.Printf("Received request %v", in)

	_, ns, err := vh.Compute(in.Payload)
	if err != nil {
		log.Printf("Failed to process: %v", err)
		return &Response{Payload: "", Valid: false, Error: err.Error()}, nil
	}
	return &Response{Payload: ns, Valid: true, Error: ""}, nil
}

// DammCheck serves a request for validating a checksum using the Damm algorithm
func (s *Server) DammCheck(ctx context.Context, in *Request) (*Response, error) {
	var dm checksum.Damm

	log.Printf("Received request %v", in)

	valid, err := dm.Check(in.Payload)
	if err != nil {
		log.Printf("Failed to process: %v", err)
		return &Response{Payload: in.Payload, Valid: false, Error: err.Error()}, nil
	}
	return &Response{Payload: in.Payload, Valid: valid, Error: ""}, nil
}

// VerhoeffCheck serves a request for validating a checksum using the Verhoeff algorithm
func (s *Server) VerhoeffCheck(ctx context.Context, in *Request) (*Response, error) {
	var vh checksum.Verhoeff

	log.Printf("Received request %v", in)

	valid, err := vh.Check(in.Payload)
	if err != nil {
		log.Printf("Failed to process: %v", err)
		return &Response{Payload: in.Payload, Valid: false, Error: err.Error()}, nil
	}
	return &Response{Payload: in.Payload, Valid: valid, Error: ""}, nil
}
