package repository

import (
	"context"

	productv1 "github.com/octohelm/objectkind/internal/example/apis/product/v1"
	"github.com/octohelm/objectkind/internal/example/domain/product"
	productconvert "github.com/octohelm/objectkind/internal/example/domain/product/convert"
	"github.com/octohelm/objectkind/pkg/idgen"
	"github.com/octohelm/storage/pkg/sqlpipe"
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
		if err := repo.idGen.NewTo(&pdt.ID); err != nil {
			return err
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
	)).PipeE(
		sqlpipe.OnConflictDoUpdateSet(
			product.ProductT.I.Primary,

			product.ProductT.Name,
			product.ProductT.Description,

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
