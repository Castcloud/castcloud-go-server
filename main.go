package main

import (
	"runtime"

	_ "github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/mattn/go-isatty"

	"github.com/Castcloud/castcloud-go-server/cli"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cli.Execute()
}
