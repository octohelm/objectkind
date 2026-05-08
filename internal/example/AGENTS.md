# internal/example AGENTS

## 文件约定

- 命令入口见上级 `justfile` 中的 `just example ...` 和 `just go ...`
- 项目分层说明见本目录 `README.md`
- 行为约束：根 `AGENTS.md` 的通用规则 + 本文档的 workspace 局部覆盖

## 分层

```
pkg/apis ──→ pkg/endpoints ──→ domain/* ──→ cmd/example
对象契约      操作契约           数据操作        HTTP 组装
```

- `pkg/apis`：objectkind、variant、reference、list、业务错误
- `pkg/endpoints`：path、method、Operation 入参、`ResponseData()`、`ResponseErrors()`
- `domain/{name}/`：表 model → `convert/` 转换 → `repository/` 单域数据访问 → `service/` 业务编排
- `cmd/example/`：基于 endpoint 声明 operator，通过 inject 接入 service

## 编码约束

- 主 objectkind 与 variant 分文件维护
- 有明确语义的 id、code、状态值声明对应类型，不使用裸 `string`
- `ResponseData()` 返回 `new(T)`；204 场景用 `new(courierhttp.NoContent)`
- API 层 operator 直接复用 endpoint 声明
- repository 测试必须先于 service
- service 依赖 repository 优先非 pointer 字段；API operator 注入 `*Service`
- Serve 和 Agent 通过 inject 接入依赖，只做组装不承载业务逻辑

## 生成顺序

```
go generate ./internal/example/pkg/apis/...
go generate ./internal/example/pkg/endpoints/...
go generate ./internal/example/domain/...
go generate ./internal/example/cmd/example/...
```

修改生成器或生成配置后执行对应范围，生成文件不手改。

## 什么时候停

- 无法明确逻辑应落在哪一层
- 需要新增规则但无法确认是 workspace 局部还是仓库级
- 变更无法完成最小生成与最小验证
- 需要新业务模型或跨 domain 协调但边界未定

## 规则生效

- 本文档覆盖根 `AGENTS.md` 的同主题规则，未覆盖部分继续继承上级
