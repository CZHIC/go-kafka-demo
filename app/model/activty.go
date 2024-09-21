package model

import (
	"iv-test/library/logger"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var (
	TableActivty = "t_activity"
)

type TActivty struct {
	Id        int    `orm:"id,primary" json:"id"`       // id
	Template  string `orm:"template" json:"template"`   // 模版
	PlainTime int    `orm:"plainTime" json:"plainTime"` // 计划时间
}

func GetActivtyInfo(Id int64) (*TActivty, error) {
	activtyInfo := (*TActivty)(nil)
	db := g.DB()
	sql := "select * from " + TableActivty + " where id = ? "
	record, err := db.GetOne(sql, Id)
	if err != nil {
		logger.Error(err)
	}
	// 转换
	if record != nil {
		if err := gconv.Struct(record, &activtyInfo); err != nil {
			logger.Errorf("GetActivtyInfo -> "+
				"gconv.Struct() execute failed. err = %v", err)
			return nil, err
		}
	}
	return activtyInfo, nil
}
