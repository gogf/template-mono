package main

import (
	_ "github.com/gogf/template-mono/app/job-template/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/template-mono/app/job-template/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
