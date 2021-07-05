package android

import (
	"github.com/yzchan/umeng-go/push/notification"
)

func NewBroadcast() *notification.AndroidNotification {
	n := notification.NewAndroidNotification()
	n.SetNotificationType(notification.Broadcast)
	n.Payload.Initial()
	return n
}
