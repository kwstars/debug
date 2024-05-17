run:
	nohup go run main.go > output.log 2>&1 &

tool:
	curl -o trace.out http://localhost:6060/debug/pprof/trace?seconds=5
	go tool trace trace.out

clean:
	-pkill -f "go run main.go"
	rm -f output.log
	rm -f trace.out