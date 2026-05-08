---
name: objectkind-guideline
description: 开发或修改 Go 后端服务时使用；覆盖从契约定义、数据操作、CLI 单例、api-server/sidecar 组装到测试验证的完整流程。
---

# ObjectKind Guideline

从零开发一个完整的 CRUD 服务。框架全景：

```
objectkind (对象建模)
 ├─ courier (HTTP Router) → infra (CLI / Server)
 ├─ storage (SQL / session) → converter (Object ↔ Model)
 └─ gengo (代码生成) → enumeration
```

## 核心流程

按 6 步流水线推进：

### 1. 契约层 — 定义对象和接口

```go
// pkg/apis/{domain}/v1/org.go
// +gengo:objectkind
type Org struct {
    ID   OrgID  `json:"id"`
    Name string `json:"name"`
}
```

```go
// pkg/endpoints/{domain}/v1/orgs.go
type CreateOrg struct {
    courierhttp.MethodPost `path:"/orgs"`
    Body CreateOrgRequest  `in:"body"`
}
func (CreateOrg) ResponseData() *Org { return new(Org) }
```

→ 细节见 [contract-layer.md](references/contract-layer.md)

### 2. 数据层 — 表模型和操作

```go
// domain/{domain}/table.go — storage sqlbuilder 表定义
// domain/{domain}/convert/ — objectkind converter 双向转换
// domain/{domain}/repository/ — 单 domain 数据访问
// domain/{domain}/service/ — 业务编排、跨 domain 组合
```

→ 细节见 [data-layer.md](references/data-layer.md)

### 3. CLI 入口 — 依赖注入

`cmd/{app}/` 声明 provider，通过 `+gengo:injectable` 生命周期接入。

→ 细节见 [cli-singletons.md](references/cli-singletons.md)

### 4. 组装层 — HTTP / agent

operator embed endpoint + 注入 service，`infra/pkg/http` 挂载 courier Router。

→ 细节见 [assembly-layer.md](references/assembly-layer.md)

### 5. 刷新生成

涉及 `+gengo:*` 修改后执行 `go generate`。

### 6. 验证

按 repository → service → api-server 顺序验证。

→ 细节见 [testing-and-fixtures.md](references/testing-and-fixtures.md)

## 所依赖的 skill

具体工具约定由对应 skill 负责：`courier-guideline` / `storage-guideline` / `infra-guideline` / `gengo-guideline` / `enumeration-guideline` / `testing-guideline`。
