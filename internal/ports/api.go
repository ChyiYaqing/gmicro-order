package ports

import (
	"context"

	"github.com/chyiyaqing/gmicro-order/internal/application/core/domain"
)

type APIPort interface {
	SaveOrder(ctx context.Context, order domain.Order) (domain.Order, error)
	GetOrder(ctx context.Context, id int64) (domain.Order, error)
}
