package android

import (
	"github.com/yzchan/umeng-go/push/notification"
	"strings"
)

type Listcast struct {
	notification.Cast
	DeviceTokens string  `json:"device_tokens"`
	Payload      Payload `json:"payload"`
	MiPush
}

func NewListcast() *Listcast {
	cast := &Listcast{}
	cast.Type = "listcast"
	cast.SetProductionMode(true)
	cast.Payload.Initial()
	cast.Payload.SetDisplayType("notification")
	return cast
}

func (l *Listcast) SetDeviceTokens(tokens []string) *Listcast {
	l.DeviceTokens = strings.Join(tokens, ",")
	return l
}
