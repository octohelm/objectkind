package repository_test

import (
	"testing"

	"github.com/innoai-tech/infra/pkg/configuration/testingutil"
	"github.com/innoai-tech/infra/pkg/otel"
	"github.com/octohelm/storage/pkg/filter"
	sessiondb "github.com/octohelm/storage/pkg/session/db"
	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"

	orderdomain "github.com/octohelm/objectkind/internal/example/domain/order"
	orderfilter "github.com/octohelm/objectkind/internal/example/domain/order/filter"
	orderrepository "github.com/octohelm/objectkind/internal/example/domain/order/repository"
	orderv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/order/v1"
	productv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/product/v1"
	"github.com/octohelm/objectkind/pkg/idgen"
	"github.com/octohelm/objectkind/pkg/runtime"
)

func TestOrderRepository(t *testing.T) {
	d := &struct {
		otel.Otel
		idgen.IDGen
		sessiondb.Database
		orderrepository.OrderRepository
		orderrepository.OrderItemRepository
	}{}

	d.LogLevel = "debug"
	d.LogFormat = "text"
	d.EnableMigrate = true
	d.ApplyCatalog("test", orderdomain.T)

	ctx := testingutil.NewContext(t, d)

	order := runtime.Build(func(o *orderv1.Order) {
		o.Name = "测试订单"
		o.Status.State = orderv1.ORDER_STATE__CREATED
		o.Status.TotalAmount = 99
	})

	Must(t, func() error {
		return d.PutOrders(ctx, order)
	})

	item := runtime.Build(func(i *orderv1.OrderItem) {
		i.Spec.Sku = &productv1.Sku{}
		i.Spec.Sku.ID = 1001
		i.Spec.Quantity = 2
		i.Status.TotalPrice = 198
		i.Status.FinalPrice = 198
	})

	item.Name = "订单项-1"
	item.Description = "测试订单项"

	Must(t, func() error {
		return d.PutOrderItems(ctx, order.ID, item)
	})

	got := MustValue(t, func() (*orderv1.Order, error) {
		return d.FindOneOrder(ctx, &orderfilter.OrderByID{
			ID: filter.Eq(order.ID),
		})
	})

	Then(t, "订单及订单项正确",
		Expect(got.ID, Equal(order.ID)),
		Expect(got.Spec.Items, Be(cmp.Len[[]*orderv1.OrderItem](1))),
		Expect(got.Spec.Items[0].Spec.Sku.ID, Equal(productv1.SkuID(1001))),
		Expect(got.Spec.Items[0].Spec.Quantity, Equal(int64(2))),
	)
}
