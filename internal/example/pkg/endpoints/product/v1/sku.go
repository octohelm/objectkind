package v1

import (
	"github.com/octohelm/courier/pkg/courierhttp"

	productv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/product/v1"
)

// ListSkuByProductID 查询商品下的规格列表。
type ListSkuByProductID struct {
	courierhttp.MethodGet `path:"/products/{productID}/skus"`

	ProductID productv1.ProductID `name:"productID" in:"path"`
	// 分页偏移
	Offset int64 `name:"offset,omitzero" in:"query"`
	// 分页数
	Limit int64 `name:"limit,omitzero" validate:"@int[-1,50] = 10" in:"query"`
}

func (ListSkuByProductID) ResponseData() *productv1.SkuList { return new(productv1.SkuList) }

func (ListSkuByProductID) ResponseErrors() []error {
	return []error{
		&productv1.ErrProductForbidden{},
		&productv1.ErrProductNotFound{},
	}
}

// GetSkuByID 查询单个规格详情。
type GetSkuByID struct {
	courierhttp.MethodGet `path:"/skus/{skuID}"`

	SkuID productv1.SkuID `name:"skuID" in:"path"`
}

func (GetSkuByID) ResponseData() *productv1.Sku { return new(productv1.Sku) }

func (GetSkuByID) ResponseErrors() []error {
	return []error{
		&productv1.ErrProductForbidden{},
		&productv1.ErrSkuNotFound{},
	}
}

// CreateSkuByProductID 为商品创建规格。
type CreateSkuByProductID struct {
	courierhttp.MethodPost `path:"/products/{productID}/skus"`

	ProductID productv1.ProductID           `name:"productID" in:"path"`
	Body      productv1.SkuRequestForCreate `in:"body"`
}

func (CreateSkuByProductID) ResponseData() *productv1.Sku { return new(productv1.Sku) }

func (CreateSkuByProductID) ResponseErrors() []error {
	return []error{
		&productv1.ErrProductForbidden{},
		&productv1.ErrProductNotFound{},
	}
}

// UpdateSkuByID 更新规格。
type UpdateSkuByID struct {
	courierhttp.MethodPut `path:"/skus/{skuID}"`

	SkuID productv1.SkuID               `name:"skuID" in:"path"`
	Body  productv1.SkuRequestForUpdate `in:"body"`
}

func (UpdateSkuByID) ResponseData() *productv1.Sku { return new(productv1.Sku) }

func (UpdateSkuByID) ResponseErrors() []error {
	return []error{
		&productv1.ErrProductForbidden{},
		&productv1.ErrSkuNotFound{},
	}
}

// DeleteSkuByID 删除规格。
type DeleteSkuByID struct {
	courierhttp.MethodDelete `path:"/skus/{skuID}"`

	SkuID productv1.SkuID `name:"skuID" in:"path"`
}

func (DeleteSkuByID) ResponseData() *courierhttp.NoContent { return new(courierhttp.NoContent) }

func (DeleteSkuByID) ResponseErrors() []error {
	return []error{
		&productv1.ErrProductForbidden{},
		&productv1.ErrSkuNotFound{},
	}
}
