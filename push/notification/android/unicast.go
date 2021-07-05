package android

import (
	"github.com/yzchan/umeng-go/push/notification"
)

func NewUnicast() *notification.AndroidNotification {
	n := notification.NewAndroidNotification()
	n.SetNotificationType(notification.Unicast)
	n.Payload.Initial()
	return n
}
