package main

import (
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/template-mono/app/svc-template/internal/cmd"
)

func main() {
	var (
		ctx = gctx.New()
		command, err = gcmd.NewFromObject(cmd.Main)
	)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	if err = command.Run(ctx); err != nil {
		g.Log().Fatal(ctx, err)
	}
}
