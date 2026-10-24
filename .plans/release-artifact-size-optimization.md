# Beacon 镜像与 Collia 二进制体积优化实施计划

## Component Map

- 本任务不修改 Vue 组件、组件状态或数据流。
- `beacon/web/vite.config.ts` 只负责生产输出目录生命周期，确保每次构建只保留当前资源。
- `beacon/Dockerfile` 负责前端隔离构建、Go 发布构建、双架构 Agent 打包和精简运行时组装。
- `beacon/Taskfile.yml` 与 `collia/Taskfile.yml` 负责向发布构建传递版本并统一裁剪参数。
- 发布/模块文档负责记录新的产物边界与验证命令。

## Steps

1. 记录 v3.0.6 镜像、未裁剪 Collia、累积 Web dist 和干净 Web dist 的体积基线。
2. 更新 Release Install Domain Spec，约束发布二进制裁剪、版本保留与前端产物清理。
3. 启用 Vite 输出目录清理，并在 Docker 中增加隔离的 Beacon Web 构建阶段。
4. 为 Beacon 与两个架构的 Collia 启用静态、路径裁剪、符号裁剪构建，并从 Taskfile 注入版本。
5. 将运行时切换到精简 Linux 基础镜像，保持 Nginx、Supervisor、CA、健康检查命令和现有路径/端口不变。
6. 统一本地 Collia Taskfile 的发布参数，防止镜像内外产物规格漂移。
7. 更新部署与 Beacon/Collia 模块文档，记录构建产物、体积预算和验证方式。
8. 构建 amd64/arm64 Collia 与本地架构 Beacon 镜像，检查架构、版本、静态链接、文件大小、镜像大小和容器健康。
9. 运行 Go 测试、前端类型/测试/生产构建、文档门禁及 `git diff --check`。

## Risk Controls

- 不使用 UPX，保留标准 ELF 行为和常规故障分析能力。
- 继续打包两个 Agent 架构，不以缩小镜像为由破坏跨架构安装。
- 保持 Nginx/Supervisor 编排、容器 80 端口及 `/app/downloads/collia/<arch>/collia` 路径稳定。
- 运行时切换后必须以真实容器验证 Supervisor、Nginx、Beacon、健康检查和安装包文件。
- 体积比较使用相同平台、相同 Docker 引擎的 `docker image inspect .Size`，避免把架构差异误判为优化收益。

## Verification

- `cd beacon/web && pnpm run test:run && pnpm run ts && pnpm run build`
- `cd beacon && go test ./...`
- `cd collia && go test ./...`
- `task collia:amd64` 与 `task collia:arm64`
- `task beacon:amd64` 或等价本地平台镜像构建
- 检查镜像内二进制大小、ELF 架构、动态依赖和 `--version`
- 启动容器并检查 Supervisor 进程、Nginx、Beacon 与 `/health`
- 文档治理检查与 `git diff --check`
