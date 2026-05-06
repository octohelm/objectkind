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
type ListSkuByProductID struct {
	endpointsv1.ListSkuByProductID
	productService *productservice.ProductService `inject:""`
}

func (op *ListSkuByProductID) Output(ctx context.Context) (any, error) {
	return op.productService.ListSkuByProductID(ctx, op.ProductID, pager.Pager[productdomain.Sku]{
		Offset: op.Offset,
		Limit:  op.Limit,
	})
}

// +gengo:injectable
// +gengo:operator
type GetSkuByID struct {
	endpointsv1.GetSkuByID
	productService *productservice.ProductService `inject:""`
}

func (op *GetSkuByID) Output(ctx context.Context) (any, error) {
	return op.productService.FindOneSku(ctx, &productfilter.SkuByID{
		ID: filter.Eq(op.SkuID),
	})
}

// +gengo:injectable
// +gengo:operator
type CreateSkuByProductID struct {
	endpointsv1.CreateSkuByProductID
	productService *productservice.ProductService `inject:""`
}

func (op *CreateSkuByProductID) Output(ctx context.Context) (any, error) {
	return op.productService.CreateSku(ctx, op.ProductID, op.Body.AsSku())
}

// +gengo:injectable
// +gengo:operator
type UpdateSkuByID struct {
	endpointsv1.UpdateSkuByID
	productService *productservice.ProductService `inject:""`
}

func (op *UpdateSkuByID) Output(ctx context.Context) (any, error) {
	return op.productService.UpdateSku(ctx, op.SkuID, op.Body.AsSku())
}

// +gengo:injectable
// +gengo:operator
type DeleteSkuByID struct {
	endpointsv1.DeleteSkuByID
	productService *productservice.ProductService `inject:""`
}

func (op *DeleteSkuByID) Output(ctx context.Context) (any, error) {
	if err := op.productService.DeleteSku(ctx, op.SkuID); err != nil {
		return nil, err
	}

	return new(courierhttp.NoContent), nil
}
