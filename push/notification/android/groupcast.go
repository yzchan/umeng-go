package android

import (
	"encoding/json"
	"github.com/yzchan/umeng-go/push/notification"
)

type Groupcast struct {
	notification.Cast
	DeviceTokens string  `json:"device_tokens"`
	Payload      Payload `json:"payload"`
	MiPush
	Filter interface{} `json:"filter"`
}

func NewGroupcast() *Groupcast {
	cast := &Groupcast{}
	cast.Type = "unicast"
	cast.SetProductionMode(true)
	cast.Payload.Initial()
	cast.Payload.SetDisplayType("notification")
	return cast
}

func (g *Groupcast) SetFilter(condition string) *Groupcast {
	var v interface{}
	json.Unmarshal([]byte(condition), &v)
	g.Filter = v
	return g
}
