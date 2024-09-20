package api

import (
	"bytes"
	"io"
	"iv-test/app/model"
	"iv-test/app/service"
	"iv-test/library/logger"
	"iv-test/library/response"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/xuri/excelize/v2"
)

var Test = new(apiTest)

type apiTest struct{}

// @summary 导入活动文件
// @tags    接口
// @produce json
// @router  /test/upload-file [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiTest) UploadFile(r *ghttp.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		logger.Error(err.Error())
		response.JsonExit(r, 1, "上传文件失败")
	}
	defer file.Close()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		logger.Error(err.Error())
		response.JsonExit(r, 2, "上传文件失败")
	}
	reader := bytes.NewReader(fileBytes)
	f, err := excelize.OpenReader(reader)
	if err != nil {
		logger.Error(err.Error())
		response.JsonExit(r, 3, "上传文件失败")
	}
	sheets := f.GetSheetList()
	importInfo, err := f.GetRows(sheets[0])
	if err != nil {
		logger.Error(err.Error())
		response.JsonExit(r, 4, "上传文件失败")
	}

	for k, data := range importInfo {
		uploadInfo := &model.UploadDoc{}
		if k == 0 {
			continue
		}
		logger.Println("----data-----k", k, data)
		if data[0] == "" {
			continue
		}
		uploadInfo.UserId = gconv.Int64(data[0])

		if len(data) >= 2 && data[1] != "" {
			uploadInfo.Phone = gconv.Int64(data[1])
		}
		if len(data) >= 3 && data[2] != "" {
			uploadInfo.ActivtyId = gconv.Int64(data[2])
		}
		if len(data) >= 4 && data[3] != "" {
			uploadInfo.Template = data[3]
		}
		if len(data) >= 5 && data[4] != "" {
			uploadInfo.SeadTime = data[4]
		}
		sendTime := gtime.NewFromStr(uploadInfo.SeadTime).Unix()
		if sendTime < 0 { // 没传时间，读取活动表时间
			//todo
		}
		if uploadInfo.Template == "" { // 没传模板,读取活动模版
			//todo
		}

		//写入消息表
		//todo

	}

	response.JsonExit(r, 0, "ok")
}

// 定时获取消息表发送MQ
func GetMessage() {
	logger.Println("执行一次")
	service.GetMessage()
}
