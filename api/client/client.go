package main

import (
	"checksum/api"
	"flag"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
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

	c := api.NewChecksumClient(conn)
	switch {

	case *check != "":
		var err error
		if *verhoeff {
			response, err = c.VerhoeffCheck(context.Background(), &api.Request{Payload: *check})
		} else if *damm {
			response, err = c.DammCheck(context.Background(), &api.Request{Payload: *check})
		} else if *luhn {
			response, err = c.LuhnCheck(context.Background(), &api.Request{Payload: *check})
		}
		if err != nil {
			log.Fatalf("Error when calling gRPC method: %s", err)
		}

	default:
		var err error
		if *verhoeff {
			response, err = c.VerhoeffCompute(context.Background(), &api.Request{Payload: *compute})
		} else if *damm {
			response, err = c.DammCompute(context.Background(), &api.Request{Payload: *compute})
		} else if *luhn {
			response, err = c.LuhnCompute(context.Background(), &api.Request{Payload: *compute})
		}
		if err != nil {
			log.Fatalf("Error when calling gRPC method: %s", err)
		}
	}

	fmt.Printf("Response from server: \n%v\n", spew.Sprint(response))

}