# 数据操作层

## 目录

- `domain/<domain>/*.go` — 表模型定义
- `domain/<domain>/convert/` — 对象 ↔ 表模型转换
- `domain/<domain>/filter/` — 查询过滤条件
- `domain/<domain>/repository/` — 单 domain 数据访问
- `domain/<domain>/service/` — 业务编排

## 表结构

表模型放在 `domain/<domain>/*.go`，用 `+gengo:table` 声明表、索引和分组。具体注解约定见 `gengo-guideline` 与 `storage-guideline`。

storage session 通过 infra 的 `+gengo:injectable:provider` 机制注入：在 `cmd/{app}` 层声明 Database provider → service 通过 `inject:""` 获取 session → repository 从 service 获取数据库连接。

## repository

只做单 domain 数据访问：

- 插入 / 更新
- 删除
- 单条查询
- 列表查询
- 批量填充

## service

只做业务编排：

- 状态流转
- 跨 domain 查询与组合
- 聚合计算
- 供 operator / sidecar 复用

## 批量查询与填充

列表、聚合、关联补全场景优先使用批量能力：

- `filter.InSeq(...)`
- `sqlutil.FillSet(...)`
- `sqlutil/filler.Fill(...)`
- `FillOwnerSet(...)`

禁止在 service 循环里逐条 `FindOne`，禁止在 repository 里跨 domain 编排业务流程。

## 反模式

- 在 `service` 里逐条查询填充关联
- 在 repository 里跨 domain 编排业务流程
- 表结构里混入 API 契约层特有的类型或约束
