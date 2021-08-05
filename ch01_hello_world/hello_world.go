package main

import (
	"fmt"
	"os"
)

func main() {
	// 应用的参数
	fmt.Println(os.Args[1:])

	fmt.Println("hello world")

	// 返回状态
	os.Exit(0)
}
