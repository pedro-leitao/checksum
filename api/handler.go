package api

import (
	"checksum"
	"io"
	"log"
)

// Server represents the gRPC server
type Server struct {
}

// DammCompute serves a request for computing a checksum using the Damm algorithm
func (s *Server) DammCompute(srv Checksum_DammComputeServer) error {
	var algo checksum.Damm

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

// VerhoeffCompute serves a request for computing a checksum using the Verhoeff algorithm
func (s *Server) VerhoeffCompute(srv Checksum_VerhoeffComputeServer) error {
	var algo checksum.Verhoeff

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

// DammCheck serves a request for validating a checksum using the Damm algorithm
func (s *Server) DammCheck(srv Checksum_DammCheckServer) error {
	var algo checksum.Damm

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

// VerhoeffCheck serves a request for validating a checksum using the Verhoeff algorithm
func (s *Server) VerhoeffCheck(srv Checksum_VerhoeffCheckServer) error {
	var algo checksum.Verhoeff

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

// LuhnCompute serves a request for computing a checksum using the Luhn algorithm
func (s *Server) LuhnCompute(srv Checksum_LuhnComputeServer) error {
	var algo checksum.Luhn

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

// LuhnCheck serves a request for validating a checksum using the Luhn algorithm
func (s *Server) LuhnCheck(srv Checksum_LuhnCheckServer) error {
	var algo checksum.Luhn

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
