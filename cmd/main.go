package main

import (
	"os"

	"github.com/chyiyaqing/gmicro-order/config"
	"github.com/chyiyaqing/gmicro-order/internal/adapters/db"
	"github.com/chyiyaqing/gmicro-order/internal/adapters/grpc"
	"github.com/chyiyaqing/gmicro-order/internal/adapters/payment"
	"github.com/chyiyaqing/gmicro-order/internal/adapters/shipping"
	"github.com/chyiyaqing/gmicro-order/internal/adapters/user"
	"github.com/chyiyaqing/gmicro-order/internal/application/core/api"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	service     = "order"
	environment = "dev"
	id          = 1
)

type customLogger struct {
	formatter log.JSONFormatter
}

// Format(*Entry) ([]byte, error)
func (l *customLogger) Format(entry *log.Entry) ([]byte, error) {
	span := trace.SpanFromContext(entry.Context)
	entry.Data["trace_id"] = span.SpanContext().TraceID().String()
	entry.Data["span_id"] = span.SpanContext().SpanID().String()
	// Below injection is Just to understand what Context has
	entry.Data["Context"] = span.SpanContext()
	return l.formatter.Format(entry)
}

func init() {
	log.SetFormatter(&customLogger{
		formatter: log.JSONFormatter{
			FieldMap: log.FieldMap{
				"msg": "message",
			},
		},
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			attribute.String("environment", environment),
			attribute.Int64("ID", id),
		)),
	)
	return tp, nil
}

func main() {
	// tp, err := tracerProvider("http://jaeger-otel.jaeger.svc.cluster.local:14278/api/traces")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// otel.SetTracerProvider(tp)
	// otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}))

	dbAdapter, err := db.NewAdapter(config.GetSqliteDB())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err.Error())
	}

	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub, Error: %v", err)
	}

	userAdapter, err := user.NewAdapter(config.GetUserServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize user stub, Error: %v", err)
	}

	shippingAdapter, err := shipping.NewAdapter(config.GetShippingServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize shipping stub, Error: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter, userAdapter, shippingAdapter)
	grpcAdapter := grpc.NewAdaptor(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
