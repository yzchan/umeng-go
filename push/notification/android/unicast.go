package android

import "github.com/yzchan/umeng-go/push/notification"

type Unicast struct {
	notification.Cast
	DeviceTokens string  `json:"device_tokens"`
	Payload      Payload `json:"payload"`
	MiPush
}

func NewUnicast() *Unicast {
	cast := &Unicast{}
	cast.Type = "unicast"
	cast.SetProductionMode(true)
	cast.Payload.Initial()
	cast.Payload.SetDisplayType("notification")
	return cast
}

func (u *Unicast) SetDeviceToken(token string) *Unicast {
	u.DeviceTokens = token
	return u
}
