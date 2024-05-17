run:
	go run foo/cmd/foo/main.go arg1


dlv-debug:
	echo "break cmd/foo/main.go:12" > init.txt
	dlv debug --init init.txt cmd/foo/main.go -- arg1

dlv-test:
	echo "break TestMyFunction" > init.txt
	dlv test --init init.txt ./pkg/baz

clean:
	rm -f init.txt
