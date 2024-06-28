package main

import (
	_ "wy-goframe-admin/internal/packed"

	_ "wy-goframe-admin/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"wy-goframe-admin/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
