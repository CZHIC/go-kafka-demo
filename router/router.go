package router

import (
	"iv-test/app/api"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {

	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/test", api.Test)

	})
}
