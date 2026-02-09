package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/template-mono/app/svc-template/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
