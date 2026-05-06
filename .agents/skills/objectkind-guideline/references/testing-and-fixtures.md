# 测试与实例构造

## 测试框架

使用 `github.com/octohelm/x/testing/v2`。具体断言 API 见 `testing-guideline`。

## 上下文构造

- repository 测试：组装 otel、idgen、Database 与 repository，用 `testingutil.NewContext(t, d)`
- service 测试：用 `testingutil.BuildContext(t, ...)` 一次装起 Database、Repository、Service
- 具体 API 见 `infra-guideline`

## 实例构造

优先用 `runtime.Build(...)` 构造对象或请求体，避免在测试里手写大量样板赋值。

## 验证顺序

1. repository
2. service
3. `cmd/<app>/...`
4. `cmd/<app>/sidecar/...`

## 反模式

- service 还没稳定就先写 API 测试
- 断言描述不清楚，读不出业务语义
- 测试里绕开 service 直接验证 HTTP 适配层主逻辑
