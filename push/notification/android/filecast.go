package android

import (
	"github.com/yzchan/umeng-go/push/notification"
)

func NewFilecast() *notification.AndroidNotification {
	n := notification.NewAndroidNotification()
	n.SetNotificationType(notification.Filecast)
	n.Payload.Initial()
	return n
}
