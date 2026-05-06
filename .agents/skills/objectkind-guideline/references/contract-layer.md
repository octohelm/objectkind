# 契约层

## 目录

- `pkg/apis/<domain>/v1`
- `pkg/endpoints/<domain>/v1`

## 职责

`pkg/apis` 负责稳定业务语义定义：

- 主对象
- request / action 变体
- reference
- list
- 业务错误
- 强语义 ID / Code / State / Reason 类型

`pkg/endpoints` 负责调用面定义：

- HTTP method
- path
- path / query / body 输入
- `ResponseData()`
- `ResponseErrors()`

## 准则

- 先定义对象契约，再定义操作契约
- `ResponseData()` 返回 `new(T)`
- 列表、筛选、分页等输入直接定义在 Operation 上
- endpoint 只描述契约，不写数据库和业务流程

## 反模式

- 在 `pkg/apis` 混入 repository / service 逻辑
- 在 `pkg/endpoints` 里写业务状态机
- 先写 operator，再回头补 endpoint
- 用裸 `string` 代替已有语义 ID / Code / State 类型
