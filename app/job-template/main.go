package main

import (
	_ "${MODULE_NAME}/app/job-template/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"${MODULE_NAME}/app/job-template/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
