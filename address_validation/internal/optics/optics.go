package optics

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	_ "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	_ "go.opentelemetry.io/otel/propagation"
	_ "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func recordContext(ctx context.Context) {
	tracer := otel.Tracer("ex.com/basic")

	// we're ignoring errors here since we know these values are valid,
	// but do handle them appropriately if dealing with user-input
	foo, _ := baggage.NewMember("ex.com.foo", "foo1")
	bar, _ := baggage.NewMember("ex.com.bar", "bar1")
	bag, _ := baggage.New(foo, bar)
	ctx = baggage.ContextWithBaggage(ctx, bag)

	func(ctx context.Context) {
		var span trace.Span
		ctx, span = tracer.Start(ctx, "operation")
		defer span.End()
		span.AddEvent("Nice operation!", trace.WithAttributes(attribute.Int("bogons", 100)))
		func(ctx context.Context) {
			var span trace.Span
			ctx, span = tracer.Start(ctx, "Sub operation...")
			defer span.End()
			span.AddEvent("Sub span event")
		}(ctx)
	}(ctx)
}
