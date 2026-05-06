package apis

import (
	"github.com/octohelm/courier/pkg/courierhttp"

	"github.com/octohelm/objectkind/internal/example/cmd/example/apis/order"
	"github.com/octohelm/objectkind/internal/example/cmd/example/apis/product"
)

var R = courierhttp.GroupRouter("/api/example").With(
	courierhttp.GroupRouter("/v1").With(
		order.R,
		product.R,
	),
)
