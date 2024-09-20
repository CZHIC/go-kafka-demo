package request

import (
	"encoding/json"

	"iv-test/library/logger"
	"reflect"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// request结构体
type BaseRequest struct {
	Data interface{} `json:"data"`
}

// 获取请求参数中的data信息
func GetStructData(ctx *ghttp.Request, data interface{}) (interface{}, error) {
	req := new(BaseRequest)
	req.Data = data
	if err := JsonBody2Req(ctx, req); err != nil {
		logger.Error("----GetStructData---error----", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 获取请求参数中的data信息
func GetJsonStructData(ctx *ghttp.Request, data interface{}) (interface{}, error) {
	if err := JsonBodyToStruct(ctx, data); err != nil {
		return nil, err
	}
	return data, nil
}

// Staff结构体
type StaffInfo struct {
	StaffId   int64  `json:"staff_id"`
	StaffName string `json:"staff_name"`
}

// Json Body转换为对应结构体
func JsonBody2Req(ctx *ghttp.Request, req interface{}) error {

	//原始数据
	content := ctx.GetBodyString()

	//xss转义
	bodyMap := gconv.Map(content)
	if bodyMap != nil {
		data := gconv.Map(bodyMap["data"])
		for k, v := range data {
			if gconv.String(reflect.TypeOf(v)) != "string" {
				continue
			}
			if k == "staff_info" {
				continue
			}
			// data[k], _ = g.View().ParseContent("${.name}", g.Map{ // xss转义
			// 	"name": v,
			// })

			data[k] = v
		}
		tmp := map[string]interface{}{
			"data": data,
		}
		bodyJson, err := json.Marshal(tmp)
		if err == nil {
			content = string(bodyJson)
		}
	}
	if content != "" {
		if err := json.Unmarshal([]byte(content), req); err != nil {
			return err
		}
	}
	return nil
}

// Json Body转换为对应结构体
func JsonBodyToStruct(ctx *ghttp.Request, req interface{}) error {
	content := ctx.GetBodyString()
	logger.Info("JsonBodyToStruct    content    ", content)
	if err := json.Unmarshal([]byte(content), req); err != nil {
		return err
	}
	return nil
}
