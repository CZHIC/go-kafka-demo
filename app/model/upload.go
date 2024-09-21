package model

// 将word文档类容录入数据库
type UploadDoc struct {
	UserId    int64  `json:"userId"`    // 用户ID
	Phone     int64  `json:"phone"`     // 手机号
	ActivtyId int64  `json:"activtyId"` // 活动ID
	Template  string `json:"template"`  //自定义模版
	PlainTime string `json:"plainTime"` //计划发送时间
}
