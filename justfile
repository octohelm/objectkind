# Root 命令入口：注册 toolchain 与模块 justfile

# Go toolchain
mod go "tool/go/justfile"
# example cli
mod example "internal/example/cmd/example/justfile"

# 列出所有可用命令（含子模块，无输入）
[group('meta')]
default:
    @just --list --list-submodules
