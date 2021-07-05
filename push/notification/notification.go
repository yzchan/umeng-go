package notification

import (
	"encoding/json"
	"github.com/yzchan/umeng-go/push"
	"time"
)

const (
	Unicast        string = "unicast"
	Listcast       string = "listcast"
	Broadcast      string = "broadcast"
	Groupcast      string = "groupcast"
	Filecast       string = "filecast"
	Customizedcast string = "customizedcast"
)

type Notification struct {
	App            *push.App   `json:"-"`
	AppKey         string      `json:"appkey"`
	Type           string      `json:"type"`
	Timestamp      int64       `json:"timestamp"`
	Description    string      `json:"description"`
	DeviceTokens   string      `json:"device_tokens,omitempty"`
	AliasType      string      `json:"alias_type,omitempty"`
	Alias          string      `json:"alias,omitempty"`
	FileId         string      `json:"file_id,omitempty"`
	Filter         interface{} `json:"filter,omitempty"`
	Policy         *Policy     `json:"policy,omitempty"`
	ProductionMode string      `json:"production_mode,omitempty"` // 默认为true
	ReceiptUrl     string      `json:"receipt_url,omitempty"`     // U-Push Pro 回执地址 最大长度256字节。
	ReceiptType    string      `json:"receipt_type,omitempty"`    // U-Push Pro 回执类型 1：送达回执；2：点击回执；3：送达和点击/忽略回执。默认为3
}

type CastResp struct {
	Ret  string `json:"ret"`
	Data struct {
		MsgId  string `json:"msg_id"`
		TaskId string `json:"task_id"`
	} `json:"data"`
}

func (n *Notification) SetAppKey(key string) *Notification {
	n.AppKey = key
	return n
}

func (n *Notification) InitTimestamp() *Notification {
	n.Timestamp = time.Now().Unix()
	return n
}

func (n *Notification) SetProductionMode(mode bool) *Notification {
	if !mode {
		n.ProductionMode = "false"
	}
	return n
}

func (n *Notification) SetDescription(desc string) *Notification {
	n.Description = desc
	return n
}

func (n *Notification) SetAliasType(aliasType string) *Notification {
	n.AliasType = aliasType
	return n
}

func (n *Notification) SetAlias(alias string) *Notification {
	n.Alias = alias
	return n
}

func (n *Notification) SetFileId(fileId string) *Notification {
	n.FileId = fileId
	return n
}

func (n *Notification) SetFilter(condition string) *Notification {
	var v interface{}
	_ = json.Unmarshal([]byte(condition), &v)
	n.Filter = v
	return n
}

func (n *Notification) SetDeviceTokens(tokens string) *Notification {
	n.DeviceTokens = tokens
	return n
}

func (n *Notification) SetReceipt(url string, rType string) *Notification {
	n.ReceiptType = rType
	n.ReceiptUrl = url
	return n
}

func (n *Notification) SetNotificationType(t string) *Notification {
	n.Type = t
	return n
}

func (n *Notification) BindApp(app *push.App) *Notification {
	n.App = app
	n.SetAppKey(app.AppKey)
	return n
}

func (n *Notification) BaseSend(notification interface{}) (taskOrMsgId string, err error) {
	var (
		buf []byte
		r   CastResp
	)
	n.InitTimestamp()

	if buf, err = n.App.Request(push.Host+push.SendPath, notification); err != nil {
		return
	}

	if err = json.Unmarshal(buf, &r); err != nil {
		return
	}

	return r.Data.MsgId + r.Data.TaskId, nil
}
