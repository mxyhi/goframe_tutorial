// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package hello

import (
	"context"

	"goframe_tutorial/api/hello/v1"
)

type IHelloV1 interface {
	Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error)
	HelloName(ctx context.Context, req *v1.HelloNameReq) (res *v1.HelloNameRes, err error)
}
