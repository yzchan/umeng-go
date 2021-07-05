package android

import (
	"github.com/yzchan/umeng-go/push/notification"
)

func NewListcast() *notification.IOSNotification {
	n := notification.NewIOSNotification()
	n.SetNotificationType(notification.Listcast)
	n.Payload.Initial()
	return n
}
