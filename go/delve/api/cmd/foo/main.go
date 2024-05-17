package main

import (
	"fmt"
	"os"

	"demo/pkg/baz"
)

func main() {
	arg := os.Args[1]
	result := baz.MyFunction(arg)
	fmt.Println(result)
}
