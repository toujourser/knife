package im_hook

import "testing"

func TestFeiShuHook(t *testing.T) {
	hook := NewFeishuHook("xx", "https://open.feishu.cn/open-apis/bot/v2/hook/xxxx")
	err := hook.Push(1, "发送纯文本", "")
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log("text ok")
	}

	err = hook.Push(2, "发送富文本", `[{"tag": "text","text": "【搜索引擎】\n"},{"tag": "a","text": "百度\n","href": "https://www.baidu.com"},{"tag": "a","text": "必应","href": "https://cn.bing.com/"}]`)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log("fulltext ok")
	}
}
