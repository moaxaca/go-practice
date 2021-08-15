package internal

import (
	"context"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/trace"
	"io.parcely.address_validation/pkg/address_validator"
	"log"
	"os"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type IocContainer struct {
	AddressValidator *address_validator.AddressValidator
}

func (f *IocContainer) GetTracer(ctx context.Context) trace.Tracer {
	exporterStrategy := os.Getenv("OTEL_EXPORTER")
	if exporterStrategy == "gcp" {
		projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
		exporter, err := texporter.NewExporter(texporter.WithProjectID(projectID))
		if err != nil {
			log.Fatalf("texporter.NewExporter: %v", err)
		}
		tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter))
		defer func(tp *sdktrace.TracerProvider, ctx context.Context) {
			err := tp.ForceFlush(ctx)
			if err != nil {
				log.Fatalf("tp.ForceFlush: %v", err)
			}
		}(tp, ctx) // flushes any pending spans
		otel.SetTracerProvider(tp)
	} else {
		traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
		if err != nil {
			log.Fatalf("texporter.NewExporter: %v", err)
		}
		bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
		tp := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(bsp))
		otel.SetTracerProvider(tp)
	}
	return otel.GetTracerProvider().Tracer("av")
}

func CreateIoc() IocContainer {
	ioc := IocContainer{}

	smartyCredentials := address_validator.SmartyStreetsCredentials{}
	smartyCredentials.AuthId = os.Getenv("SMART_AUTH_ID")
	smartyCredentials.AuthToken = os.Getenv("SMART_AUTH_TOKEN")
	addressValidatorBuilder := address_validator.CreateBuilder()
	addressValidatorBuilder.WithInMemoryCache()
	addressValidatorBuilder.WithSmartyValidator(smartyCredentials)
	addressValidator, _ := addressValidatorBuilder.Build()
	// TODO Error Handle
	ioc.AddressValidator = &addressValidator

	return ioc
}
