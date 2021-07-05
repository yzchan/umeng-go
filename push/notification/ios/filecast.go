package android

import (
	"github.com/yzchan/umeng-go/push/notification"
)

func NewFilecast() *notification.IOSNotification {
	n := notification.NewIOSNotification()
	n.SetNotificationType(notification.Filecast)
	n.Payload.Initial()
	return n
}
