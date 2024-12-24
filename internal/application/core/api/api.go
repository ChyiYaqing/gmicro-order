package api

import (
	"context"

	"github.com/chyiyaqing/gmicro-order/internal/application/core/domain"
	"github.com/chyiyaqing/gmicro-order/internal/ports"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db       ports.DBPort
	payment  ports.PaymentPort
	user     ports.UserPort
	shipping ports.ShippingPort
}

// 依赖注入机制：将数据库适配器(特定的数据库技术的具体实现)注入到应用程序中，以便API可以将特定订单的状态存储在数据库中
func NewApplication(db ports.DBPort, payment ports.PaymentPort, user ports.UserPort, shipping ports.ShippingPort) *Application {
	return &Application{
		db:       db,
		payment:  payment,
		user:     user,
		shipping: shipping,
	}
}

func (a *Application) SaveOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	err := a.db.Save(ctx, &order)
	if err != nil {
		return domain.Order{}, err
	}
	// paymentErr
	paymentErr := a.payment.Charge(ctx, &order)
	if paymentErr != nil {
		st, _ := status.FromError(paymentErr)
		fieldErr := &errdetails.BadRequest_FieldViolation{
			Field:       "payment",
			Description: st.Message(),
		}
		badReq := &errdetails.BadRequest{}
		badReq.FieldViolations = append(badReq.FieldViolations, fieldErr)
		orderStatus := status.New(codes.InvalidArgument, "order creation failed")
		statusWithDetail, _ := orderStatus.WithDetails(badReq)
		return domain.Order{}, statusWithDetail.Err()
	}
	// user
	userModel, userErr := a.user.Get(ctx, order.CustomerID)
	if userErr != nil {
		return domain.Order{}, userErr
	}
	// shipping
	shippingErr := a.shipping.Create(ctx, &order, userModel.Address)
	if shippingErr != nil {
		return domain.Order{}, shippingErr
	}

	return order, nil
}

func (a *Application) GetOrder(ctx context.Context, order_id int64) (domain.Order, error) {
	return a.db.Get(ctx, order_id)
}
