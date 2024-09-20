package api

import (
	"iv-test/library/logger"
	"iv-test/library/msg"
	"iv-test/library/response"
	"os"
	"strings"
	"sync"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Test = new(apiTest)

type apiTest struct{}

// @summary 导入财务文件
// @tags    接口
// @produce json
// @router  /upload/financial-file [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiTest) FinancialFile(r *ghttp.Request) {

	var wg sync.WaitGroup
	ProjectName := r.GetString("companyName")
	period := r.GetString("period") //周期
	period = period[0:4] + " " + period[4:]
	logger.Print("----周期----", period)
	var ok bool
	var ProjectId int
	ProjectNames := strings.Split(ProjectName, ".")
	if len(ProjectNames) == 1 {
		ProjectName = ProjectNames[0]
	} else {
		ProjectName = ProjectName[len(ProjectNames[0])+1:]
	}
	ProjectName = strings.TrimSpace(ProjectName)
	logger.Print("----Project_name----", ProjectName)
	env := g.Cfg().GetString("ENV.env")
	if env == "test" {
		if ProjectId, ok = msg.NameToProjectId_Test[ProjectName]; !ok {
			ProjectId = 0
		}
	}
	if env == "release" {
		if ProjectId, ok = msg.NameToProjectId_Release[ProjectName]; !ok {
			ProjectId = 0
		}
	}

	if env == "pro" {
		mapList := map[string]int{}
		for k, v := range msg.NameToProjectId_Pro {
			key := strings.ToLower(k)
			mapList[key] = v
		}
		ProjectNameLower := strings.ToLower(ProjectName)

		if ProjectId, ok = mapList[ProjectNameLower]; !ok {
			ProjectId = 0
		}

		if ProjectId == 0 { // 全值匹配没找到时，用包涵关系再试一次 （用于中国公司）
			for k, v := range msg.NameToProjectId_Pro {
				key := strings.ToLower(k)
				projectName := strings.ToLower(ProjectName)
				if strings.Contains(projectName, key) {
					ProjectId = v
					break
				}

			}
		}

	}
	if ProjectId == 0 {
		response.JsonExit(r, 2, "error", "未查到项目信息 ||"+ProjectName)
	}

	files := r.GetUploadFiles("file")
	logger.Println("------文件信息列表----", files)

	Errors := []string{}
	FinancialFile := ""
	for _, v := range files {
		wg.Add(1)
		go func(v *ghttp.UploadFile, ProjectName string, wg *sync.WaitGroup) {
			name, err := v.Save("document/")
			if err != nil {
				logger.Error("保存临时文件失败：", err.Error(), name, ProjectName)
				return
			}
			file, err := os.OpenFile("document/"+name, os.O_RDONLY, 0600)
			if err != nil {
				logger.Error("打开临时文件失败：", err.Error(), name, ProjectName)
				return
			}
			defer wg.Done()
			defer os.Remove("document/" + name)
			defer file.Close()

			//	fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				logger.Error("获取文件类容失败：", err.Error(), name, ProjectName)
				return
			}
			// fileId, err := service.UploadFinancialFile(v, projectInfo, fileBytes)
			// if err != nil {
			// 	logger.Error(err.Error())
			// 	Errors = append(Errors, err.Error())
			// }

			// if fileId > 0 {
			// 	FinancialFile += gconv.String(fileId) + ","
			// }
		}(v, ProjectName, &wg)
	}
	wg.Wait() // 等待所有协程结束
	FinancialFile = strings.Trim(FinancialFile, ",")
	if len(FinancialFile) > 0 {
		// err := service.AddFinancial(int64(ProjectId), ProjectName, FinancialFile, period)
		// if err != nil {
		// 	Errors = append(Errors, err.Error())
		// }
	}

	logger.Print("---------所有进程结束--------")
	if len(Errors) > 0 {
		response.JsonExit(r, 1, "error", Errors)
	}
	response.JsonExit(r, 0, "ok")
}
