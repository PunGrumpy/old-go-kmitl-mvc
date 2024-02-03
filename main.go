package main

import (
	_ "kmitl-mvc/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"kmitl-mvc/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
