package main

import (
	_ "goframe_tutorial/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"goframe_tutorial/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
