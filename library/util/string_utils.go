package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"iv-test/library/logger"
	"iv-test/library/msg"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/gogf/gf/util/gconv"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// 正则表达式拆分字符串
func RegSplit(text string, delimeter string) []string {
	reg := regexp.MustCompile(delimeter)
	indexes := reg.FindAllStringIndex(text, -1)
	lastStart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = text[lastStart:element[0]]
		lastStart = element[1]
	}
	result[len(indexes)] = text[lastStart:len(text)]
	return result
}

// 正则表达式提取字符串
func RegFind(text string, delimeter string) string {
	reg := regexp.MustCompile(delimeter)
	return reg.FindString(text)
}

// 十六进制字符串转十进制数字
func HexToDec(hex string) int64 {
	n, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		logger.Errorf("HexToDec -> strconv.ParseInt() execute failed. err = %v", err)
	}
	return n
}

// 把list转换string
func AnalysisListToString(req []int) string {

	if req == nil {
		return ""
	}

	var buffer bytes.Buffer
	for index, info := range req {
		buffer.WriteString(strconv.Itoa(info))
		if len(req) > 1 && index != len(req)-1 {
			buffer.WriteString(",")
		}
	}

	return buffer.String()
}

// 把list转换string
func ListToString(req []int64) string {
	if len(req) > 0 {
		str := ""
		for _, v := range req {
			str += gconv.String(v) + ","
		}
		str = strings.TrimRight(str, ",")
		return str
	}
	return ""
}

// 把list转换string
func AnalysisStringToList(req string) []int {

	rsp := make([]int, 0)

	if req == "" {
		return rsp
	}

	for _, info := range strings.Split(req, ",") {
		int, _ := strconv.Atoi(info)
		if int == 0 {
			continue
		}
		rsp = append(rsp, int)
	}

	return rsp
}

// 把string转换Map
func AnalysisStringToMap(req string) map[int64]struct{} {

	rsp := make(map[int64]struct{})

	if req == "" {
		return rsp
	}

	for _, info := range AnalysisStringToList(req) {
		rsp[int64(info)] = struct{}{}
	}

	return rsp
}

// 把list转换string
func AnalysisStringToListInt64(req string) []int64 {

	rsp := make([]int64, 0)

	if req == "" {
		return rsp
	}

	for _, info := range strings.Split(req, ",") {
		int64, _ := strconv.ParseInt(info, 10, 64)
		if int64 == 0 {
			continue
		}
		rsp = append(rsp, int64)
	}

	return rsp
}

// []string 去重
//func RemoveDuplicate(list []string) []string {
//	// 这个排序很关键
//	sort.Strings(list)
//	i := 0
//	var newlist = []string{""}
//	for j := 0; j < len(list); j++ {
//		if strings.Compare(newlist[i], list[j]) == -1 {
//			newlist = append(newlist, list[j])
//			i++
//		}
//	}
//	return newlist
//}

// []string 去重
func RemoveDuplicate(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// []int 去重
func RemoveDuplicateInt(list *[]int) []int {
	var x []int = []int{}
	for _, i := range *list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	return x
}

// excel 列坐标规则生成 -- 26 英文字母
func ExcelCoordinateGenerate(index int) string {
	//index := 0

	initNum := 64 // 65-A

	// 一层坐标
	wIndex := index / 26
	// 二层坐标
	zIndex := index % 26
	if wIndex == 0 {
		logger.Info(fmt.Sprintf("%c", zIndex+initNum+1))
		return fmt.Sprintf("%c", zIndex+initNum+1)
	} else {
		logger.Info(fmt.Sprintf("%c", wIndex+initNum) + fmt.Sprintf("%c", zIndex+initNum+1))
		return fmt.Sprintf("%c", wIndex+initNum) + fmt.Sprintf("%c", zIndex+initNum+1)
	}
}

func Interface2String(value interface{}) string {
	var key string
	if value == nil {
		return ""
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
		if key == "0" {
			key = ""
		}
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
		if key == "0" {
			key = ""
		}
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
		if key == "0" {
			key = ""
		}
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
		if key == "0" {
			key = ""
		}
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
		if key == "0" {
			key = ""
		}
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
		if key == "0" {
			key = ""
		}
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
		if key == "0" {
			key = ""
		}
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
		if key == "0" {
			key = ""
		}
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
		if key == "0" {
			key = ""
		}
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
		if key == "0" {
			key = ""
		}
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
		if key == "0" {
			key = ""
		}
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
		if key == "0" {
			key = ""
		}
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}

// 过滤对比字段 只对比存在值的(投后)
func FilterInterfaceParam(value interface{}, menuType int) map[string]interface{} {

	rsp := make(map[string]interface{}, 0)
	jsonMarshalA, _ := json.Marshal(value)
	err := json.Unmarshal(jsonMarshalA, &rsp)
	logger.Info(err)
	jsonMarshalB, _ := json.Marshal(rsp)
	logger.Println(string(jsonMarshalB))

	for k := range rsp {
		if menuType == msg.InvestmentTeamMember {
			if _, ok := msg.TeamUserParamMap[k]; !ok {
				delete(rsp, k)
			}
		} else if menuType == msg.SituationInfo {
			if _, ok := msg.SituationParamMap[k]; !ok {
				delete(rsp, k)
			}
		}
	}

	return rsp
}

type HighlightBody struct {
	Id          int64  `json:"id"`
	ProjectId   int64  `json:"projectId"`
	ProjectType int    `json:"projectType"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Index       int    `json:"index"`
	IsSelf      bool   `orm:"is_self" json:"isSelf"`
}

func SplitFloatStringByIndex(s string) string {
	if s != "" {
		if s[len(s)-3:] == ".00" {
			return s[:len(s)-3]
		} else if s[len(s)-1:] == "0" {
			return s[:len(s)-1]
		}
	}
	return s
}

// 排序
func OrderByString(req string) string {
	if req == "" {
		return req
	}
	list := strings.Split(req, ",")
	sort.Strings(list)

	return strings.Join(list, ",")
}

func UTF82GBK(src string) ([]byte, error) {
	encoding := simplifiedchinese.All[0]
	return ioutil.ReadAll(transform.NewReader(bytes.NewBuffer([]byte(src)), encoding.NewEncoder()))
}

// 检查参数类型
func CheckParamString(value interface{}) bool {

	//if value == nil || value == "" || value == 0 {
	//	return true
	//}

	switch value.(type) {
	case string:
		return value == "##--.--##"
	default:
		return false
	}
}

// 检查参数类型
func CheckParamNum(value interface{}) bool {

	//if value == nil || value == "" || value == 0 {
	//	return true
	//}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		return ft == -999999
	case float32:
		ft := value.(float32)
		return ft == -999999
	case int:
		it := value.(int)
		return it == -999999
	case int64:
		it := value.(int64)
		return it == -999999
	default:
		return false
	}
}
