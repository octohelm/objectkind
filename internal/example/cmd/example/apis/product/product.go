package product

import (
	"context"

	"github.com/octohelm/courier/pkg/courierhttp"
	"github.com/octohelm/storage/pkg/filter"

	productdomain "github.com/octohelm/objectkind/internal/example/domain/product"
	productfilter "github.com/octohelm/objectkind/internal/example/domain/product/filter"
	productservice "github.com/octohelm/objectkind/internal/example/domain/product/service"
	endpointsv1 "github.com/octohelm/objectkind/internal/example/pkg/endpoints/product/v1"
	"github.com/octohelm/objectkind/pkg/sqlutil/pager"
)

// +gengo:injectable
// +gengo:operator
type ListProduct struct {
	endpointsv1.ListProduct
	productService *productservice.ProductService `inject:""`
}

func (op *ListProduct) Output(ctx context.Context) (any, error) {
	return op.productService.ListProduct(ctx, &pager.Pager[productdomain.Product]{
		Offset: op.Offset,
		Limit:  op.Limit,
	})
}

// +gengo:injectable
// +gengo:operator
type GetProductByID struct {
	endpointsv1.GetProductByID
	productService *productservice.ProductService `inject:""`
}

func (op *GetProductByID) Output(ctx context.Context) (any, error) {
	return op.productService.FindOneProduct(ctx, &productfilter.ProductByID{
		ID: filter.Eq(op.ProductID),
	})
}

// +gengo:injectable
// +gengo:operator
type CreateProduct struct {
	endpointsv1.CreateProduct
	productService *productservice.ProductService `inject:""`
}

func (op *CreateProduct) Output(ctx context.Context) (any, error) {
	return op.productService.CreateProduct(ctx, op.Body.AsProduct())
}

// +gengo:injectable
// +gengo:operator
type UpdateProductByID struct {
	endpointsv1.UpdateProductByID
	productService *productservice.ProductService `inject:""`
}

func (op *UpdateProductByID) Output(ctx context.Context) (any, error) {
	return op.productService.UpdateProduct(ctx, op.ProductID, op.Body.AsProduct())
}

// +gengo:injectable
// +gengo:operator
type DeleteProductByID struct {
	endpointsv1.DeleteProductByID
	productService *productservice.ProductService `inject:""`
}

func (op *DeleteProductByID) Output(ctx context.Context) (any, error) {
	if err := op.productService.DeleteProduct(ctx, op.ProductID); err != nil {
		return nil, err
	}

	return new(courierhttp.NoContent), nil
}

// +gengo:injectable
// +gengo:operator
type PublishProductByID struct {
	endpointsv1.PublishProductByID
	productService *productservice.ProductService `inject:""`
}

func (op *PublishProductByID) Output(ctx context.Context) (any, error) {
	return op.productService.PublishProduct(ctx, op.ProductID)
}

// +gengo:injectable
// +gengo:operator
type UnpublishProductByID struct {
	endpointsv1.UnpublishProductByID
	productService *productservice.ProductService `inject:""`
}

func (op *UnpublishProductByID) Output(ctx context.Context) (any, error) {
	return op.productService.UnpublishProduct(ctx, op.ProductID)
}
