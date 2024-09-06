// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package echo

import (
	"context"

	"${MODULE_NAME}/app/svc-template/api/echo/v1"
)

type IEchoV1 interface {
	Say(ctx context.Context, req *v1.SayReq) (res *v1.SayRes, err error)
}
