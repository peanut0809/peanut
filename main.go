package main

import (
	_ "golang8090/internal/packed"

	_ "golang8090/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"golang8090/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
