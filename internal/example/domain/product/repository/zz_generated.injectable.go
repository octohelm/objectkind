/*
Package repository GENERATED BY gengo:injectable
DON'T EDIT THIS FILE
*/
package repository

import (
	context "context"
)

func (v *ProductRepository) Init(ctx context.Context) error {
	if err := v.idGen.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (v *SkuRepository) Init(ctx context.Context) error {
	if err := v.skuID.Init(ctx); err != nil {
		return err
	}

	return nil
}
