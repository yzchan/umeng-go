package android

import (
	"github.com/yzchan/umeng-go/push/notification"
)

func NewCustomizedcast() *notification.IOSNotification {
	n := notification.NewIOSNotification()
	n.SetNotificationType(notification.Customizedcast)
	n.Payload.Initial()
	return n
}
