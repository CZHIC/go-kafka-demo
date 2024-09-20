package mysql

import (
	"sync"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// 数据源
type dataSource struct {
	Read      gdb.DB
	Write     gdb.DB
	AuthRead  gdb.DB
	AuthWrite gdb.DB
}

var once sync.Once
var ds *dataSource

// 获取数据源
func GetDataSource() *dataSource {
	once.Do(func() {
		ds = &dataSource{
			Read:      g.DB("default"),
			Write:     g.DB("default"),
			AuthRead:  g.DB("default"),
			AuthWrite: g.DB("default"),
		}

	})
	return ds
}
