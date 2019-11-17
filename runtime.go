package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOROOT--->", runtime.GOROOT())
	fmt.Println("os/platform--->", runtime.GOOS)
	fmt.Println("CPU数量--->", runtime.NumCPU())
}
