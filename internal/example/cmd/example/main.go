package main

import (
	"context"

	"github.com/innoai-tech/infra/pkg/cli"
)

var App = cli.NewApp("example", "1.0.0")

// main 启动 example workspace 的 CLI 应用。
func main() {
	cli.Exec(context.Background(), App)
}
