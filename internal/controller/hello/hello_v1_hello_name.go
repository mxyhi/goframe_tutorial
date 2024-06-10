package hello

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "goframe_tutorial/api/hello/v1"
)

func (c *ControllerV1) HelloName(ctx context.Context, req *v1.HelloNameReq) (res *v1.HelloNameRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
