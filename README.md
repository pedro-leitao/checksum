# A set of standard checksum algorithms implemented in Go

This is an implementation of various checksum standard algorithms (and associated tests), mostly in use, or described in academic papers. It acts as more of a learning exercise around gRPC than anything else, but the algorithms work as intended.

Implemented as of now:

- Damm algorithm (https://en.wikipedia.org/wiki/Damm_algorithm)
- Verhoeff scheme (https://en.wikipedia.org/wiki/Verhoeff_algorithm)
- Luhn algorithm (https://en.wikipedia.org/wiki/Luhn_algorithm)

It comes with a gRPC endpoint implementation, and a streaming client/server which can be used to invoke/serve calls.

### Compiling the protobuf definition

To compile the protobuf definition of the gRPC API, run:

    $ protoc -I api/ -I${GOPATH}/src --go_out=plugins=grpc:api api/checksum.proto

You can then build the client/server:

    $ cd api/client; go build
    $ cd api/server; go build
    
And finally run a few examples:

    $ ./server -help
    Usage of ./server:
    -port int
    	port the server should listen on (default 4040)
    $ ./server
    2020/05/04 00:44:40 Listening on port 4040
    $ ./client -help
      -addr string
    	address the server is listening on (default "localhost:4040")
      -algo string
    	the checksum algorithm to use ('luhn', 'damm', 'verhoeff') (default "luhn")
      -check
    	check a given numeric string for its checksum
      -compute
    	compute the checksum for a given numeric string
      -payload string
    	the payload to compute the checksum (default "123456789")
      -stream
    	stream content from stdin, one line at a time
     $ ./client -addr :4040 -compute -payload "1234567897" -algo=verhoeff
	Response from server:
	uuid:"ed13dc88-19d1-4aa3-9b4c-478482e3f59e"  payload:"12345678973"  valid:true
     $ ./client -addr :4040 -check -payload "12345678973" -algo=verhoeff
	Response from server:
	uuid:"479e82d7-1780-446d-b9f3-941cdc8d47e6"  payload:"12345678973"  valid:true
     $ ./client -addr :4040 -compute -stream -algo=verhoeff < random.txt
	Response from server:
	uuid:"917de14c-5f2e-4d91-8ed8-3fc1fb962f49"  payload:"8740320014"  valid:true
	Response from server:
	uuid:"b0107364-11ed-4eb8-8e50-57d603519916"  payload:"1539518308"  valid:true
	Response from server:
	uuid:"44fcb134-267e-40e0-9d7b-ffba8d67b0df"  payload:"8107150221"  valid:true
	Response from server:
	uuid:"88b7ed89-e2b4-4813-badb-d20c982924a3"  payload:"2942262509"  valid:true
	...
    
### Benchmarking and testing

Tests for the various algorithms are included, as are benchmark tests:

    go test -bench=.
    goos: linux
    goarch: arm
    pkg: checksum
    BenchmarkDamm-4       	  351817	      3467 ns/op
    BenchmarkLuhn-4       	 1257541	       920 ns/op
    BenchmarkVerhoeff-4   	  195346	      6122 ns/op
    PASS
    ok  	checksum	4.654s
    
Even on my incredibly tiny Raspberry Pi the algorithms are blazing fast, with the Damm algorithm running 10 digit compute/check cycles at about 350,000 cycles per second, and the Luhn algorithm running over 1.2 million cycles per second (one has to love Go for performance).
