package logger

import (
	"fmt"
	"iv-test/library/email"
	"strings"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

// 获取log对象
func Log(name ...string) *glog.Logger {
	if len(name) > 0 && name[0] != "" {
		return g.Log(name[0]).Skip(1).Line()
	}
	return g.Log().Skip(1).Line()
}

// 普通打印，不带任何标签
func Print(v ...interface{}) {
	Log("access").Print(v)
}

// 普通打印，不带任何标签
func Println(v ...interface{}) {
	Log("access").Println(v)
}

// INFO打印，带[INFO]标签
func Info(v ...interface{}) {
	Log("access").Info(v)
}

// INFO打印，带[INFO]标签
func Infof(format string, v ...interface{}) {
	Log("access").Infof(format, v...)
}

// Debug打印，带[Debug]标签
func Debug(v ...interface{}) {
	Log("access").Debug(v)
}

// Debug打印，带[Debug]标签
func Debugf(format string, v ...interface{}) {
	Log("access").Debugf(format, v...)
}

// Error打印，带[Error]标签
func Error(v ...interface{}) {
	Log("error").Error(v)
	SendErrorEmail("", v...)
}

// Error打印，带[Error]标签
func Errorf(format string, v ...interface{}) {
	Log("error").Errorf(format, v...)
	SendErrorEmail("", v...)
}

// 系统错误信息发邮件提醒
func SendErrorEmail(format string, v ...interface{}) {
	//获取发送人
	users := g.Cfg().Get("ENV.error")
	if users == nil {
		return
	}
	user := users.([]interface{})
	if len(user) == 0 {
		return
	}

	EmailInfo := new(email.ReqEmail)
	EmailInfo.Title = "【IV-CRON】系统报错 :" + g.Cfg().GetString("ENV.env")
	EmailInfo.Content = format + "<br>"
	for _, str := range v {
		EmailInfo.Content += gconv.String(str) + "<br>"
	}
	var (
		tempStr  = ""
		valueStr = ""
	)

	//传值v是error类型时，获取最底层堆栈
	for _, v := range v {
		if err, ok := v.(error); ok {
			tempStr = fmt.Sprintf("%+v", err)
		} else {
			tempStr = gconv.String(v)
		}
		if len(valueStr) > 0 {
			if valueStr[len(valueStr)-1] == '\n' {
				// Remove one blank line(\n\n).
				if tempStr[0] == '\n' {
					valueStr += tempStr[1:]
				} else {
					valueStr += tempStr
				}
			} else {
				valueStr += " " + tempStr
			}
		} else {
			valueStr = tempStr
		}
	}
	valueStr = strings.Replace(gconv.String(valueStr), "\n", "<br>", 100)
	EmailInfo.Content += "<br> Stack: <hr>" + valueStr + "<br>"

	//Stack  针对传值v不是error时，获取堆栈信息
	Stack := []interface{}{}
	if s := Log("error").GetStack(); s != "" {
		Stack = append(Stack, s)
	}
	for _, strs := range Stack {
		strs = strings.Replace(gconv.String(strs), "\n", "<br>", 100)
		EmailInfo.Content += gconv.String(strs) + "<br>"
	}

	TO := ""
	for _, v := range user {
		TO += gconv.String(v) + ","
	}
	TO = strings.TrimRight(TO, ",")
	EmailInfo.To = TO
	EmailInfo.CC = ""
	_, err := email.SendEmail(EmailInfo)
	if err != nil { //如果失败重发一次
		fmt.Println("-------邮件发送失败----------:", err.Error())
		time.Sleep(1 * time.Second)
		email.SendEmail(EmailInfo)
	}
}
