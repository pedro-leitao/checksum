package api

import (
	"checksum"
	"io"
	"log"
)

// checksummer is a generic interface to checksum algorithms
type checksummer interface {
	Check(s string) (bool, error)
	Compute(s string) (int, string, error)
}

// Server represents the gRPC server
type Server struct {
}

// Compute serves a request for computing a checksum
func (s *Server) Compute(srv Checksum_ComputeServer) error {
	var algo checksummer

	ctx := srv.Context()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		req, err := srv.Recv()
		if err == io.EOF {
			log.Printf("Stream was closed: %v", err)
			return nil
		}
		if err != nil {
			log.Printf("Failed to receive: %v", err)
		}

		switch req.Algo {
		case "damm":
			algo = &checksum.Damm{}
		case "verhoeff":
			algo = &checksum.Verhoeff{}
		default:
			algo = &checksum.Luhn{}
		}
		_, ns, err := algo.Compute(req.Payload)
		if err != nil {
			log.Printf("Failed to process: %v", err)
			if err := srv.Send(&Response{Uuid: req.Uuid, Payload: "", Valid: false, Error: err.Error()}); err != nil {
				log.Printf("Failed to send: %v", err)
			}
		}

		if err := srv.Send(&Response{Uuid: req.Uuid, Payload: ns, Valid: true, Error: ""}); err != nil {
			log.Printf("Failed to send: %v", err)
		}

		log.Printf("Handled request %v", req)
	}
}

// Check serves a request for validating a checksum
func (s *Server) Check(srv Checksum_CheckServer) error {
	var algo checksummer

	ctx := srv.Context()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Failed to receive: %v", err)
		}

		switch req.Algo {
		case "damm":
			algo = &checksum.Damm{}
		case "verhoeff":
			algo = &checksum.Verhoeff{}
		default:
			algo = &checksum.Luhn{}
		}
		valid, err := algo.Check(req.Payload)
		if err != nil {
			log.Printf("Failed to process: %v", err)
			if err := srv.Send(&Response{Uuid: req.Uuid, Payload: req.Payload, Valid: valid, Error: err.Error()}); err != nil {
				log.Printf("Failed to send: %v", err)
			}
		}

		if err := srv.Send(&Response{Uuid: req.Uuid, Payload: req.Payload, Valid: valid, Error: ""}); err != nil {
			log.Printf("Failed to send: %v", err)
		}

		log.Printf("Handled request %v", req)
	}

}
