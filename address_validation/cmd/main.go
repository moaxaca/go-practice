package main

import (
	"context"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"io.parcely.address_validation/api/grpc"
	"io.parcely.address_validation/api/rest"
	"io.parcely.address_validation/internal"
	"log"
	"net"
	"os"
	"sync"
)

func main() {
	ctx := context.Background()
	ioc := internal.CreateIoc()

	tracer := ioc.GetTracer(ctx)
	err := func(ctx context.Context) error {
		_, span := tracer.Start(ctx, "foo")
		defer span.End()
		return nil
	}(ctx)
	if err != nil {
		log.Fatalf("span test: %v", err)
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		// GRPC Construction
		lis, err := net.Listen("tcp", ":"+os.Getenv("APP_GRPC_PORT"))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.CreateServer(ioc)
		log.Printf("server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve grpc: %v", err)
		}
		wg.Done()
	}()

	go func() {
		// Rest Server
		restServerConfig := rest.RestServerConfiguration{}
		restServerConfig.Name = "Address Validation Service"
		restServerConfig.Address = ":"+os.Getenv("APP_REST_PORT")
		restServer := rest.CreateRestServer(restServerConfig, ioc)
		// Service
		service := micro.NewService(
			micro.Server(restServer),
			micro.Registry(registry.NewRegistry()),
		)
		service.Init()
		err := service.Run()
		if err != nil {
			log.Fatalf("failed to serve rest: %v", err)
			return
		}
		wg.Done()
	}()
	wg.Wait()
}
