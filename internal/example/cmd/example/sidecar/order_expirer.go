package sidecar

import (
	"context"
	"time"

	"github.com/innoai-tech/infra/pkg/agent"
	"github.com/innoai-tech/infra/pkg/cron"
	"github.com/octohelm/x/sync/singleflight"

	orderservice "github.com/octohelm/objectkind/internal/example/domain/order/service"
)

// OrderExpirer 定期关闭超过阈值仍未支付的订单。
// +gengo:injectable
type OrderExpirer struct {
	agent.Agent

	Period cron.Spec `flag:",omitzero"`
	TTL    Duration  `flag:",omitzero"`

	orderService *orderservice.OrderService `inject:""`
}

// SetDefaults 设置默认调度周期与订单过期阈值。
func (a *OrderExpirer) SetDefaults() {
	if a.Period == "" {
		a.Period = "*/30 * * * *"
	}

	if a.TTL == 0 {
		a.TTL = Duration(30 * time.Minute)
	}
}

// Disabled 返回当前 agent 是否应被禁用。
func (a *OrderExpirer) Disabled(ctx context.Context) bool {
	return a.Period.Schedule() == nil
}

// afterInit 注册定期关闭过期订单的后台 worker。
func (a *OrderExpirer) afterInit(ctx context.Context) error {
	if a.Disabled(ctx) {
		return nil
	}

	sfg := &singleflight.Group[string]{}

	a.Host("CloseExpiredOrders", func(ctx context.Context) error {
		for range a.Period.Times(ctx) {
			a.Go(ctx, func(ctx context.Context) error {
				err, _ := sfg.Do("close-expired-orders", func() error {
					defer sfg.Forget("close-expired-orders")

					_, err := a.orderService.CloseExpiredOrders(ctx, time.Now().Add(-time.Duration(a.TTL)))
					return err
				})

				return err
			})
		}

		return nil
	})

	return nil
}

// Duration 是 sidecar 场景下用于 flag/配置解析的时间长度包装。
type Duration time.Duration

func (d Duration) String() string {
	return time.Duration(d).String()
}

func (d *Duration) UnmarshalText(text []byte) error {
	v, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}

	*d = Duration(v)
	return nil
}
