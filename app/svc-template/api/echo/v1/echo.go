package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SayReq struct {
	g.Meta  `path:"/echo/say" tags:"Echo Service" method:"get" summary:"You say, I echo"`
	Content string `v:"required" dc:"Say something?"`
}

type SayRes struct {
	Content string `dc:"Reply content"`
}
