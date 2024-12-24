package ports

import (
	"context"

	"github.com/chyiyaqing/gmicro-order/internal/application/core/domain"
)

type UserPort interface {
	Get(context.Context, int64) (domain.User, error)
}
