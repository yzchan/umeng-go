package android

import (
	"github.com/yzchan/umeng-go/push/notification"
)

func NewListcast() *notification.AndroidNotification {
	n := notification.NewAndroidNotification()
	n.SetNotificationType(notification.Listcast)
	n.Payload.Initial()
	return n
}
