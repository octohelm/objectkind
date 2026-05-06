package main

import (
	"context"

	sessiondb "github.com/octohelm/storage/pkg/session/db"

	orderdomain "github.com/octohelm/objectkind/internal/example/domain/order"
	productdomain "github.com/octohelm/objectkind/internal/example/domain/product"
)

// +gengo:injectable:provider
type Database struct {
	sessiondb.Database
}

func (d *Database) SetDefaults() {
	if d.Database.NameOverwrite == "" {
		d.Database.NameOverwrite = "example"
	}

	d.Database.SetDefaults()

	d.Database.ApplyCatalog(
		d.Database.DBName(),

		productdomain.T,
		orderdomain.T,
	)
}

func (s *Database) Run(ctx context.Context) error {
	if err := s.Database.Run(ctx); err != nil {
		return err
	}
	return nil
}
