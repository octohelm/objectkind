# Go 工具链入口
[group: 'toolchain']
mod go "tool/go/justfile"

# 示例 CLI 入口
[group: 'app']
mod example "internal/example/cmd/example/justfile"

# 列出所有可用命令（含子模块）
[group('meta')]
default:
    @just --list --list-submodules
