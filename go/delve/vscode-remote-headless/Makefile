# https://github.com/golang/vscode-go/wiki/debugging#remote-debugging

run-headless-app:
	nohup dlv debug cmd/foo/main.go --headless --listen=0.0.0.0:12345 --accept-multiclient --continue -- arg1 > myapp.log 2>&1 &

clean:
	-pkill -f "dlv debug cmd/foo/main.go"
	rm -f myapp.log