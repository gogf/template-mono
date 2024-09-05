package echo

import (
	"context"

	"${MODULE_NAME}/app/svc-template/api/echo/v1"
)

func (c *ControllerV1) Say(ctx context.Context, req *v1.SayReq) (res *v1.SayRes, err error) {
	return &v1.SayRes{Content: req.Content}, nil
}
