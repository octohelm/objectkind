package main

import (
	"context"
	"os"

	"github.com/innoai-tech/infra/devpkg/gengo"
	"github.com/innoai-tech/infra/pkg/cli"
	"github.com/innoai-tech/infra/pkg/otel"

	_ "github.com/octohelm/courier/devpkg/clientgen"
	_ "github.com/octohelm/courier/devpkg/injectablegen"
	_ "github.com/octohelm/courier/devpkg/operatorgen"
	_ "github.com/octohelm/courier/devpkg/uintstrgen"
	_ "github.com/octohelm/gengo/devpkg/deepcopygen"
	_ "github.com/octohelm/gengo/devpkg/runtimedocgen"
	_ "github.com/octohelm/storage/devpkg/filteropgen"
	_ "github.com/octohelm/storage/devpkg/tablegen"

	_ "github.com/octohelm/objectkind/devpkg/objectkindgen"
)

var App = cli.NewApp("gengo", "dev")

func init() {
	c := &struct {
		cli.C `name:"gen"`
		otel.Otel
		gengo.Gengo
	}{}

	c.LogLevel = otel.InfoLevel

	cli.AddTo(App, c)
}

func main() {
	if err := cli.Execute(context.Background(), App, os.Args[1:]); err != nil {
		panic(err)
	}
}
