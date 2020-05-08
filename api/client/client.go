package main

import (
	"bufio"
	"checksum/api"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	var response *api.Response
	var err error

	addr := flag.String("addr", "localhost:4040", "address the server is listening on")
	check := flag.Bool("check", false, "check a given numeric string for its checksum")
	compute := flag.Bool("compute", false, "compute the checksum for a given numeric string")
	payload := flag.String("payload", "123456789", "the payload to compute the checksum")
	algo := flag.String("algo", "luhn", "the checksum algorithm to use ('luhn', 'damm', 'verhoeff')")
	stream := flag.Bool("stream", false, "stream content from stdin, one line at a time")
	flag.Parse()

	conn, err = grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	client := api.NewChecksumClient(conn)

	switch {

	case *check:

		strm, err := client.Check(context.Background())
		defer strm.Context().Done()
		defer strm.CloseSend()

		if err != nil {
			log.Fatalf("Failed to create stream: %v", err)
		}
		if *stream {
			scanner := bufio.NewScanner(os.Stdin)
			log.Printf("Reading payload from <stdin> one line at a time...")
			for scanner.Scan() {
				if err := strm.Send(&api.Request{Uuid: uuid.New().String(), Algo: *algo, Payload: scanner.Text()}); err != nil {
					log.Printf("Failed to send to stream: %v", err)
					break
				}
				if response, err = strm.Recv(); err != nil {
					log.Printf("Failed to receive from stream: %v", err)
					break
				}
				fmt.Printf("Response from server: \n%v\n", response)
			}
		} else {
			if err := strm.Send(&api.Request{Uuid: uuid.New().String(), Algo: *algo, Payload: *payload}); err != nil {
				log.Printf("Failed to send to stream: %v", err)
				break
			}
			if response, err = strm.Recv(); err != nil {
				log.Printf("Failed to receive from stream: %v", err)
			}
			fmt.Printf("Response from server: \n%v\n", response)
		}

	case *compute:
		strm, err := client.Compute(context.Background())
		defer strm.Context().Done()
		defer strm.CloseSend()

		if err != nil {
			log.Fatalf("Failed to create stream: %v", err)
		}
		if *stream {
			scanner := bufio.NewScanner(os.Stdin)
			log.Printf("Reading payload from <stdin> one line at a time...")
			for scanner.Scan() {
				if err := strm.Send(&api.Request{Uuid: uuid.New().String(), Algo: *algo, Payload: scanner.Text()}); err != nil {
					log.Printf("Failed to send to stream: %v", err)
					break
				}
				if response, err = strm.Recv(); err != nil {
					log.Printf("Failed to receive from stream: %v", err)
					break
				}
				fmt.Printf("Response from server: \n%v\n", response)
			}
		} else {
			if err := strm.Send(&api.Request{Uuid: uuid.New().String(), Algo: *algo, Payload: *payload}); err != nil {
				log.Printf("Failed to send to stream: %v", err)
				break
			}
			if response, err = strm.Recv(); err != nil {
				log.Printf("Failed to receive from stream: %v", err)
			}
			fmt.Printf("Response from server: \n%v\n", response)
		}
	}

}
