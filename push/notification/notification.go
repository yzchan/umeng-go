package notification

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/yzchan/umeng-go/v2/push"
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
	TemplateName   string      `json:"template_name,omitempty"`
}

func (n *Notification) SetAppKey(key string) {
	n.AppKey = key
}

func (n *Notification) GetRequestUri() string {
	return push.Host + push.SendPath
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

func (n *Notification) SetDeviceToken(token string) *Notification {
	n.DeviceTokens = token
	return n
}

func (n *Notification) SetDeviceTokens(tokens string) *Notification {
	n.DeviceTokens = tokens
	return n
}

func (n *Notification) SetDeviceTokenList(tokens []string) *Notification {
	if len(tokens) > 500 {
		tokens = tokens[:500]
	}
	n.DeviceTokens = strings.Join(tokens, ",")
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
