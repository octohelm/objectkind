# API Server / Agent 组装层

## 目录

- `cmd/<app>/api_server.go` — HTTP server 挂载
- `cmd/<app>/apis/<domain>/` — operator 适配
- `cmd/<app>/sidecar/` — 后台 agent

## api-server

`api_server.go` 负责挂载路由并接入 infra server 生命周期。具体用法见 `infra-guideline`。

## operator

`cmd/<app>/apis/<domain>/` 负责最薄适配：

- embed endpoint 定义
- 注入 `*Service`
- 在 `Output(ctx)` 转发参数

## sidecar / agent

`cmd/<app>/sidecar/` 负责后台任务：

- 注入 service
- 声明调度周期
- 复用 service 能力，不直连 table / repository
- 具体用法见 `infra-guideline`

## 准则

- operator 不重复声明 path、method、response 契约
- operator 不写业务状态机
- sidecar 只调用 service，不直接操作数据库细节
- server 和 agent 都通过 injectable / singleton 生命周期接入依赖

## 反模式

- 在 operator 里手写业务编排
- sidecar 直接拼 SQL 或直接操作 repository 明细
- server 初始化阶段顺手写业务逻辑
