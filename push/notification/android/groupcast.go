package android

import (
	"encoding/json"
	"github.com/yzchan/umeng-go/push/notification"
)

type Groupcast struct {
	notification.Cast
	DeviceTokens string      `json:"device_tokens"`
	Payload      Payload     `json:"payload"`
	Policy       Policy      `json:"policy,omitempty"`
	Filter       interface{} `json:"filter"`
	MiPush
}

func NewGroupcast() *Groupcast {
	cast := &Groupcast{}
	cast.Type = "groupcast"
	cast.Payload.Initial()
	return cast
}

func (g *Groupcast) SetFilter(condition string) *Groupcast {
	var v interface{}
	_ = json.Unmarshal([]byte(condition), &v)
	g.Filter = v
	return g
}

func (g *Groupcast) Send() (string, error) {
	g.SetPackageName(g.App.PackageName)
	return g.BaseSend(g)
}
