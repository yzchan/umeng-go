package android

import (
	"github.com/yzchan/umeng-go/push/notification"
	"time"
)

type Unicast struct {
	notification.Cast
	DeviceTokens string  `json:"device_tokens"`
	Payload      Payload `json:"payload"`
	Policy       Policy  `json:"policy,omitempty"`
	MiPush
}

func NewUnicast() *Unicast {
	cast := &Unicast{}
	cast.Type = "unicast"
	cast.Payload.Initial()
	cast.Timestamp = time.Now().Unix()
	return cast
}

func (u *Unicast) SetDeviceToken(token string) *Unicast {
	u.DeviceTokens = token
	return u
}
