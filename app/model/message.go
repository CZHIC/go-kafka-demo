package model

import (
	"iv-test/library/logger"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var (
	TableMessage = "t_message"
)

type TMessage struct {
	Id         int    `orm:"id,primary" json:"id"`         // id
	UserId     int    `orm:"userId" json:"userId"`         // 用户ID
	Phone      int    `orm:"phone" json:"phone"`           // 手机号
	ActivityId int    `orm:"activityId" json:"activityId"` // 活动ID
	Message    string `orm:"message" json:"message"`       // 公司名称
	SendTime   int    `orm:"sendTime" json:"sendTime"`     // 发送时间
	IfSend     int    `orm:"ifSend" json:"ifSend"`         // 是否发送
	PlainTime  int    `orm:"plainTime" json:"plainTime"`   // 计划时间
}

func GetMessageAllList() ([]*TMessage, error) {
	messageList := ([]*TMessage)(nil)
	db := g.DB()
	nowtime := gtime.Now().Unix()
	sql := "select * from " + TableMessage + "where ifSend = 0  and sendTime < ? "
	record, err := db.GetAll(sql, nowtime)
	if err != nil {
		logger.Error(err)
	}
	// 转换
	if record != nil {
		if err := gconv.Structs(record, &messageList); err != nil {
			logger.Errorf("GetInvestMentHistoryList -> "+
				"gconv.Struct() execute failed. err = %v", err)
			return nil, err
		}
	}
	return messageList, nil
}

func AddMessage(data *TMessage) error {
	db := g.DB()
	_, err := db.Save(TableMessage, data)
	if err != nil {
		return err
	}
	return nil
}

// 更新发送状态
func UpdateMessageIfSend(id int) error {
	db := g.DB()
	_, err := db.Update(TableMessage, " `ifSend` = 1", " id = ? ", id)
	if err != nil {
		logger.Error(err.Error())
	}
	return nil
}
