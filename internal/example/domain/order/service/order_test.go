package service_test

import (
	"testing"
	"time"

	"github.com/innoai-tech/infra/pkg/configuration/testingutil"
	"github.com/innoai-tech/infra/pkg/otel"
	"github.com/octohelm/storage/pkg/filter"
	sessiondb "github.com/octohelm/storage/pkg/session/db"
	"github.com/octohelm/storage/pkg/sqlbuilder"
	"github.com/octohelm/storage/pkg/sqlpipe"
	sqltypetime "github.com/octohelm/storage/pkg/sqltype/time"
	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"

	orderdomain "github.com/octohelm/objectkind/internal/example/domain/order"
	orderdomainfilter "github.com/octohelm/objectkind/internal/example/domain/order/filter"
	orderrepository "github.com/octohelm/objectkind/internal/example/domain/order/repository"
	orderservice "github.com/octohelm/objectkind/internal/example/domain/order/service"
	productdomain "github.com/octohelm/objectkind/internal/example/domain/product"
	productrepository "github.com/octohelm/objectkind/internal/example/domain/product/repository"
	productservice "github.com/octohelm/objectkind/internal/example/domain/product/service"
	orderv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/order/v1"
	productv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/product/v1"
	transactionv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/transaction/v1"
	"github.com/octohelm/objectkind/pkg/idgen"
	"github.com/octohelm/objectkind/pkg/runtime"
	"github.com/octohelm/objectkind/pkg/sqlutil/pager"
)

func TestOrderService(t *testing.T) {
	ctx, d := testingutil.BuildContext(t, func(d *struct {
		otel.Otel
		idgen.IDGen
		sessiondb.Database
		productrepository.ProductRepository
		productrepository.SkuRepository
		productservice.ProductService
		orderrepository.OrderRepository
		orderrepository.OrderItemRepository
		orderservice.OrderService
	},
	) {
		d.LogLevel = "debug"
		d.LogFormat = "text"
		d.EnableMigrate = true
		d.ApplyCatalog("test", productdomain.T, orderdomain.T)
	})

	productSvc := &d.ProductService
	orderSvc := &d.OrderService

	product := runtime.Build(func(p *productv1.Product) {
		p.Name = "下单商品"
	})

	sku := runtime.Build(func(s *productv1.Sku) {
		s.Code = productv1.SkuCode("order-sku-1")
		s.Name = "order-sku-1"
		s.Spec.Price = 100
		s.Spec.Currency = transactionv1.CurrencyCNY
	})

	Must(t, func() error {
		return productSvc.UpsertProduct(ctx, product, sku)
	})

	orderRequest := runtime.Build(func(r *orderv1.OrderRequestForCreate) {
		r.Name = "订单-1"
		r.Spec.Items = []*orderv1.OrderItemRequestForCreate{
			{
				Spec: orderv1.OrderItemSpecRequestForCreate{
					Quantity: 2,
				},
			},
		}
		r.Spec.Items[0].Spec.Sku.ID = sku.ID
	})

	created := MustValue(t, func() (*orderv1.Order, error) {
		return orderSvc.CreateOrder(ctx, orderRequest)
	})

	Then(
		t, "创建订单后金额与商品明细正确",
		Expect(created.Status.State, Equal(orderv1.ORDER_STATE__CREATED)),
		Expect(created.Status.TotalAmount, Equal(200)),
		Expect(created.Spec.Items, Be(cmp.Len[[]*orderv1.OrderItem](1))),
		Expect(created.Spec.Items[0].Spec.Sku.ID, Equal(sku.ID)),
		Expect(created.Spec.Items[0].Status.FinalPrice, Equal(float64(200))),
	)

	paid := MustValue(t, func() (*orderv1.Order, error) {
		return orderSvc.PayOrder(ctx, created.ID, orderv1.OrderPaymentChannel("cashier"))
	})

	Then(
		t, "支付后状态变更为已支付",
		Expect(paid.Status.State, Equal(orderv1.ORDER_STATE__PAID)),
	)

	completed := MustValue(t, func() (*orderv1.Order, error) {
		return orderSvc.CompleteOrder(ctx, created.ID)
	})

	Then(
		t, "完成后状态变更为已完成",
		Expect(completed.Status.State, Equal(orderv1.ORDER_STATE__COMPLETED)),
	)

	cancelCandidate := MustValue(t, func() (*orderv1.Order, error) {
		return orderSvc.CreateOrder(ctx, runtime.Build(func(r *orderv1.OrderRequestForCreate) {
			r.Name = "订单-2"
			r.Spec.Items = []*orderv1.OrderItemRequestForCreate{
				{
					Spec: orderv1.OrderItemSpecRequestForCreate{
						Quantity: 1,
					},
				},
			}
			r.Spec.Items[0].Spec.Sku.ID = sku.ID
		}))
	})

	canceled := MustValue(t, func() (*orderv1.Order, error) {
		return orderSvc.CancelOrder(ctx, cancelCandidate.ID, orderv1.OrderCancelReason("user_request"))
	})

	Then(
		t, "取消后状态变更为已取消",
		Expect(canceled.Status.State, Equal(orderv1.ORDER_STATE__CANCELED)),
	)

	t.Run("列表查询返回完整订单上下文", func(t *testing.T) {
		list := MustValue(t, func() (*orderv1.OrderList, error) {
			return orderSvc.ListOrder(ctx, &pager.Pager[orderdomain.Order]{
				Offset: 0,
				Limit:  10,
			})
		})

		Then(
			t, "订单列表包含订单项与商品规格",
			Expect(list.Items, Be(cmp.Len[[]*orderv1.Order](2))),
			Expect(list.Items[0].Spec.Items, Be(cmp.NotNil[[]*orderv1.OrderItem]())),
			Expect(list.Items[0].Spec.Items[0].Spec.Sku, Be(cmp.NotNil[*productv1.Sku]())),
		)
	})

	t.Run("关闭超时未支付订单", func(t *testing.T) {
		expiredCandidate := MustValue(t, func() (*orderv1.Order, error) {
			return orderSvc.CreateOrder(ctx, runtime.Build(func(r *orderv1.OrderRequestForCreate) {
				r.Name = "订单-过期"
				r.Spec.Items = []*orderv1.OrderItemRequestForCreate{
					{
						Spec: orderv1.OrderItemSpecRequestForCreate{
							Quantity: 1,
						},
					},
				}
				r.Spec.Items[0].Spec.Sku.ID = sku.ID
			}))
		})

		Must(t, func() error {
			return orderSvc.OrderRepository.Order.PipeE(
				&orderdomainfilter.OrderByID{
					ID: filter.Eq(expiredCandidate.ID),
				},
				sqlpipe.DoUpdate(orderdomain.OrderT.CreatedAt, sqlbuilder.Value(sqltypetime.Add(expiredCandidate.CreationTimestamp, -2*time.Hour))),
			).Commit(ctx)
		})

		closed := MustValue(t, func() (int, error) {
			return orderSvc.CloseExpiredOrders(ctx, time.Now().Add(-30*time.Minute))
		})

		reloaded := MustValue(t, func() (*orderv1.Order, error) {
			return orderSvc.GetOrder(ctx, expiredCandidate.ID)
		})

		Then(
			t, "超时未支付订单被自动关闭",
			Expect(closed, Equal(1)),
			Expect(reloaded.Status.State, Equal(orderv1.ORDER_STATE__CANCELED)),
		)
	})
}
