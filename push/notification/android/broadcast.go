package android

import "github.com/yzchan/umeng-go/push/notification"

type Broadcast struct {
	notification.Cast
	Payload      Payload `json:"payload"`
	MiPush
}

func NewBroadcast() *Broadcast {
	cast := &Broadcast{}
	cast.Type = "broadcast"
	cast.SetProductionMode(true)
	cast.Payload.Initial()
	cast.Payload.SetDisplayType("notification")
	return cast
}
