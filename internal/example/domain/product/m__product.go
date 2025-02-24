package product

import (
	productv1 "github.com/octohelm/objectkind/internal/example/apis/product/v1"
	sqltypecompose "github.com/octohelm/objectkind/pkg/sqltype/compose"
)

// Product
// +gengo:table:register=T
// +gengo:table:group=product
// @def primary ID
// @def index i_state State
type Product struct {
	sqltypecompose.Resource[productv1.ProductID]

	State productv1.ProductState `db:"state,default=1"`
}
