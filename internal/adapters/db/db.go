package db

import (
	"context"
	"fmt"

	"github.com/chyiyaqing/gmicro-order/internal/application/core/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	ProducCode string
	UnitPrice  float32
	Quantity   int32
	OrderID    int64
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(sqliteDB string) (*Adapter, error) {
	db, openErr := gorm.Open(sqlite.Open(sqliteDB), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db open %v error: %v", sqliteDB, openErr)
	}
	err := db.AutoMigrate(&Order{}, &OrderItem{}) // 确保表创建正确
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}

func (a *Adapter) Get(ctx context.Context, id int64) (domain.Order, error) {
	var orderEntity Order
	res := a.db.First(&orderEntity, id)
	var orderItems []domain.OrderItem
	for _, orderItem := range orderEntity.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProducCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	order := domain.Order{
		ID:         int64(orderEntity.ID),
		CustomerID: orderEntity.CustomerID,
		Status:     orderEntity.Status,
		OrderItems: orderItems,
		CreatedAt:  orderEntity.CreatedAt.UnixNano(),
	}
	return order, res.Error
}

func (a *Adapter) Save(ctx context.Context, order *domain.Order) error {
	var orderItems []OrderItem
	for _, orderItem := range order.OrderItems {
		orderItems = append(orderItems, OrderItem{
			ProducCode: orderItem.ProductCode,
			UnitPrice:  orderItem.UnitPrice,
			Quantity:   orderItem.Quantity,
		})
	}
	orderModel := Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: orderItems,
	}
	res := a.db.Create(&orderModel)
	if res.Error == nil {
		order.ID = int64(orderModel.ID)
	}
	return res.Error
}
