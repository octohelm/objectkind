# 总览与目录结构

## 目录结构

```
<module>/
├── pkg/
│   ├── apis/<domain>/v1/        # 对象、request、reference、list、错误
│   └── endpoints/<domain>/v1/   # method、path、输入输出
├── domain/<domain>/
│   ├── *.go                     # 表模型
│   ├── convert/                 # 对象 ↔ 表模型转换
│   ├── filter/                  # 查询过滤条件
│   ├── repository/              # 单 domain 数据访问与批量填充
│   └── service/                 # 业务编排与跨 domain 协作
└── cmd/<app>/
    ├── main.go                  # CLI 应用入口
    ├── serve.go                 # 运行时组件装配
    ├── database.go              # 数据库单例 provider
    ├── api_server.go            # HTTP server 单例与路由挂载
    ├── apis/<domain>/           # operator 适配层
    └── sidecar/                 # 后台 agent
```

## 层级职责

- `pkg/apis`：稳定业务对象与业务错误
- `pkg/endpoints`：method、path、输入输出、错误集合
- `domain/repository`：单 domain 数据访问、批量查询、填充器
- `domain/service`：业务编排、状态流转、跨 domain 组合
- `cmd/<app>`：运行入口、依赖注入、HTTP server 与 agent 组装

## 依赖方向

1. `pkg/apis`
2. `pkg/endpoints`
3. `domain`
4. `cmd/<app>/apis`
5. `cmd/<app>/sidecar`
6. `cmd/<app>/serve.go` / `api_server.go` / `main.go`

反向依赖禁止：

- endpoint 反向依赖 operator
- repository 反向依赖 service
- service 感知 HTTP 细节
- sidecar 直连 table / repository
