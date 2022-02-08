package controller

import (
	"context"

	"github.com/gogf/template-mono/app/svc-template/apiv1"
)

var (
	Echo = cEcho{}
)

type cEcho struct{}

func (c *cEcho) Say(ctx context.Context, req *apiv1.EchoSayReq) (res *apiv1.EchoSayRes, err error) {
	return &apiv1.EchoSayRes{Content: req.Content}, nil
}
