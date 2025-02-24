package repository

import (
	"context"
	"iter"

	productv1 "github.com/octohelm/objectkind/internal/example/apis/product/v1"
	"github.com/octohelm/objectkind/internal/example/domain/product"
	productconvert "github.com/octohelm/objectkind/internal/example/domain/product/convert"
	productfilter "github.com/octohelm/objectkind/internal/example/domain/product/filter"
	"github.com/octohelm/objectkind/pkg/sqlutil"
	sqlutilfiller "github.com/octohelm/objectkind/pkg/sqlutil/filler"
	sqlutilquery "github.com/octohelm/objectkind/pkg/sqlutil/query"
	"github.com/octohelm/storage/pkg/dberr"
	"github.com/octohelm/storage/pkg/filter"
	"github.com/octohelm/storage/pkg/sqlpipe"
	sqlpipeex "github.com/octohelm/storage/pkg/sqlpipe/ex"
	iterx "github.com/octohelm/x/iter"
)

type SkuQuerier struct {
	Sku sqlpipeex.Executor[product.Sku]
}

func (q *SkuQuerier) Skus(ctx context.Context, ex sqlpipeex.SourceExecutor[product.Sku]) iter.Seq2[*productv1.Sku, error] {
	return iterx.Action(func(yield func(*productv1.Sku) bool) error {
		products := sqlpipeex.OneToMulti[productv1.ProductID, productv1.Product]{}

		for m, err := range ex.Items(ctx) {
			if err != nil {
				return err
			}

			o, err := productconvert.Sku.ToObject(m)
			if err != nil {
				return err
			}

			if !yield(o) {
				return nil
			}

			if o.Product != nil {
				products.Record(o.Product.ID, o.Product)
			}
		}

		if sqlutilquery.NeedResourceOwner(ctx) {
			if err := sqlutilfiller.FillOwnerSet(ctx, products); err != nil {
				return err
			}
		}

		return nil
	})
}

func (q *SkuQuerier) ListSku(ctx context.Context, operators ...sqlpipe.SourceOperator[product.Sku]) (*productv1.SkuList, error) {
	return sqlutil.List(ctx, q.Sku.PipeE(operators...), q.Skus)
}

func (q *SkuQuerier) FindOneSku(ctx context.Context, operators ...sqlpipe.SourceOperator[product.Sku]) (*productv1.Sku, error) {
	sku, err := sqlutil.FindOne(ctx, q.Sku.PipeE(operators...), q.Skus)
	if err != nil {
		if dberr.IsErrNotFound(err) {
			return nil, &product.ErrSkuNotFound{}
		}
		return nil, err
	}
	return sku, nil
}

func init() {
	sqlutilfiller.Register(&SkuQuerier{})
}

func (q *SkuQuerier) FillSet(ctx context.Context, skus sqlpipeex.Set[productv1.SkuID, productv1.Sku]) error {
	ex := q.Sku.PipeE(&productfilter.SkuByID{
		ID: filter.InSeq(skus.Keys()),
	})
	return sqlutil.FillSet(ctx, skus, ex, q.Skus)
}

func (q *SkuQuerier) FillOwnerSet(ctx context.Context, products sqlpipeex.Set[productv1.ProductID, productv1.Product]) error {
	if products == nil || products.IsZero() {
		return nil
	}

	ex := q.Sku.PipeE(&productfilter.SkuByProductID{
		ProductID: filter.InSeq(products.Keys()),
	})

	for sku, err := range q.Skus(sqlutilquery.With(ctx, sqlutilquery.SkipResourceOwner), ex) {
		if err != nil {
			return err
		}

		if sku.Product != nil {
			for pdt := range products.Records(sku.Product.ID) {
				pdt.Skus = append(pdt.Skus, sku)
			}

			sku.Product = nil
		}
	}

	return nil
}
