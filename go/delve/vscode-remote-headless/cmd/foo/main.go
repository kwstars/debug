package main

import (
	"fmt"
	"os"
	"time"

	"demo/pkg/baz"
)

func main() {
	arg := os.Args[1]
	result := baz.MyFunction(arg)
	fmt.Println(result)

	for {
		time.Sleep(1 * time.Second)
		fmt.Println(result)
	}
}
