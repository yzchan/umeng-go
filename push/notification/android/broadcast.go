package android

import (
	"github.com/yzchan/umeng-go/push/notification"
	"time"
)

type Broadcast struct {
	notification.Cast
	Payload Payload `json:"payload"`
	Policy  Policy  `json:"policy,omitempty"`
	MiPush
}

func NewBroadcast() *Broadcast {
	cast := &Broadcast{}
	cast.Type = "broadcast"
	cast.Payload.Initial()
	cast.Timestamp = time.Now().Unix()
	return cast
}
