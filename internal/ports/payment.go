package ports

import (
	"context"

	"github.com/chyiyaqing/gmicro-order/internal/application/core/domain"
)

type PaymentPort interface {
	Charge(context.Context, *domain.Order) error
}
