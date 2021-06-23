package android

import (
	"github.com/yzchan/umeng-go/push/notification"
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
	return cast
}

func (b *Broadcast) Send() (string, error) {
	b.SetPackageName(b.App.PackageName)
	return b.BaseSend(b)
}
