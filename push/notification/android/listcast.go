package android

import (
	"encoding/json"
	"github.com/yzchan/umeng-go/push/notification"
	"strings"
)

type Listcast struct {
	notification.Cast
	DeviceTokens string      `json:"device_tokens"`
	Payload      Payload     `json:"payload"`
	Policy       Policy      `json:"policy,omitempty"`
	Filter       interface{} `json:"filter"`
	Channel      Channel     `json:"channel_properties,omitempty"`
	MiPush
}

func NewListcast() *Listcast {
	cast := &Listcast{}
	cast.Type = "listcast"
	cast.Payload.Initial()
	return cast
}

func (cast *Listcast) SetDeviceTokens(tokens []string) *Listcast {
	if len(tokens) > 500 {
		tokens = tokens[:500]
	}
	cast.DeviceTokens = strings.Join(tokens, ",")
	return cast
}

func (cast *Listcast) Send() (string, error) {
	cast.Channel.SetChannelActivity(cast.App.PackageName)
	return cast.BaseSend(cast)
}

func (cast *Listcast) String() string {
	marshaled, _ := json.MarshalIndent(cast, "", "    ")
	return string(marshaled)
}
