package echo

import (
	"context"

	"github.com/gogf/template-mono/app/svc-template/api/echo/v1"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) Say(ctx context.Context, req *v1.SayReq) (res *v1.SayRes, err error) {
	return &v1.SayRes{Content: req.Content}, nil
}
