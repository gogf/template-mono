package handler

import (
	"context"
	"fmt"

	"github.com/gogf/template-mono/app/svc-template/protobuf"
)

var (
	Echo = &handlerEcho{}
)

type handlerEcho struct{}

func (a *handlerEcho) Say(ctx context.Context, r *protobuf.SayReq) (*protobuf.SayRes, error) {
	text := fmt.Sprintf(`> %s`, r.Content)
	return &protobuf.SayRes{Content: text}, nil
}
