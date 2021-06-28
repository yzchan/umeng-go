package android

import (
	"encoding/json"
	"github.com/yzchan/umeng-go/push/notification"
)

type Broadcast struct {
	notification.Cast
	Payload Payload `json:"payload"`
	Policy  Policy  `json:"policy,omitempty"`
	Channel Channel `json:"channel_properties,omitempty"`
	MiPush
}

func NewBroadcast() *Broadcast {
	cast := &Broadcast{}
	cast.Type = "broadcast"
	cast.Payload.Initial()
	return cast
}

func (cast *Broadcast) Send() (string, error) {
	cast.Channel.SetChannelActivity(cast.App.PackageName)
	return cast.BaseSend(cast)
}

func (cast *Broadcast) String() string {
	marshaled, _ := json.Marshal(cast)
	return string(marshaled)
}
