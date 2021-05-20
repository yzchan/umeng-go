package notification

import "time"

type Caster interface {
	SetAppkey(key string) *Cast
}

type Cast struct {
	Appkey         string `json:"appkey"`
	Type           string `json:"type"`
	Timestamp      int64  `json:"timestamp"`
	Description    string `json:"description"`
	ProductionMode string `json:"production_mode,omitempty"` // 默认为true
	ReceiptUrl     string `json:"receipt_url,omitempty"`     // U-Push Pro 回执地址 最大长度256字节。
	ReceiptType    string `json:"receipt_type,omitempty"`    // U-Push Pro 回执类型 1：送达回执；2：点击回执；3：送达和点击/忽略回执。默认为3
}

func (u *Cast) SetAppkey(key string) *Cast {
	u.Appkey = key
	return u
}

func (u *Cast) SetTimestamp() *Cast {
	u.Timestamp = time.Now().Unix()
	return u
}

func (u *Cast) SetProductionMode(mode bool) *Cast {
	if !mode {
		u.ProductionMode = "false"
	}
	return u
}

func (u *Cast) SetDescription(desc string) *Cast {
	u.Description = desc
	return u
}

func (u *Cast) SetReceipt(url string, rType string) *Cast {
	u.ReceiptType = rType
	u.ReceiptUrl = url
	return u
}
