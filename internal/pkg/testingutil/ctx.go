package testingutil

import (
	"context"
	"os"
	"testing"

	"github.com/innoai-tech/infra/pkg/configuration"
	"golang.org/x/sync/errgroup"
)

func NewContext(t testing.TB, v any) context.Context {
	tmp := t.TempDir()
	_ = os.Chdir(tmp)

	t.Cleanup(func() {
		_ = os.RemoveAll(tmp)
	})

	ctx := context.Background()
	if v != nil {
		singletons := configuration.SingletonsFromStruct(v)
		c, err := singletons.Init(ctx)
		if err != nil {
			panic(err)
		}

		ctx = c

		for s := range singletons.Configurators() {
			if r, ok := s.(configuration.Runner); ok {
				err := r.Run(ctx)
				if err != nil {
					panic(err)
				}
			}
		}

		go func() {
			g, c := errgroup.WithContext(ctx)

			for s := range singletons.Configurators() {
				if server, ok := s.(configuration.Server); ok {
					g.Go(func() error {
						err := server.Serve(c)
						return err
					})
				}
			}

			_ = g.Wait()
		}()

		t.Cleanup(func() {
			c := configuration.ContextInjectorFromContext(ctx).InjectContext(ctx)

			for s := range singletons.Configurators() {
				if canShutdown, ok := s.(configuration.CanShutdown); ok {
					_ = configuration.Shutdown(c, canShutdown)
				}
			}
		})
	}

	return configuration.ContextInjectorFromContext(ctx).InjectContext(ctx)
}
