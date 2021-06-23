package notification

import (
	"encoding/json"
	"github.com/yzchan/umeng-go/push"
	"time"
)

type Cast struct {
	App            *push.App `json:"-"`
	AppKey         string    `json:"appkey"`
	Type           string    `json:"type"`
	Timestamp      int64     `json:"timestamp"`
	Description    string    `json:"description"`
	ProductionMode string    `json:"production_mode,omitempty"` // 默认为true
	ReceiptUrl     string    `json:"receipt_url,omitempty"`     // U-Push Pro 回执地址 最大长度256字节。
	ReceiptType    string    `json:"receipt_type,omitempty"`    // U-Push Pro 回执类型 1：送达回执；2：点击回执；3：送达和点击/忽略回执。默认为3
}

type CastResp struct {
	Ret  string `json:"ret"`
	Data struct {
		MsgId  string `json:"msg_id"`
		TaskId string `json:"task_id"`
	} `json:"data"`
}

func (c *Cast) SetAppkey(key string) *Cast {
	c.AppKey = key
	return c
}

func (c *Cast) SetTimestamp() *Cast {
	c.Timestamp = time.Now().Unix()
	return c
}

func (c *Cast) SetProductionMode(mode bool) *Cast {
	if !mode {
		c.ProductionMode = "false"
	}
	return c
}

func (c *Cast) SetDescription(desc string) *Cast {
	c.Description = desc
	return c
}

func (c *Cast) SetReceipt(url string, rType string) *Cast {
	c.ReceiptType = rType
	c.ReceiptUrl = url
	return c
}

func (c *Cast) BindApp(app *push.App) *Cast {
	c.App = app
	return c
}

func (c *Cast) BaseSend(cast interface{}) (taskId string, err error) {
	var (
		buf []byte
		r   CastResp
	)

	c.AppKey = c.App.AppKey
	c.SetTimestamp()

	push.PrettyJson = true
	defer func() {
		push.PrettyJson = false
	}()
	if buf, err = c.App.Request(push.Host+push.SendPath, cast); err != nil {
		return
	}

	if err = json.Unmarshal(buf, &r); err != nil {
		return
	}

	return r.Data.MsgId + r.Data.TaskId, nil
}
