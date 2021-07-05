package android

import (
	"github.com/yzchan/umeng-go/push/notification"
)

func NewBroadcast() *notification.IOSNotification {
	n := notification.NewIOSNotification()
	n.SetNotificationType(notification.Broadcast)
	n.Payload.Initial()
	return n
}
