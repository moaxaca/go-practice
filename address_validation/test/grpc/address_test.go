package grpc_test

import (
	"context"
	pb "io.parcely.address_validation/api/grpc/proto"
	"io.parcely.address_validation/test/util"
	"log"
	"testing"
	"time"
)

func TestAddressValidationGrpc(t *testing.T) {
	th := util.CreateTestGrpcHarness()
	tests := []struct {
		status  int
		request pb.AddressValidationRequest
	}{
		{status: 200, request: pb.AddressValidationRequest{AddressLines: []string {"6272 Pacific Coast Hwy"}, Locality: "Long Beach", PostalCode: "90803", Region: "CA", CountryCode: "USA"}},
	}
	for _, tc := range tests {
		client := pb.NewAddressClient(th.Connection)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second * 30)
		defer cancel()
		r, err := client.Validate(ctx, &tc.request)
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetSuccessResponse().Uuid)
	}
	th.Cleanup()
}
