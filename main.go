// main.go
package main

import (
	"context"
	"github.com/go-logr/stdr"
	"github.com/kizmey/order_management_system/config"
	"github.com/kizmey/order_management_system/database"
	logger "github.com/kizmey/order_management_system/logs"
	"github.com/kizmey/order_management_system/pkg"
	serverPkg "github.com/kizmey/order_management_system/server/httpEchoServer"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"log"
	"os"
)

func main() {
	// Initialize Logger
	logger.InitLogger()
	defer func(LogFile *os.File) {
		err := LogFile.Close()
		if err != nil {
		}
	}(logger.LogFile)

	// Initialize Loki client
	//logger.InitLokiClient()

	// Log to file and Loki
	fields := logrus.Fields{"module": "main", "function": "main"}
	logger.LogInfo("Service started", fields)
	//logger.LogToLoki(logrus.InfoLevel, "Service started", fields)

	conf := config.GettingConfig()
	db := database.NewPostgresDatabase(conf.Database)

	usecases := pkg.InitUsecase(db)
	err := InitOpenTelemetry()
	if err != nil {
		logger.LogError("Failed to initialize OpenTelemetry"+err.Error(), fields)
	}

	server := serverPkg.NewEchoServer(conf, usecases)
	server.Start()
}

func InitOpenTelemetry() error {
	// Set up loggerx
	loggerx := stdr.New(log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile))
	otel.SetLogger(loggerx)

	// Set up Propagators
	propagator := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
	otel.SetTextMapPropagator(propagator)

	// Set up trace provider
	ctx := context.Background()
	traceExporter, err := otlptrace.New(ctx, otlptracehttp.NewClient(
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithEndpoint("localhost:4318"),
	))
	if err != nil {
		return err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExporter),
	)
	otel.SetTracerProvider(tp)
	return nil
}
