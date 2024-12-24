package shipping

import (
	"context"
	"time"

	"github.com/chyiyaqing/gmicro-order/internal/application/core/domain"
	"github.com/chyiyaqing/gmicro-proto/golang/shipping"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	shipping shipping.ShippingClient
}

func NewAdapter(shippingServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
				grpc_retry.WithMax(5),
				grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)),
			),
		),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	conn, err := grpc.NewClient(shippingServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	client := shipping.NewShippingClient(conn)
	return &Adapter{shipping: client}, nil
}

func (a *Adapter) Create(ctx context.Context, order *domain.Order, address string) error {
	_, err := a.shipping.Create(ctx, &shipping.CreateShippingRequest{
		UserId:  order.CustomerID,
		OrderId: order.ID,
		Address: address,
	})
	return err
}
