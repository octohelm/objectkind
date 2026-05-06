# CLI 单例入口

## 目录

- `cmd/<app>/main.go`
- `cmd/<app>/serve.go`
- `cmd/<app>/database.go`

## 入口职责

### main.go

- 创建 CLI app
- 调用 `cli.Exec(...)`
- 具体 API 见 `infra-guideline`

### serve.go

- 注册顶层运行组件
- 聚合 Database、Service、Server 等单例
- 通过字段声明单例依赖

### database.go

- 定义数据库 provider
- 设置默认库名、catalog 和迁移配置
- 作为下游 repository / service 的单例来源
- 具体用法见 `storage-guideline`

## 准则

- CLI 入口负责装配，不承载业务逻辑
- provider 通过 `+gengo:injectable` 生命周期接入，不手拼对象图
- 基础设施单例在 `cmd/<app>` 层声明，不散落到 domain

## 反模式

- 在 `main.go` 里直接 new 业务对象
- 在 `serve.go` 写业务分支
- 把数据库初始化细节散到 repository / service
