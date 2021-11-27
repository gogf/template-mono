package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/katyusha/krpc"
	"github.com/gogf/template-mono/app/svc-template/internal/consts"
	"github.com/gogf/template-mono/app/svc-template/internal/handler"
	"github.com/gogf/template-mono/app/svc-template/protobuf"
)

var (
	Main = gcmd.Command{
		Name:          "main",
		Usage:         "main",
		Brief:         "start grpc server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 配置对象配置
			c := krpc.Server.NewGrpcServerConfig()
			// 指定服务名称
			c.AppId = consts.AppId

			// Server对象配置
			s := krpc.Server.NewGrpcServer(c)
			// 服务对象注册
			protobuf.RegisterEchoServer(s.Server, handler.Echo)
			s.Run()
			return nil
		},
	}
)
