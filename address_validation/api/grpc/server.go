package grpc

import (
	"google.golang.org/grpc"
	"io.parcely.address_validation/api/grpc/handlers"
	example "io.parcely.address_validation/api/grpc/proto"
	"io.parcely.address_validation/internal"
)

func CreateServer(ioc internal.IocContainer) *grpc.Server {
	server := grpc.NewServer()
	addressHandler := handlers.CreateAddressHandler(*ioc.AddressValidator)
	example.RegisterAddressServer(server, addressHandler)
	exampleHandler := handlers.CreateExampleHandler()
	example.RegisterExampleServer(server, exampleHandler)
	return server
}
