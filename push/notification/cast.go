package notification

import "time"

type Caster interface {
	SetAppkey(key string) *Cast
	SetTimestamp() *Cast
}

type Cast struct {
	Appkey         string `json:"appkey"`
	Type           string `json:"type"`
	Timestamp      int64  `json:"timestamp"`
	Description    string `json:"description"`
	ProductionMode bool   `json:"production_mode"`
	ReceiptUrl     string `json:"receipt_url,omitempty"`  // U-Push Pro 回执地址
	ReceiptType    string `json:"receipt_type,omitempty"` // U-Push Pro 回执类型
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
	u.ProductionMode = mode
	return u
}

func (u *Cast) SetDescription(desc string) *Cast {
	u.Description = desc
	return u
}
