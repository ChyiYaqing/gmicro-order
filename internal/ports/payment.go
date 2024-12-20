package ports

import (
	"context"

	"github.com/chyiyaqing/gmicro-order/internal/application/core/domain"
)

//go:generate mockgen -destination payment_mock.go -package ports -source payment.go
type PaymentPort interface {
	Charge(context.Context, *domain.Order) error
}
