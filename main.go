package main

import (
	"runtime"

	"github.com/khlieng/castcloud-go/cli"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cli.Execute()
}
