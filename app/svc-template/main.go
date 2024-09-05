package main

import (
	_ "${MODULE_NAME}/app/svc-template/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"${MODULE_NAME}/app/svc-template/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
