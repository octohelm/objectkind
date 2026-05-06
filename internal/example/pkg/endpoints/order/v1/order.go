package v1

import (
	"github.com/octohelm/courier/pkg/courierhttp"

	orderv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/order/v1"
	productv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/product/v1"
)

// ListOrder 查询订单列表。
type ListOrder struct {
	courierhttp.MethodGet `path:"/orders"`
	// 分页偏移
	Offset int64 `name:"offset,omitzero" in:"query"`
	// 分页数
	Limit int64 `name:"limit,omitzero" validate:"@int[-1,50] = 10" in:"query"`
}

func (ListOrder) ResponseData() *orderv1.OrderList { return new(orderv1.OrderList) }

func (ListOrder) ResponseErrors() []error {
	return []error{
		&orderv1.ErrOrderForbidden{},
	}
}

// GetOrderByID 查询单个订单详情。
type GetOrderByID struct {
	courierhttp.MethodGet `path:"/orders/{orderID}"`

	OrderID orderv1.OrderID `name:"orderID" in:"path"`
}

func (GetOrderByID) ResponseData() *orderv1.Order { return new(orderv1.Order) }

func (GetOrderByID) ResponseErrors() []error {
	return []error{
		&orderv1.ErrOrderForbidden{},
		&orderv1.ErrOrderNotFound{},
	}
}

// CreateOrder 创建订单。
type CreateOrder struct {
	courierhttp.MethodPost `path:"/orders"`

	Body orderv1.OrderRequestForCreate `in:"body"`
}

func (CreateOrder) ResponseData() *orderv1.Order { return new(orderv1.Order) }

func (CreateOrder) ResponseErrors() []error {
	return []error{
		&orderv1.ErrOrderForbidden{},
		&productv1.ErrSkuNotFound{},
	}
}

// PayOrderByID 执行订单支付。
type PayOrderByID struct {
	courierhttp.MethodPost `path:"/orders/{orderID}/pay"`

	OrderID orderv1.OrderID            `name:"orderID" in:"path"`
	Body    orderv1.OrderRequestForPay `in:"body"`
}

func (PayOrderByID) ResponseData() *orderv1.Order { return new(orderv1.Order) }

func (PayOrderByID) ResponseErrors() []error {
	return []error{
		&orderv1.ErrOrderForbidden{},
		&orderv1.ErrOrderNotFound{},
		&orderv1.ErrOrderStateConflict{},
	}
}

// CancelOrderByID 执行订单取消。
type CancelOrderByID struct {
	courierhttp.MethodPost `path:"/orders/{orderID}/cancel"`

	OrderID orderv1.OrderID               `name:"orderID" in:"path"`
	Body    orderv1.OrderRequestForCancel `in:"body"`
}

func (CancelOrderByID) ResponseData() *orderv1.Order { return new(orderv1.Order) }

func (CancelOrderByID) ResponseErrors() []error {
	return []error{
		&orderv1.ErrOrderForbidden{},
		&orderv1.ErrOrderNotFound{},
		&orderv1.ErrOrderStateConflict{},
	}
}

// CompleteOrderByID 执行订单完成。
type CompleteOrderByID struct {
	courierhttp.MethodPost `path:"/orders/{orderID}/complete"`

	OrderID orderv1.OrderID                 `name:"orderID" in:"path"`
	Body    orderv1.OrderRequestForComplete `in:"body"`
}

func (CompleteOrderByID) ResponseData() *orderv1.Order { return new(orderv1.Order) }

func (CompleteOrderByID) ResponseErrors() []error {
	return []error{
		&orderv1.ErrOrderForbidden{},
		&orderv1.ErrOrderNotFound{},
		&orderv1.ErrOrderStateConflict{},
	}
}
