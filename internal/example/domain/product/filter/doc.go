//go:generate go tool devtool gen .
package filter

import (
	"github.com/octohelm/objectkind/internal/example/domain/product"
)

// +gengo:filterop
type filterOfProduct struct {
	product.Product
}

// +gengo:filterop
type filterOfSku struct {
	product.Sku
}
