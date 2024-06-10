package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloNameReq struct {
	// 动态路由
	g.Meta `path:"/hello/:name" tags:"Hello" method:"get" summary:"You first hello api"`
	// name 名字 首字母大写为公共字段，小写为私有字段
	Name string `dc:"名字"`
}

type HelloNameRes struct {
	g.Meta `mime:"text/html" example:"string"` // example 为值示例
}
