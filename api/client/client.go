package main

import (
	"checksum/api"
	"flag"
	"fmt"
	"log"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	addr := flag.String("addr", "localhost:4040", "address the server is listening on")
	check := flag.String("check", "", "check a given numeric string for its checksum")
	compute := flag.String("compute", "123456789", "compute the checksum for a given numeric string")
	damm := flag.Bool("damm", true, "use the Damm algorithm")
	verhoeff := flag.Bool("verhoeff", false, "use the Verhoeff algorithm")
	luhn := flag.Bool("luhn", false, "use the Luhn algorithm")
	flag.Parse()

	var conn *grpc.ClientConn
	var response *api.Response
	var err error

	conn, err = grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	client := api.NewChecksumClient(conn)
	switch {

	case *check != "":
		if *verhoeff {
			stream, err := client.VerhoeffCheck(context.Background())
			defer stream.CloseSend()

			if err != nil {
				log.Fatalf("Failed to create stream: %v", err)
			}
			if err := stream.Send(&api.Request{Uuid: uuid.New().String(), Payload: *check}); err != nil {
				log.Printf("Failed to send to stream: %v", err)
				break
			}
			if response, err = stream.Recv(); err != nil {
				log.Printf("Failed to receive from stream: %v", err)
			}
		} else if *damm {
			stream, err := client.DammCheck(context.Background())
			defer stream.CloseSend()

			if err != nil {
				log.Fatalf("Failed to create stream: %v", err)
			}
			if err := stream.Send(&api.Request{Uuid: uuid.New().String(), Payload: *check}); err != nil {
				log.Printf("Failed to send to stream: %v", err)
				break
			}
			if response, err = stream.Recv(); err != nil {
				log.Printf("Failed to receive from stream: %v", err)
			}
		} else if *luhn {
			stream, err := client.LuhnCheck(context.Background())
			defer stream.CloseSend()

			if err != nil {
				log.Fatalf("Failed to create stream: %v", err)
			}
			if err := stream.Send(&api.Request{Uuid: uuid.New().String(), Payload: *check}); err != nil {
				log.Printf("Failed to send to stream: %v", err)
				break
			}
			if response, err = stream.Recv(); err != nil {
				log.Printf("Failed to receive from stream: %v", err)
			}
		}
		if err != nil {
			log.Fatalf("Error when calling gRPC method: %s", err)
		}

	default:
		if *verhoeff {
			stream, err := client.VerhoeffCompute(context.Background())
			defer stream.CloseSend()

			if err != nil {
				log.Fatalf("Failed to create stream: %v", err)
			}
			if err := stream.Send(&api.Request{Uuid: uuid.New().String(), Payload: *compute}); err != nil {
				log.Printf("Failed to send to stream: %v", err)
				break
			}
			if response, err = stream.Recv(); err != nil {
				log.Printf("Failed to receive from stream: %v", err)
			}
		} else if *damm {
			stream, err := client.DammCompute(context.Background())
			defer stream.CloseSend()

			if err != nil {
				log.Fatalf("Failed to create stream: %v", err)
			}
			if err := stream.Send(&api.Request{Uuid: uuid.New().String(), Payload: *compute}); err != nil {
				log.Printf("Failed to send to stream: %v", err)
				break
			}
			if response, err = stream.Recv(); err != nil {
				log.Printf("Failed to receive from stream: %v", err)
			}
		} else if *luhn {
			stream, err := client.DammCompute(context.Background())
			defer stream.CloseSend()

			if err != nil {
				log.Fatalf("Failed to create stream: %v", err)
			}
			if err := stream.Send(&api.Request{Uuid: uuid.New().String(), Payload: *compute}); err != nil {
				log.Printf("Failed to send to stream: %v", err)
				break
			}
			if response, err = stream.Recv(); err != nil {
				log.Printf("Failed to receive from stream: %v", err)
			}
		}
		if err != nil {
			log.Fatalf("Error when calling gRPC method: %s", err)
		}
	}

	fmt.Printf("Response from server: \n%v\n", response)

}
