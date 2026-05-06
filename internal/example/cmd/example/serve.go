package main

import (
	"github.com/innoai-tech/infra/pkg/cli"
	"github.com/innoai-tech/infra/pkg/otel"

	orderservice "github.com/octohelm/objectkind/internal/example/domain/order/service"
	productservice "github.com/octohelm/objectkind/internal/example/domain/product/service"
	"github.com/octohelm/objectkind/pkg/idgen"
)

func init() {
	cli.AddTo(App, &Serve{})
}

// Serve 是 example workspace 的对外服务入口。
type Serve struct {
	cli.C `component:"server"`
	otel.Otel
	idgen.IDGen

	Database Database

	ProductService productservice.ProductService
	OrderService   orderservice.OrderService

	Server APIServer
}
