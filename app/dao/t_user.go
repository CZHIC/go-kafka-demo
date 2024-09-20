package dao

import (
	"iv-test/library/logger"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

var t_users = "t_user"

// 获取用户名
func GetStaffNames() gdb.Result {
	db := g.DB("")
	sql := "select id,staff_name from " + t_users
	ret, err := db.GetAll(sql)
	if err != nil {
		logger.Error("查询错误：", err)
	}
	return ret
}

func UpdateStaffName(staffName string, id int) {
	db := g.DB("")
	sql := "update " + t_users + " set staff_name =  ?  where id = ?"
	_, err := db.Exec(sql, g.Slice{staffName, id})
	if err != nil {
		logger.Error("插入错误：", err)
	}
}

// 获取用户名
func GetStaffNameById(Id int64) gdb.Record {
	db := g.DB("")
	sql := "select id,staff_name from " + t_users + " where id = ?"
	ret, err := db.GetOne(sql, g.Slice{Id})
	if err != nil {
		logger.Error("查询错误：", err)
	}
	return ret
}
