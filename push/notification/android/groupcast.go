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
	Channel      Channel     `json:"channel_properties,omitempty"`
}

func NewGroupcast() *Groupcast {
	cast := &Groupcast{}
	cast.Type = "groupcast"
	cast.Payload.Initial()
	return cast
}

func (cast *Groupcast) SetFilter(condition string) *Groupcast {
	var v interface{}
	_ = json.Unmarshal([]byte(condition), &v)
	cast.Filter = v
	return cast
}

func (cast *Groupcast) Send() (string, error) {
	cast.Channel.SetChannelActivity(cast.App.PackageName)
	return cast.BaseSend(cast)
}

func (cast *Groupcast) String() string {
	marshaled, _ := json.MarshalIndent(cast, "", "    ")
	return string(marshaled)
}
