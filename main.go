package main

import (
	"iv-test/app/api"
	_ "iv-test/boot"
	_ "iv-test/router"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
)

func main() {

	// 定时任务 ，每分钟执行，查询消息表，发送到kafka
	gcron.Add(g.Cfg().GetString("GCRON.GetMessage"), api.GetMessage)

	g.Server().Run()
}
