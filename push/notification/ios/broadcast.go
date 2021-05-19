package ios

import "github.com/yzchan/umeng-go/push/notification"

type Broadcast struct {
	notification.Cast
	Payload      Payload `json:"payload"`
}

func NewBroadcast() *Broadcast {
	cast := &Broadcast{}
	cast.Type = "broadcast"
	cast.SetProductionMode(true)
	cast.Payload = make(Payload)
	cast.Payload.Initial()
	return cast
}
