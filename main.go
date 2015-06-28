package main

import (
	"runtime"

	"github.com/Castcloud/castcloud-go-server/cli"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cli.Execute()
}
