package main

import (
	_ "wy-goframe-admin/internal/packed"

	_ "wy-goframe-admin/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"wy-goframe-admin/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
