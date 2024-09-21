package main

import (
	"iv-test/app/api"
	_ "iv-test/boot"
	_ "iv-test/router"
	"time"

	"github.com/go-echarts/statsview"
	"github.com/go-echarts/statsview/viewer"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
)

func main() {

	go func() {
		//增加一个系统监控
		mgr := statsview.New()
		viewer.SetConfiguration(viewer.WithTheme(viewer.ThemeMacarons), viewer.WithMaxPoints(1000))
		mgr.Start()
		//访问。 http://localhost:18066/debug/statsview
		// busy working....
		time.Sleep(time.Second * 10)
	}()

	// 定时任务 ，每分钟执行，查询消息表，发送到kafka
	gcron.Add(g.Cfg().GetString("GCRON.GetMessage"), api.GetMessage)

	g.Server().Run()
}
