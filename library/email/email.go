package email

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/pkg/errors"
	"github.com/rs/xid"
)

type ReqEmail struct {
	Title      string `json:"Title"`      // 标题
	Content    string `json:"Content"`    // 标题
	To         string `json:"To"`         // 标题
	AppendPath string `json:"AppendPath"` // 附件，多个文件以;隔开
	CC         string `json:"CC"`         // 抄送`,`隔开
}
type Tof4EmailInfo struct {
	To string `json:"To"` //邮件接受者

	CC string `json:"CC"` //抄送人(邮箱格式),支持多人，按","分隔

	Bcc string `json:"Bcc"` //密送人(邮箱格式),支持多人，按","分隔

	From string `json:"From"` //发送者

	Title string `json:"Title"` //标题

	Content string `json:"Content"` //发送内容

	Priority int `json:"Priority"` //发送优先级 0：普通 1: 高

	EmailType int `json:"EmailType"` //指定邮件类型，默认1。（0 外部邮件 1内部邮件）

	BodyFormat int `json:"BodyFormat"` //内容格式，默认1（0 文本 1 html格式）

	AppendPath string `json:"AppendPath"` //附件路径 多个路径以;隔开 例：/doc/golang.png;/doc/golang1.png;
}

type Tof4Rsp struct {
	ErrCode int    `json:"ErrCode"`
	ErrMsg  string `json:"ErrMsg"`
	MsgId   string `json:"MsgId"`
}

//发送普通邮件
func SendEmail(req *ReqEmail) (bool, error) {

	url := g.Cfg().GetString("TOF.Host") + g.Cfg().GetString("TOF.EmailPath")

	postData := Tof4EmailInfo{}
	postData.From = g.Cfg().GetString("TOF.From")
	postData.To = req.To
	postData.Title = req.Title
	postData.Content = req.Content
	postData.CC = req.CC
	postData.BodyFormat = 1
	postData.EmailType = 0

	ret, err := tOf4Post(&postData, url)
	if err != nil {
		return false, err
	}

	return ret, nil
}

//带附件发送邮件
func SendEmailWith(req *ReqEmail) (bool, error) {
	url := g.Cfg().GetString("TOF.Host") + g.Cfg().GetString("TOF.EmailWithPath")

	postData := Tof4EmailInfo{}
	postData.From = g.Cfg().GetString("TOF.From")
	postData.To = req.To
	postData.Title = req.Title
	postData.Content = req.Content
	postData.CC = req.CC
	postData.BodyFormat = 1
	postData.EmailType = 0
	//dir, _ := os.Getwd()
	//postData.AppendPath = dir + "/doc/golang.png;" + dir + "/doc/golang1.png"
	postData.AppendPath = req.AppendPath

	ret, err := sendWith(&postData, url)
	if err != nil {
		return false, err
	}
	return ret, nil
}

func sendWith(postData *Tof4EmailInfo, url string) (bool, error) {
	var b bytes.Buffer
	writer, err := multipartWriteFields(&b, postData)
	if err != nil {
		return false, errors.WithMessage(err, "")
	}
	writer.Close()

	secret := g.Cfg().GetString("TOF.Token")
	headerTimestamp := time.Now().Unix()
	headerNonce := xid.New().String()
	sign := tofSign(secret, gconv.String(headerTimestamp), headerNonce)

	cc := g.Client()
	cc.SetHeader("x-rio-paasid", g.Cfg().GetString("TOF.PassId"))
	cc.SetHeader("x-rio-signature", sign)
	cc.SetHeader("x-rio-timestamp", gconv.String(headerTimestamp))
	cc.SetHeader("x-rio-nonce", headerNonce)
	cc.SetHeader("Content-Type", writer.FormDataContentType())
	cc.Timeout(100 * time.Second) // 设置超时时间

	if err != nil {
		return false, err
	}
	s, err := cc.Post(url, &b)
	if err != nil {
		return false, err
	}
	body := s.ReadAllString()
	fmt.Print("Email发送类容:", postData, "结果:", body)
	result := Tof4Rsp{}
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return false, errors.Wrap(err, fmt.Sprintf("rsp: %s", string(body)))
	}
	//发送失败才返回错误码
	if result.ErrCode != 0 {
		return false, errors.New(fmt.Sprintf("rsp: %s", string(body)))
	}
	return true, nil

}

func tOf4Post(postData *Tof4EmailInfo, url string) (bool, error) {

	secret := g.Cfg().GetString("TOF.Token")
	headerTimestamp := time.Now().Unix()
	headerNonce := xid.New().String()
	sign := tofSign(secret, gconv.String(headerTimestamp), headerNonce)
	postStr, _ := json.Marshal(postData)

	cc := g.Client()
	cc.SetHeader("x-rio-paasid", g.Cfg().GetString("TOF.PassId"))
	cc.SetHeader("x-rio-signature", sign)
	cc.SetHeader("x-rio-timestamp", gconv.String(headerTimestamp))
	cc.SetHeader("x-rio-nonce", headerNonce)
	cc.Timeout(100 * time.Second) // 设置超时时间

	s, err := cc.Post(url, string(postStr))
	if err != nil {
		return false, err
	}
	body := s.ReadAllString()
	fmt.Print("Email发送类容:", postData, "结果:", body)
	result := Tof4Rsp{}
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return false, errors.Wrap(err, fmt.Sprintf("rsp: %s", string(body)))
	}
	//发送失败才返回错误码
	if result.ErrCode != 0 {
		return false, errors.New(fmt.Sprintf("rsp: %s", string(body)))
	}
	return true, nil
}

func tofSign(secret, timestamp, nonce string) string {
	signData := fmt.Sprintf("%s%s%s%s", timestamp, secret, nonce, timestamp)
	res := strings.ToUpper(fmt.Sprintf("%x", sha256.Sum256([]byte(signData))))
	return res
}

func multipartWriteFields(b *bytes.Buffer, mailInfo *Tof4EmailInfo) (*multipart.Writer, error) {
	w := multipart.NewWriter(b)
	// 使用反射来读取struct字段和值
	typ := reflect.TypeOf(*mailInfo)
	values := reflect.ValueOf(*mailInfo)
	num := values.NumField()

	for i := 0; i < num; i++ {
		fieldName := typ.Field(i).Tag.Get("json")
		fieldKind := values.Field(i).Kind()
		// 只有string和int两种
		var err error
		switch fieldKind {
		case reflect.String:
			err = w.WriteField(fieldName, values.Field(i).String())
		case reflect.Int:
			err = w.WriteField(fieldName, strconv.FormatInt(values.Field(i).Int(), 10))
		default:
			err = errors.New("unknown field kind")
		}
		if err != nil {
			return nil, errors.Errorf("%v", err)
		}
	}
	// 创建文件字段，多附件上传
	appendPathList := strings.Split(mailInfo.AppendPath, ";")
	fmt.Print(len(appendPathList))
	lenNum := len(appendPathList)
	for _, path := range appendPathList[0:lenNum] {
		file, err := os.Open(path)

		if err != nil {
			fmt.Print("读取文件出错了:", err.Error())
			break
		}
		defer file.Close()
		fileWriter, err := w.CreateFormFile("file", filepath.Base(path))
		if err != nil {
			return nil, errors.WithMessage(err, "create form file err")
		}
		_, err = io.Copy(fileWriter, file)
		if err != nil {
			return nil, errors.WithMessage(err, "writing file err")
		}
	}
	return w, nil
}
