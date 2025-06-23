# 参考文档: https://github.com/golang/vscode-go/wiki/debugging#remote-debugging
#
# 重要提示：需要使用 Delve v1.7.3 或更高版本通过 DAP (Debug Adapter Protocol) 连接到外部 dlv --headless 服务器
# 较旧版本会在终端中记录错误："error layer=rpc rpc:invalid character 'C' looking for beginning of value"

# 运行无头调试服务器 - 程序在启动时暂停
# 此模式下程序不会自动开始执行，需要调试器连接后手动启动
run-headless-app-paused:
	nohup dlv debug cmd/foo/main.go --headless --listen=0.0.0.0:12345 --api-version 2 \
	--log --log-output=debugger,gdbwire,lldbout,debuglineerr,rpc,dap,fncall,minidump,stack --log-dest=./dlv.log -- arg1 > myapp.log 2>&1 &

# 运行无头调试服务器 - 程序自动开始执行
# 此模式下程序会立即开始运行，支持多客户端连接
run-headless-app-running:
	nohup dlv debug cmd/foo/main.go --headless --listen=0.0.0.0:12345 --api-version 2 --accept-multiclient --continue \
	--log --log-output=debugger,gdbwire,lldbout,debuglineerr,rpc,dap,fncall,minidump,stack --log-dest=./dlv.log -- arg1 > myapp.log 2>&1 &

# 清理进程和日志文件
# 停止所有相关的 dlv 调试进程并删除日志文件
clean:
	-pkill -f "dlv debug cmd/foo/main.go"  # 终止匹配的 dlv 进程 (- 前缀表示忽略错误)
	rm -f myapp.log dlv.log                # 删除日志文件