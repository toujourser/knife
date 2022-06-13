package im_hook

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/toujourser/utils/algorithm"
	"github.com/toujourser/utils/request"
)

type FeiShuMsgType int

const (
	FText     FeiShuMsgType = 1 // 普通文本消息
	FFullText FeiShuMsgType = 2 // 富文本消息
)

func NewFeishuHook(token, url string) *FeiShuHook {
	return &FeiShuHook{Enable: true, Token: token, Url: url}
}

type FeiShuHook struct {
	Enable bool
	Token  string
	Url    string
}

type FeiShuResponse struct {
	Extra         interface{} `json:"Extra"`
	StatusCode    int         `json:"StatusCode"`
	StatusMessage string      `json:"StatusMessage"`
	Code          int         `json:"code"`
	Msg           string      `json:"msg"`
	Data          interface{} `json:"data"`
}

//飞书自定义机器人, 目前仅支持文字和富文本
//	msgType: 1-纯文本 2-富文本
//		msgType为1时，title为发送的文字信息，content置空
//		msgType为2时，title为富文本标题，content为富文本内容
func (f *FeiShuHook) Push(msgType FeiShuMsgType, title, content string) error {
	if !f.Enable {
		return nil
	}
	var (
		data      string
		timestamp = time.Now().Unix()
	)
	sign, err := algorithm.GenSign(f.Token, timestamp)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to generate a signature, msg: %s", err))
	}
	switch msgType {
	case FText:
		data = fmt.Sprintf(`{"timestamp": "%d","sign": "%s","msg_type":"text","content":{"text":"%s"}}`, timestamp, sign, title)
	case FFullText:
		data = fmt.Sprintf(`{"timestamp": "%d","sign": "%s", "msg_type": "post","content": {"post": {"zh_cn": {"title": "%s","content": [%s]}}}}`, timestamp, sign, title, content)
	}

	resp, err := request.Request("POST", f.Url, map[string]string{"Content-Type": "application/json"}, []byte(data))
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fRsp := new(FeiShuResponse)
	err = json.Unmarshal(b, fRsp)
	if err != nil {
		return errors.New(fmt.Sprintf("feishu response unmarshal failed. err: %s", err))
	}
	if fRsp.Code != 0 || fRsp.StatusCode != 0 {
		return errors.New(fmt.Sprintf("feishu send msg failed, err: %s", fRsp.Msg))
	}
	return nil
}
