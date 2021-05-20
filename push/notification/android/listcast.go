package android

import (
	"github.com/yzchan/umeng-go/push/notification"
	"strings"
	"time"
)

type Listcast struct {
	notification.Cast
	DeviceTokens string      `json:"device_tokens"`
	Payload      Payload     `json:"payload"`
	Policy       Policy      `json:"policy,omitempty"`
	Filter       interface{} `json:"filter"`
	MiPush
}

func NewListcast() *Listcast {
	cast := &Listcast{}
	cast.Type = "listcast"
	cast.Payload.Initial()
	cast.Timestamp = time.Now().Unix()
	return cast
}

func (l *Listcast) SetDeviceTokens(tokens []string) *Listcast {
	if len(tokens) > 500 {
		tokens = tokens[:500]
	}
	l.DeviceTokens = strings.Join(tokens, ",")
	return l
}
