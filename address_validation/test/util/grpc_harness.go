package util

import (
	"google.golang.org/grpc"
	"log"
)

const (
	address     = "localhost:3001"
)

type TestGrpcHarness struct {
	Connection *grpc.ClientConn
}

func (th *TestGrpcHarness) Cleanup() {
	defer func(Connection *grpc.ClientConn) {
		err := Connection.Close()
		if err != nil {
			log.Fatalf("did close connection: %v", err)
		}
	}(th.Connection)
}

func CreateTestGrpcHarness() TestGrpcHarness {
	th := TestGrpcHarness{}
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	th.Connection = conn
	return th
}
