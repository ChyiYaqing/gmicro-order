package api

import (
	"context"

	"github.com/chyiyaqing/gmicro-order/internal/application/core/domain"
	"github.com/chyiyaqing/gmicro-order/internal/ports"
)

type Application struct {
	db ports.DBPort
}

// 依赖注入机制：将数据库适配器(特定的数据库技术的具体实现)注入到应用程序中，以便API可以将特定订单的状态存储在数据库中
func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a *Application) SaveOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	err := a.db.Save(ctx, &order)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

func (a *Application) GetOrder(ctx context.Context, order_id int64) (domain.Order, error) {
	return a.db.Get(ctx, order_id)
}
