run:
	go run cmd/foo/main.go arg1

dlv-headless:
	dlv debug --headless --listen 127.0.0.1:2345 cmd/foo/main.go -- arg1

dlv-connect:
	dlv connect localhost:2345