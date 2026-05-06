package v1

import (
	"github.com/octohelm/courier/pkg/courierhttp"

	productv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/product/v1"
)

// ListProduct 查询商品列表。
type ListProduct struct {
	courierhttp.MethodGet `path:"/products"`

	// 分页偏移
	Offset int64 `name:"offset,omitzero" in:"query"`
	// 分页数
	Limit int64 `name:"limit,omitzero" validate:"@int[-1,50] = 10" in:"query"`
}

func (ListProduct) ResponseData() *productv1.ProductList { return new(productv1.ProductList) }

func (ListProduct) ResponseErrors() []error {
	return []error{
		&productv1.ErrProductForbidden{},
	}
}

// GetProductByID 查询单个商品详情。
type GetProductByID struct {
	courierhttp.MethodGet `path:"/products/{productID}"`

	ProductID productv1.ProductID `name:"productID" in:"path"`
}

func (GetProductByID) ResponseData() *productv1.Product { return new(productv1.Product) }

func (GetProductByID) ResponseErrors() []error {
	return []error{
		&productv1.ErrProductForbidden{},
		&productv1.ErrProductNotFound{},
	}
}

// CreateProduct 创建商品。
type CreateProduct struct {
	courierhttp.MethodPost `path:"/products"`

	Body productv1.ProductRequestForCreate `in:"body"`
}

func (CreateProduct) ResponseData() *productv1.Product { return new(productv1.Product) }

func (CreateProduct) ResponseErrors() []error {
	return []error{
		&productv1.ErrProductForbidden{},
	}
}

// UpdateProductByID 更新商品。
type UpdateProductByID struct {
	courierhttp.MethodPut `path:"/products/{productID}"`

	ProductID productv1.ProductID               `name:"productID" in:"path"`
	Body      productv1.ProductRequestForUpdate `in:"body"`
}

func (UpdateProductByID) ResponseData() *productv1.Product { return new(productv1.Product) }

func (UpdateProductByID) ResponseErrors() []error {
	return []error{
		&productv1.ErrProductForbidden{},
		&productv1.ErrProductNotFound{},
	}
}

// DeleteProductByID 删除商品。
type DeleteProductByID struct {
	courierhttp.MethodDelete `path:"/products/{productID}"`

	ProductID productv1.ProductID `name:"productID" in:"path"`
}

func (DeleteProductByID) ResponseData() *courierhttp.NoContent { return new(courierhttp.NoContent) }

func (DeleteProductByID) ResponseErrors() []error {
	return []error{
		&productv1.ErrProductForbidden{},
		&productv1.ErrProductNotFound{},
	}
}

// PublishProductByID 执行商品上架。
type PublishProductByID struct {
	courierhttp.MethodPost `path:"/products/{productID}/publish"`

	ProductID productv1.ProductID                `name:"productID" in:"path"`
	Body      productv1.ProductRequestForPublish `in:"body"`
}

func (PublishProductByID) ResponseData() *productv1.Product { return new(productv1.Product) }

func (PublishProductByID) ResponseErrors() []error {
	return []error{
		&productv1.ErrProductForbidden{},
		&productv1.ErrProductNotFound{},
		&productv1.ErrProductStateConflict{},
	}
}

// UnpublishProductByID 执行商品下架。
type UnpublishProductByID struct {
	courierhttp.MethodPost `path:"/products/{productID}/unpublish"`

	ProductID productv1.ProductID                  `name:"productID" in:"path"`
	Body      productv1.ProductRequestForUnpublish `in:"body"`
}

func (UnpublishProductByID) ResponseData() *productv1.Product { return new(productv1.Product) }

func (UnpublishProductByID) ResponseErrors() []error {
	return []error{
		&productv1.ErrProductForbidden{},
		&productv1.ErrProductNotFound{},
		&productv1.ErrProductStateConflict{},
	}
}
