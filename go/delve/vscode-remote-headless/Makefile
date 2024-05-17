# https://github.com/golang/vscode-go/wiki/debugging#remote-debugging

run-headless-app-paused:
	nohup dlv debug cmd/foo/main.go --headless --listen=0.0.0.0:12345 --api-version 2 \
	--log --log-output=debugger,gdbwire,lldbout,debuglineerr,rpc,dap,fncall,minidump,stack --log-dest=./dlv.log -- arg1 > myapp.log 2>&1 &

run-headless-app-running:
	nohup dlv debug cmd/foo/main.go --headless --listen=0.0.0.0:12345 --api-version 2 --accept-multiclient --continue \
	--log --log-output=debugger,gdbwire,lldbout,debuglineerr,rpc,dap,fncall,minidump,stack --log-dest=./dlv.log -- arg1 > myapp.log 2>&1 &

clean:
	-pkill -f "dlv debug cmd/foo/main.go"
	rm -f myapp.log dlv.log