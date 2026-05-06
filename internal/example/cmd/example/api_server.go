package main

import (
	"context"

	infrahttp "github.com/innoai-tech/infra/pkg/http"

	"github.com/octohelm/objectkind/internal/example/cmd/example/apis"
)

// APIServer 负责挂载 HTTP 路由并接入 infra server 生命周期。
// +gengo:injectable
type APIServer struct {
	infrahttp.Server
}

func (s *APIServer) beforeInit(ctx context.Context) error {
	s.Server.ApplyRouter(apis.R)
	return nil
}
