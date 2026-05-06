# internal/example AGENTS

## 适用范围

- 作用域：`internal/example`（目录及其子目录）
- 目标：约束 example workspace 下从契约定义、数据操作层到 HTTP API 组装层的完整构建流程

## 可改范围与禁区

- 可改：`pkg/apis` 下的资源、变体、业务错误与生成文件
- 可改：`pkg/endpoints` 下的 Operation 声明、Operation 内输入参数与生成文件
- 可改：`domain/*` 下的 table model、convert、repository、service 及其测试
- 可改：`cmd/example` 下基于契约与 service 的 HTTP API 组装代码与生成文件
- 可改：为保持当前 workspace 一致性所做的最小 import 调整、生成刷新与文档更新
- 禁区：未授权修改当前作用域之外的内容
- 禁区：脱离目标的顺手重构、批量命名清理、目录重排
- 禁区：在 `service` 中下沉 transport 协议、HTTP status、auth、rbac 等 HTTP 封装细节
- 禁区：在 `repository` 中跨 domain 访问其他 domain 的表或 repository

## Workspace 分层约束

- `pkg/apis` 是对象契约层：定义 objectkind、variant、reference、list、业务错误
- `pkg/endpoints` 是操作契约层：定义 path、method、Operation 入参、`ResponseData()`、`ResponseErrors()`
- `domain/*` 是数据操作层：
  - `{table}.go` 定义数据库表 model
  - `convert/` 负责表 model 与契约对象转换
  - `repository/` 只做单 domain 数据访问
  - `service/` 做业务编排、跨 domain 组合、后续 cache 或 agent loop 支撑
- `cmd/example/apis/*` 是 HTTP API 适配层：基于 endpoint 声明 operator，通过 inject 接入 `service`
- `cmd/example` 是 workspace 的运行入口层：
  - `Serve` 负责对外 HTTP API 与 server 生命周期
  - `sidecar/*` 负责后台 agent、cron job、治理型异步任务
  - 二者都只做组装，不承载业务逻辑

## 实现约束

- 主 objectkind 与 variant 分文件维护，不把对象本体、引用、请求、列表堆在一个文件中
- 业务上有明确语义的 `id`、`code`、状态值、原因值等，需要声明对应类型，不直接使用裸 `string`
- 列表、筛选、分页等输入直接定义到对应 Operation 上，避免抽出独立 `*Query` 形成额外跳转
- `ResponseData()` 必须返回 `new(T)`；204 场景使用 `new(courierhttp.NoContent)`
- API 层 operator 直接复用 endpoint 声明，不重复手写 path、method、response 契约
- 可直接复用的查询能力优先通过 repository 或 embed 复用，不在 service 中机械重复代理
- `service` 依赖 `repository` 时优先使用非 pointer 字段；API operator 注入 `*Service`
- `Serve` 和 `Agent` 都通过 inject / singleton 生命周期接入 `Database`、`Service` 等依赖，不手工拼接业务对象
- `Agent` 只消费 `service` 能力，不直接操作 table model 或跨层下沉到 repository 明细

## 生成约束

- 涉及 `objectkind`、variant、`uintstr`、`runtimedoc`、`table`、`filterop`、`injectable`、`operator`、`client` 等生成产物的修改，必须执行对应范围的 `go generate`
- 修改生成依赖前，先确认 `internal/cmd/devtool/main.go` 已注册所需生成器
- 默认生成顺序：
  1. `go generate ./internal/example/pkg/apis/...`
  2. `go generate ./internal/example/pkg/endpoints/...`
  3. `go generate ./internal/example/domain/...`
  4. `go generate ./internal/example/cmd/example/...`

## 验证顺序

- `repository` 测试必须先于 `service`
- 推荐验证顺序：
  1. `go test ./internal/example/domain/product/repository ./internal/example/domain/order/repository`
  2. `go test ./internal/example/domain/product/service ./internal/example/domain/order/service`
  3. `go test ./internal/example/cmd/example/...`
  4. `go test ./internal/example/cmd/example/sidecar/...`
  5. 必要时再补运行时或回归验证

## 停止条件

- 无法明确某段逻辑应落在契约层、repository、service 还是 API 适配层
- 需要新增规则，但无法确认是当前 workspace 局部规则还是应上提到仓库级规则
- 变更无法完成最小生成与最小验证
- 存在多种互斥实现路线，且会改变分层边界

## 人工接管条件

- 需要新的业务模型、事务边界或跨 domain 协调规则决策
- 需要引入缓存、异步任务、内部 agent loop 等执行模型，但边界尚未确定
- 需要目录外权限、依赖安装、外部系统接入或发布流程支持

## 规则优先级

1. 仓库根 `AGENTS.md`
2. 当前 workspace `AGENTS.md`
3. 子目录 `AGENTS.md`
4. 当前任务临时约束

## 新规则审校机制

1. 新规则必须说明目标、边界、风险、回滚方式
2. 若与仓库根规则或当前 workspace 既有分层职责冲突，需先完成冲突消解再启用
3. 涉及生成或测试流程变更时，需同步更新 README 中的执行步骤与验证顺序
4. 新规则落地时只做最小增量修改，不一次性重写整套文档或目录结构
