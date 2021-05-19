package ios

import "github.com/yzchan/umeng-go/push/notification"

type Unicast struct {
	notification.Cast
	DeviceTokens string  `json:"device_tokens"`
	Payload      Payload `json:"payload"`
}

func NewUnicast() *Unicast {
	cast := &Unicast{}
	cast.Type = "unicast"
	cast.SetProductionMode(true)
	cast.Payload = make(Payload)
	cast.Payload.Initial()
	return cast
}

func (u *Unicast) SetDeviceToken(token string) *Unicast {
	u.DeviceTokens = token
	return u
}
