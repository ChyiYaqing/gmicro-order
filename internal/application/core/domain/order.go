package domain

import "time"

/**
域对象主要包含字段类型、字段名称和通过标签进行序列化配置的结构体
*/

type OrderItem struct {
	ProductCode string  `json:"product_code"` // 产品唯一编码
	UnitPrice   float32 `json:"unit_price"`   // 单品价格
	Quantity    int32   `json:"quantity"`     // 产品数量
}

type Order struct {
	ID         int64       `json:"id"`          // 订单的唯一标识
	CustomerID int64       `json:"customer_id"` // 订单所有者
	Status     string      `json:"status"`      // 订单状态
	OrderItems []OrderItem `json:"order_items"` // 订单中购买商品清单
	CreatedAt  int64       `json:"created_at"`  // 订单创建时间
}

func NewOrder(customerId int64, orderItems []OrderItem) Order {
	return Order{
		CreatedAt:  time.Now().Unix(),
		Status:     "Pending",
		CustomerID: customerId,
		OrderItems: orderItems,
	}
}

func (o *Order) TotalPrice() float32 {
	var totalPrice float32
	for _, orderItem := range o.OrderItems {
		totalPrice += orderItem.UnitPrice * float32(orderItem.Quantity)
	}
	return totalPrice
}
