package handler

import (
	"context"

	"github.com/gogf/gf/net/ghttp"
)

var (
	Hello = handlerHello{}
)

type handlerHello struct {}

// Index is a demonstration route handler for output "Hello World!".
func (*handlerHello) Index(ctx context.Context, r *ghttp.Request) {
	r.Response.Writeln("Hello World!")
}
