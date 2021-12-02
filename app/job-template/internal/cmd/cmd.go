package cmd

import (
	"context"
	"os"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gproc"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start crontab job",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Info(ctx, `cron job start`)
			_, err = gcron.Add(ctx, "* * * * * *", func(ctx context.Context) {
				g.Log().Debug(ctx, `cron job running`)
			})
			if err != nil {
				return err
			}
			// Register shutdown handler.
			gproc.AddSigHandlerShutdown(func(sig os.Signal) {
				g.Log().Info(ctx, `cron job shutdown`)
			})
			// Block listening the shutdown signal.
			g.Listen()
			return
		},
	}
)
