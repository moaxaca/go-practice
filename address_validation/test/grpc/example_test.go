package grpc_test

import (
	"context"
	pb "io.parcely.address_validation/api/grpc/proto"
	"io.parcely.address_validation/test/util"
	"log"
	"testing"
	"time"
)

func TestExample(t *testing.T) {
	th := util.CreateTestGrpcHarness()
	client := pb.NewExampleClient(th.Connection)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.Ping(ctx, &pb.ExampleRequest{Message: "test"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetRelay())
	th.Cleanup()
}
