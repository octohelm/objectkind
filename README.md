# objectkind

提供对象 `Kind/APIVersion` 建模、运行时构建与配套代码生成能力。

仓库聚焦通用对象类型、运行时和开发工具，不承载独立业务服务；`internal/example` 仅作为最小端到端示例，用来验证契约、运行时与生成链路如何协作。

## 主要入口

- `pkg/object`、`pkg/apis/meta/v1`：对象元接口与元数据定义，是大部分上层对象建模的起点。
- `pkg/runtime`、`pkg/runtime/converter`：对象构建和对象转换契约。
- `devpkg/objectkindgen`：`objectkind` 生成器实现。
- `tool/internal/cmd/devtool`：仓库开发工具 CLI 入口，承载 `go tool devtool` 使用的格式化与生成能力。
- `internal/example`：示例 project，展示契约层、domain 层和运行入口如何组合。

## 怎么开始

1. 安装 Go 1.26+，并保证 `go.mod` 依赖可解析。
2. 运行 `just` 查看仓库暴露的稳定入口。
3. 需要刷新生成产物时运行 `just go::gen ./...`。
4. 需要做最小回归时运行 `just go::test ./...`。

## 更多信息

- [AGENTS.md](AGENTS.md)：仓库级协作约束与验证义务。
- [justfile](justfile)：root 执行入口，聚合 Go toolchain、skill 和 example 子模块。
- [internal/example/README.md](internal/example/README.md)：示例 workspace 的分层说明与阅读路径。
