package service

import (
	"context"

	"github.com/octohelm/storage/pkg/filter"

	"github.com/octohelm/objectkind/internal/example/domain/product"
	productfilter "github.com/octohelm/objectkind/internal/example/domain/product/filter"
	productrepository "github.com/octohelm/objectkind/internal/example/domain/product/repository"
	productv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/product/v1"
	"github.com/octohelm/objectkind/pkg/sqlutil/pager"
)

// +gengo:injectable:provider
type ProductService struct {
	productrepository.ProductRepository
	productrepository.SkuRepository
}

func (svc *ProductService) UpsertProduct(ctx context.Context, product *productv1.Product, skus ...*productv1.Sku) error {
	if err := svc.ProductRepository.PutProducts(ctx, product); err != nil {
		return err
	}

	for _, sku := range skus {
		sku.Product = product
	}

	return svc.SkuRepository.PutSkus(ctx, skus...)
}

func (svc *ProductService) ListSkuByProductID(ctx context.Context, productID productv1.ProductID, p pager.Pager[product.Sku]) (*productv1.SkuList, error) {
	if _, err := svc.ProductRepository.FindOneProduct(ctx, &productfilter.ProductByID{
		ID: filter.Eq(productID),
	}); err != nil {
		return nil, err
	}

	return svc.SkuRepository.ListSku(ctx,
		&productfilter.SkuByProductID{
			ProductID: filter.Eq(productID),
		},
		&p,
	)
}

func (svc *ProductService) CreateProduct(ctx context.Context, product *productv1.Product) (*productv1.Product, error) {
	if err := svc.ProductRepository.PutProducts(ctx, product); err != nil {
		return nil, err
	}

	return svc.ProductRepository.FindOneProduct(ctx, &productfilter.ProductByID{
		ID: filter.Eq(product.ID),
	})
}

func (svc *ProductService) UpdateProduct(ctx context.Context, id productv1.ProductID, product *productv1.Product) (*productv1.Product, error) {
	current, err := svc.ProductRepository.FindOneProduct(ctx, &productfilter.ProductByID{
		ID: filter.Eq(id),
	})
	if err != nil {
		return nil, err
	}

	product.ID = current.ID
	product.CreationTimestamp = current.CreationTimestamp
	product.Status = current.Status

	if err := svc.ProductRepository.PutProducts(ctx, product); err != nil {
		return nil, err
	}

	return svc.ProductRepository.FindOneProduct(ctx, &productfilter.ProductByID{
		ID: filter.Eq(id),
	})
}

func (svc *ProductService) DeleteProduct(ctx context.Context, id productv1.ProductID) error {
	if _, err := svc.ProductRepository.FindOneProduct(ctx, &productfilter.ProductByID{
		ID: filter.Eq(id),
	}); err != nil {
		return err
	}

	if err := svc.SkuRepository.DeleteSkus(ctx, &productfilter.SkuByProductID{
		ProductID: filter.Eq(id),
	}); err != nil {
		return err
	}

	return svc.ProductRepository.DeleteProducts(ctx, &productfilter.ProductByID{
		ID: filter.Eq(id),
	})
}

func (svc *ProductService) PublishProduct(ctx context.Context, id productv1.ProductID) (*productv1.Product, error) {
	product, err := svc.ProductRepository.FindOneProduct(ctx, &productfilter.ProductByID{
		ID: filter.Eq(id),
	})
	if err != nil {
		return nil, err
	}

	if product.Status.State != productv1.PRODUCT_STATE__OFF_SALE {
		return nil, &productv1.ErrProductStateConflict{}
	}

	product.Status.State = productv1.PRODUCT_STATE__ON_SALE

	if err := svc.ProductRepository.PutProducts(ctx, product); err != nil {
		return nil, err
	}

	return svc.ProductRepository.FindOneProduct(ctx, &productfilter.ProductByID{
		ID: filter.Eq(id),
	})
}

func (svc *ProductService) UnpublishProduct(ctx context.Context, id productv1.ProductID) (*productv1.Product, error) {
	product, err := svc.ProductRepository.FindOneProduct(ctx, &productfilter.ProductByID{
		ID: filter.Eq(id),
	})
	if err != nil {
		return nil, err
	}

	if product.Status.State != productv1.PRODUCT_STATE__ON_SALE {
		return nil, &productv1.ErrProductStateConflict{}
	}

	product.Status.State = productv1.PRODUCT_STATE__OFF_SALE

	if err := svc.ProductRepository.PutProducts(ctx, product); err != nil {
		return nil, err
	}

	return svc.ProductRepository.FindOneProduct(ctx, &productfilter.ProductByID{
		ID: filter.Eq(id),
	})
}

func (svc *ProductService) CreateSku(ctx context.Context, productID productv1.ProductID, sku *productv1.Sku) (*productv1.Sku, error) {
	product, err := svc.ProductRepository.FindOneProduct(ctx, &productfilter.ProductByID{
		ID: filter.Eq(productID),
	})
	if err != nil {
		return nil, err
	}

	sku.Product = product

	if err := svc.SkuRepository.PutSkus(ctx, sku); err != nil {
		return nil, err
	}

	return svc.SkuRepository.FindOneSku(ctx, &productfilter.SkuByID{
		ID: filter.Eq(sku.ID),
	})
}

func (svc *ProductService) UpdateSku(ctx context.Context, id productv1.SkuID, sku *productv1.Sku) (*productv1.Sku, error) {
	current, err := svc.SkuRepository.FindOneSku(ctx, &productfilter.SkuByID{
		ID: filter.Eq(id),
	})
	if err != nil {
		return nil, err
	}

	sku.ID = current.ID
	sku.Code = current.Code
	sku.CreationTimestamp = current.CreationTimestamp
	sku.Product = current.Product

	if err := svc.SkuRepository.PutSkus(ctx, sku); err != nil {
		return nil, err
	}

	return svc.SkuRepository.FindOneSku(ctx, &productfilter.SkuByID{
		ID: filter.Eq(id),
	})
}

func (svc *ProductService) DeleteSku(ctx context.Context, id productv1.SkuID) error {
	if _, err := svc.SkuRepository.FindOneSku(ctx, &productfilter.SkuByID{
		ID: filter.Eq(id),
	}); err != nil {
		return err
	}

	return svc.SkuRepository.DeleteSkus(ctx, &productfilter.SkuByID{
		ID: filter.Eq(id),
	})
}
