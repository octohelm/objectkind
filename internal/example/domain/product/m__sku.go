package product

import (
	"github.com/octohelm/objectkind/pkg/annotate"
	sqltypecompose "github.com/octohelm/objectkind/pkg/sqltype/compose"
	"github.com/octohelm/storage/pkg/sqltype/json"

	productv1 "github.com/octohelm/objectkind/internal/example/apis/product/v1"
	transactionv1 "github.com/octohelm/objectkind/internal/example/apis/transaction/v1"
)

// Sku
// +gengo:table:register=T
// +gengo:table:group=product
// @def primary ID
// @def unique_index i_sku ProductID Code
type Sku struct {
	sqltypecompose.CodableResource[productv1.SkuID, productv1.SkuCode]
	// 所属产品
	ProductID productv1.ProductID `db:"f_product_id"`
	// 单价
	Price int64 `db:"f_price"`
	// 货币单位
	Currency transactionv1.Currency `db:"f_currency"`
	// 其他属性
	Annotations json.Value[annotate.Annotations] `db:"f_annotations,null"`
}
