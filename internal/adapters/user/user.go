package user

import (
	"context"
	"time"

	"github.com/chyiyaqing/gmicro-order/internal/application/core/domain"
	"github.com/chyiyaqing/gmicro-proto/golang/user"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	user user.UserClient
}

func NewAdapter(userServiceUrl string) (*Adapter, error) {
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
	conn, err := grpc.NewClient(userServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	client := user.NewUserClient(conn)
	return &Adapter{user: client}, nil
}

func (a *Adapter) Get(ctx context.Context, user_id int64) (domain.User, error) {
	getUserResponse, err := a.user.Get(ctx, &user.GetUserRequest{
		UserId: user_id,
	})
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		ID:      getUserResponse.UserId,
		Name:    getUserResponse.Name,
		Email:   getUserResponse.Email,
		Phone:   getUserResponse.Phone,
		Address: getUserResponse.Address,
	}, nil
}
