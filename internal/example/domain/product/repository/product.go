package repository

import (
	"context"

	"github.com/octohelm/storage/pkg/sqlpipe"

	"github.com/octohelm/objectkind/internal/example/domain/product"
	productconvert "github.com/octohelm/objectkind/internal/example/domain/product/convert"
	productv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/product/v1"
	"github.com/octohelm/objectkind/pkg/idgen"
)

// +gengo:injectable
type ProductRepository struct {
	ProductQuerier

	idGen idgen.Typed[productv1.ProductID]
}

func (repo *ProductRepository) PutProducts(ctx context.Context, products ...*productv1.Product) error {
	if len(products) == 0 {
		return nil
	}

	mProducts := make([]*product.Product, 0, len(products))

	for _, pdt := range products {
		if pdt.ID == 0 {
			if err := repo.idGen.NewTo(&pdt.ID); err != nil {
				return err
			}
		}

		mProduct, err := productconvert.Product.FromObject(pdt)
		if err != nil {
			return err
		}

		mProducts = append(mProducts, mProduct)
	}

	ex := repo.Product.From(sqlpipe.Values(
		mProducts,

		product.ProductT.ID,

		product.ProductT.Name,
		product.ProductT.Description,

		product.ProductT.CreatedAt,
		product.ProductT.UpdatedAt,
		product.ProductT.State,
	)).PipeE(
		sqlpipe.OnConflictDoUpdateSet(
			product.ProductT.I.Primary,

			product.ProductT.Name,
			product.ProductT.Description,
			product.ProductT.State,

			product.ProductT.UpdatedAt,
		),
		sqlpipe.Returning(
			product.ProductT.ID,
			product.ProductT.CreatedAt,
		),
	)

	i := 0
	for x, err := range ex.Items(ctx) {
		if err != nil {
			return err
		}

		products[i].ID = x.ID
		products[i].CreationTimestamp = x.CreatedAt

		i++
	}

	return nil
}

func (repo *ProductRepository) DeleteProducts(ctx context.Context, operators ...sqlpipe.SourceOperator[product.Product]) error {
	return repo.Product.PipeE(operators...).PipeE(sqlpipe.DoDelete[product.Product]()).Commit(ctx)
}
