package main

import (
	"fmt"
	"runtime"

	lunatic "github.com/alecthomas/lunatic-go"
)

func main() {
	fmt.Println(runtime.Compiler, runtime.Version(), runtime.GOARCH, runtime.GOOS)
	major, minor, patch := lunatic.Version()
	fmt.Printf("Lunatic %d.%d.%d\n", major, minor, patch)
}
