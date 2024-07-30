package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "shop-v2/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"

	_ "shop-v2/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"shop-v2/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
