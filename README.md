# A set of standard checksum algorithms implemented in Go

This is an implementation of various checksum standard algorithms (and associated tests), mostly in use, or described in academic papers.

Implemented as of now:

- Damm algorithm (https://en.wikipedia.org/wiki/Damm_algorithm)
- Verhoeff scheme (https://en.wikipedia.org/wiki/Verhoeff_algorithm)

## Compiling the protobuf definition

protoc -I api/ -I${GOPATH}/src --go_out=plugins=grpc:api api/checksum.proto
