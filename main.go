package main

import (
	_ "golang8090/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"golang8090/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
