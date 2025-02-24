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
type SkuRepository struct {
	SkuQuerier

	skuID idgen.Typed[productv1.SkuID]
}

func (repo *SkuRepository) PutSkus(ctx context.Context, skus ...*productv1.Sku) error {
	if len(skus) == 0 {
		return nil
	}

	mSkus := make([]*product.Sku, 0)

	for _, sku := range skus {
		if err := repo.skuID.NewTo(&sku.ID); err != nil {
			return err
		}

		mSku, err := productconvert.Sku.FromObject(sku)
		if err != nil {
			return err
		}

		mSkus = append(mSkus, mSku)
	}

	ex := repo.Sku.From(sqlpipe.Values(mSkus)).PipeE(
		sqlpipe.OnConflictDoUpdateSet(
			product.SkuT.I.ISku,

			product.SkuT.Name,
			product.SkuT.Description,

			product.SkuT.Price,
			product.SkuT.Currency,

			product.SkuT.UpdatedAt,
		),
		sqlpipe.Returning(
			product.SkuT.ID,
		),
	)

	i := 0
	for m, err := range ex.Items(ctx) {
		if err != nil {
			return err
		}

		skus[i].ID = m.ID
		i++
	}

	return nil
}
