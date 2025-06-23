# 参考文档: https://github.com/golang/vscode-go/wiki/debugging#remote-debugging

build:
	@echo "构建应用程序..."
	mkdir -p bin
	go build -o bin/foo cmd/foo/main.go

# 构建带调试信息的可执行文件（用于调试）
build-debug:
	@echo "构建调试版本..."
	mkdir -p bin
	go build -gcflags="all=-N -l" -o bin/foo-debug cmd/foo/main.go

# 运行应用程序（需要提供参数）
run:
	@echo "运行应用程序（示例参数：test-arg）..."
	go run cmd/foo/main.go test-arg

# 调试相关目标
# =============

# 本地调试：运行无头调试服务器 - 程序在启动时暂停
# 此模式下程序不会自动开始执行，需要调试器连接后手动启动
run-headless-app-paused:
	nohup dlv debug cmd/foo/main.go --headless --listen=0.0.0.0:12345 --api-version 2 \
	--log --log-output=debugger,gdbwire,lldbout,debuglineerr,rpc,dap,fncall,minidump,stack --log-dest=./dlv.log -- arg1 > myapp.log 2>&1 &

# 本地调试：运行无头调试服务器 - 程序自动开始执行
# 此模式下程序会立即开始运行，支持多客户端连接
run-headless-app-running:
	nohup dlv debug cmd/foo/main.go --headless --listen=0.0.0.0:12345 --api-version 2 --accept-multiclient --continue \
	--log --log-output=debugger,gdbwire,lldbout,debuglineerr,rpc,dap,fncall,minidump,stack --log-dest=./dlv.log -- arg1 > myapp.log 2>&1 &

# 远程调试：使用已构建的二进制文件进行调试 - 程序自动开始执行
# 适用于部署到远程服务器的场景，使用预编译的二进制文件
remote-debug-running:
	@echo "启动远程调试服务器（自动运行模式）..."
	@echo "二进制文件: bin/foo-debug"
	@echo "监听地址: 0.0.0.0:12345"
	@echo "程序参数: arg1"
	nohup dlv exec bin/foo-debug --headless --listen=0.0.0.0:12345 --api-version 2 --accept-multiclient --continue \
	--log --log-output=debugger,gdbwire,lldbout,debuglineerr,rpc,dap,fncall,minidump,stack --log-dest=./dlv.log -- arg1 > myapp.log 2>&1 &

# 远程调试：使用已构建的二进制文件进行调试 - 程序启动时暂停
# 适用于需要在程序开始前设置断点的场景
remote-debug-paused:
	@echo "启动远程调试服务器（暂停模式）..."
	@echo "二进制文件: bin/foo-debug"
	@echo "监听地址: 0.0.0.0:12345"
	@echo "程序参数: arg1"
	nohup dlv exec bin/foo-debug --headless --listen=0.0.0.0:12345 --api-version 2 --accept-multiclient \
	--log --log-output=debugger,gdbwire,lldbout,debuglineerr,rpc,dap,fncall,minidump,stack --log-dest=./dlv.log -- arg1 > myapp.log 2>&1 &

# 远程调试：attach 到正在运行的进程
# 适用于调试已经在运行的程序
remote-debug-attach:
	@echo "使用以下命令 attach 到运行中的进程："
	@echo "1. 首先找到进程 PID: ps aux | grep foo"
	@echo "2. 然后运行: dlv attach <PID> --headless --listen=0.0.0.0:12345 --api-version 2 --accept-multiclient"
	@echo "示例: dlv attach 1234 --headless --listen=0.0.0.0:12345 --api-version 2 --accept-multiclient"

# 清理进程和日志文件
# 停止所有相关的 dlv 调试进程并删除日志文件
clean:
	-pkill -f "dlv"                        # 终止所有 dlv 进程 (- 前缀表示忽略错误)
	rm -f myapp.log dlv.log                # 删除日志文件
	rm -rf bin/                            # 删除构建产物

# 显示帮助信息
help:
	@echo "可用的构建目标："
	@echo "  build                    - 构建可执行文件到 bin/foo"
	@echo "  build-debug              - 构建调试版本到 bin/foo-debug"
	@echo "  run                      - 直接运行应用程序"
	@echo "  run-headless-app-paused  - 启动暂停状态的调试服务器"
	@echo "  run-headless-app-running - 启动运行状态的调试服务器"
	@echo "  remote-debug-running      - 启动远程调试服务器（自动运行模式）"
	@echo "  remote-debug-paused       - 启动远程调试服务器（暂停模式）"
	@echo "  remote-debug-attach       - attach 到正在运行的进程"
	@echo "  clean                    - 清理进程和文件"
	@echo "  help                     - 显示此帮助信息"

# 默认目标
.DEFAULT_GOAL := help

# 声明伪目标
.PHONY: build build-debug run test test-verbose fmt vet deps run-headless-app-paused run-headless-app-running remote-debug-running remote-debug-paused remote-debug-attach clean help