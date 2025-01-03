package api

import (
	"context"
	"errors"
	"testing"

	"github.com/chyiyaqing/gmicro-order/internal/application/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mockedPayment struct {
	mock.Mock
}

func (p *mockedPayment) Charge(ctx context.Context, order *domain.Order) error {
	args := p.Called(ctx, order) // Called() 模拟对象上的方法
	return args.Error(0)
}

type mockedDb struct {
	mock.Mock
}

func (d *mockedDb) Save(ctx context.Context, order *domain.Order) error {
	args := d.Called(ctx, order)
	return args.Error(0)
}

func (d *mockedDb) Get(ctx context.Context, id int64) (domain.Order, error) {
	args := d.Called(ctx, id)
	return args.Get(0).(domain.Order), args.Error(1)
}

type mockedUser struct {
	mock.Mock
}

func (d *mockedUser) Get(ctx context.Context, id int64) (domain.User, error) {
	args := d.Called(ctx, id)
	return args.Get(0).(domain.User), args.Error(1)
}

type mockedShipping struct {
	mock.Mock
}

func (d *mockedShipping) Create(ctx context.Context, order *domain.Order, address string) error {
	args := d.Called(ctx, order, address)
	return args.Error(0)
}

func TestSaveOrder(t *testing.T) {
	payment := new(mockedPayment)
	db := new(mockedDb)
	shipping := new(mockedShipping)
	user := new(mockedUser)
	payment.On("Charge", mock.Anything, mock.Anything).Return(nil)
	db.On("Save", mock.Anything, mock.Anything).Return(nil)
	shipping.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	user.On("Get", mock.Anything, mock.Anything).Return(domain.User{ID: 123}, nil)

	application := NewApplication(db, payment, user, shipping)
	_, err := application.SaveOrder(context.Background(), domain.Order{
		CustomerID: 123,
		OrderItems: []domain.OrderItem{
			{
				ProductCode: "camera",
				UnitPrice:   12.3,
				Quantity:    3,
			},
		},
		CreatedAt: 0,
	})
	assert.Nil(t, err)
}

func Test_Should_Return_Error_When_Db_Persistence_Fail(t *testing.T) {
	payment := new(mockedPayment)
	db := new(mockedDb)
	shipping := new(mockedShipping)
	user := new(mockedUser)
	payment.On("Charge", mock.Anything, mock.Anything).Return(nil)
	db.On("Save", mock.Anything, mock.Anything).Return(errors.New("connection error"))
	shipping.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	user.On("Get", mock.Anything, mock.Anything).Return(domain.User{ID: 123}, nil)

	application := NewApplication(db, payment, user, shipping)
	_, err := application.SaveOrder(context.Background(), domain.Order{
		CustomerID: 123,
		OrderItems: []domain.OrderItem{
			{
				ProductCode: "phone",
				UnitPrice:   12.3,
				Quantity:    3,
			},
		},
		CreatedAt: 0,
	})
	assert.EqualError(t, err, "connection error")
}

func Test_Should_Return_Error_When_Payment_Fail(t *testing.T) {
	payment := new(mockedPayment)
	db := new(mockedDb)
	shipping := new(mockedShipping)
	user := new(mockedUser)
	payment.On("Charge", mock.Anything, mock.Anything).Return(errors.New("insufficient balance"))
	db.On("Save", mock.Anything, mock.Anything).Return(nil)
	shipping.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	user.On("Get", mock.Anything, mock.Anything).Return(domain.User{ID: 123}, nil)

	application := NewApplication(db, payment, user, shipping)
	_, err := application.SaveOrder(context.Background(), domain.Order{
		CustomerID: 123,
		OrderItems: []domain.OrderItem{
			{
				ProductCode: "bag",
				UnitPrice:   2.5,
				Quantity:    6,
			},
		},
		CreatedAt: 0,
	})
	st, _ := status.FromError(err)
	assert.Equal(t, st.Message(), "order creation failed")
	assert.Equal(t, st.Details()[0].(*errdetails.BadRequest).FieldViolations[0].Description, "insufficient balance")
	assert.Equal(t, st.Code(), codes.InvalidArgument)
}
