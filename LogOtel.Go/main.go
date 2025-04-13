package main

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	otlogsdk "go.opentelemetry.io/otel/sdk/log"
)

func main() {
	//Set up OTLP exporter
	ctx:=context.Background()
	exporter, err := otlploggrpc.New(ctx,
		otlploggrpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}

	processor := otlogsdk.NewBatchProcessor(exporter)
	provider := otlogsdk.NewLoggerProvider(otlogsdk.WithProcessor(processor))
	defer func() {
		if err := provider.Shutdown(ctx); err != nil {
			panic(err)
		}
	}()
	
	// 3. Bridge OpenTelemetry with slog
	otelLogger := otelslog.NewLogger("otel-logger",  otelslog.WithLoggerProvider(provider))
	slog.SetDefault(otelLogger)
	
	//Log messages (will be sent to OTEL Collector)
	slog.Info("hello from slog", 
		slog.String("app", "my-app"),
	)
}