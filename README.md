# A set of standard checksum algorithms implemented in Go

This is an implementation of various checksum standard algorithms (and associated tests), mostly in use, or described in academic papers. It acts as more of a learning exercise around gRPC than anything else, but the algorithms work as intended.

Implemented as of now:

- Damm algorithm (https://en.wikipedia.org/wiki/Damm_algorithm)
- Verhoeff scheme (https://en.wikipedia.org/wiki/Verhoeff_algorithm)

It comes with a gRPC endpoint implementation, and a client/server which can be used to invoke/serve calls.

### Compiling the protobuf definition

To compile the protobuf definition of the gRPC API, run:

    $ protoc -I api/ -I${GOPATH}/src --go_out=plugins=grpc:api api/checksum.proto

You can then build the client/server:

    $ cd api/client; go build
    $ cd api/server; go build
    
And finally run a few examples:

    $ ./server
    2020/05/04 00:44:40 Listening on port 4040
    $ ./client -addr :4040 -compute "123456789" -damm
    Response from server:
    <*>payload:"1234567894"  valid:true
    $ ./client -addr :4040 -check "1234567894" -damm
    Response from server:
    <*>payload:"1234567894"  valid:true
    
### Benchmarking and testing

Tests for the various algorithms are included, as are benchmark tests:

    $ go test
    PASS
    ok  	checksum	0.023s
    $ go test -bench=.
    goos: linux
    goarch: arm
    pkg: checksum
    BenchmarkDamm-4       	  349975	      3298 ns/op
    BenchmarkVerhoeff-4   	  195798	      6260 ns/op
    PASS
    ok  	checksum	2.493s
    
Even on my incredibly tiny Raspberry Pi the algorithms are blazing fast, with the Damm algorithm running 10 digit compute/check cycles at about 350,000 cycles per second (one has to love Go for performance).
