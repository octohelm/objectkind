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

type ProductQuerier struct {
	Product sqlpipeex.Executor[product.Product]
}

func (q *ProductQuerier) Products(ctx context.Context, ex sqlpipeex.SourceExecutor[product.Product]) iter.Seq2[*productv1.Product, error] {
	return iterx.Action(func(yield func(*productv1.Product) bool) error {
		products := sqlpipeex.OneToOne[productv1.ProductID, productv1.Product]{}

		for m, err := range ex.Items(ctx) {
			if err != nil {
				return err
			}

			pdt, err := productconvert.Product.ToObject(m)
			if err != nil {
				return err
			}

			if !yield(pdt) {
				return nil
			}

			products.Record(pdt.ID, pdt)
		}

		if sqlutilquery.NeedSubResources(ctx) {
			if err := sqlutilfiller.FillSubResourcesOfOwnerSet(ctx, products); err != nil {
				return err
			}
		}

		return nil
	})
}

func (q *ProductQuerier) ListProduct(ctx context.Context, operators ...sqlpipe.SourceOperator[product.Product]) (*productv1.ProductList, error) {
	return sqlutil.List(ctx, q.Product.PipeE(operators...), q.Products)
}

func (q *ProductQuerier) FindOneProduct(ctx context.Context, operators ...sqlpipe.SourceOperator[product.Product]) (*productv1.Product, error) {
	epr, err := sqlutil.FindOne(ctx, q.Product.PipeE(operators...), q.Products)
	if err != nil {
		if dberr.IsErrNotFound(err) {
			return nil, &product.ErrProductNotFound{}
		}
		return nil, err
	}
	return epr, nil
}

func init() {
	sqlutilfiller.Register(&ProductQuerier{})
}

func (q *ProductQuerier) FillSet(ctx context.Context, products sqlpipeex.Set[productv1.ProductID, productv1.Product]) error {
	ex := q.Product.PipeE(&productfilter.ProductByID{
		ID: filter.InSeq(products.Keys()),
	})
	return sqlutil.FillSet(ctx, products, ex, q.Products)
}
