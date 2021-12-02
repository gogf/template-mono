package handler

import (
	"context"

	"github.com/gogf/template-mono/app/svc-template/apiv1"
)

var (
	Echo = &handlerEcho{}
)

type handlerEcho struct{}

func (a *handlerEcho) Say(ctx context.Context, req *apiv1.EchoSayReq) (res *apiv1.EchoSayRes, err error) {
	return &apiv1.EchoSayRes{Content: req.Content}, nil
}
