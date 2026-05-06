---
name: objectkind-guideline
description: 开发或修改 Go 后端服务时使用；覆盖从契约定义、数据操作、CLI 单例、api-server/sidecar 组装到测试验证的完整流程。
---

# ObjectKind Guideline

## 依赖 skill

本 skill 只描述分层方法论与开发流水线。具体工具/框架约定由对应 skill 负责，不在此重复：

| 涉及内容 | 加载 skill |
|---------|-----------|
| `+gengo:*` 注解与生成器 | `gengo-guideline` |
| `+gengo:enum` 枚举类型 | `enumeration-guideline` |
| courier endpoint / method / path | `courier-guideline` |
| storage 表结构 / repository / filter | `storage-guideline` |
| infra CLI / server / agent / 单例 | `infra-guideline` |
| `testing/v2` 断言与实例构造 | `testing-guideline` |

不在上述列表的工具/框架约定，优先查 `go doc` 或对应包源码，不在本 skill 附加说明。

## 分层模型

```
<module>/
├── pkg/apis/<domain>/v1/          # 对象、request、reference、list、错误
├── pkg/endpoints/<domain>/v1/     # method、path、输入输出
├── domain/<domain>/               # 表模型、convert、filter、repository、service
└── cmd/<app>/                     # CLI 入口、DI、server、operator、sidecar
```

依赖方向：`pkg/apis` ← `pkg/endpoints` ← `domain` ← `cmd/<app>`

→ 详见 [总览](references/overview.md)

## 开发流水线

严格按以下环节推进。每步只在该层需要时才打开对应 reference。

### 1. 契约层 — 先定义稳定语义

- `pkg/apis/<domain>/v1/`：先落对象，再落 request、list、错误
- `pkg/endpoints/<domain>/v1/`：再落 method、path、`ResponseData()`、`ResponseErrors()`
- endpoint 只声明契约，不写业务逻辑

→ [contract-layer.md](references/contract-layer.md)

### 2. 数据层 — 再写操作与编排

- `domain/<domain>/`：表模型、convert、filter
- `domain/<domain>/repository/`：单 domain 数据访问、批量填充
- `domain/<domain>/service/`：业务编排、状态流转、跨 domain 组合
- 列表/聚合/关联补全优先用 `FillSet`、`filler.Fill`、`FillOwnerSet`，禁止循环逐条 `FindOne`

→ [data-layer.md](references/data-layer.md)

### 3. CLI 入口 — 依赖注入与基础设施

- `cmd/<app>/main.go`、`serve.go`、`database.go`
- provider 通过 `+gengo:injectable` 生命周期接入，不手拼对象图
- 基础设施单例在 `cmd/<app>` 层声明，不散落到 domain

→ [cli-singletons.md](references/cli-singletons.md)

### 4. 组装层 — 最后接 HTTP / agent

- `cmd/<app>/api_server.go`：HTTP server 挂载
- `cmd/<app>/apis/<domain>/`：operator 适配，embed endpoint + 注入 service，不重复声明契约
- `cmd/<app>/sidecar/`：后台 agent，只调 service，不直连 table/repository

→ [assembly-layer.md](references/assembly-layer.md)

### 5. 刷新生成

涉及 `+gengo:*` 注解修改时，执行 `go generate`。未刷新或无法刷新时必须在交付中说明原因。

### 6. 最小验证

按 repository → service → api-server 顺序验证。优先用 `testingutil` 构造上下文，用 `Must`/`Then` 断言。

→ [testing-and-fixtures.md](references/testing-and-fixtures.md)

## 任务路由

- 新增资源：按流水线 1→6 全流程推进
- 改已有接口：先判是否动契约；动则从步骤 1 开始，否则从步骤 2 或 4 切入
- 修 repository / service 缺陷：直接看步骤 2
- 修入口 / server 组装：直接看步骤 3 或 4
- 补测试：直接看步骤 6

## 交付要求

完成后说明：变更层级、生成刷新情况、验证结果、残余风险。
