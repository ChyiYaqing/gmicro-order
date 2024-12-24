package ports

import (
	"context"

	"github.com/chyiyaqing/gmicro-order/internal/application/core/domain"
)

type ShippingPort interface {
	Create(ctx context.Context, order *domain.Order, address string) error
}
