package android

import (
	"github.com/yzchan/umeng-go/push/notification"
)

func NewCustomizedcast() *notification.AndroidNotification {
	n := notification.NewAndroidNotification()
	n.SetNotificationType(notification.Customizedcast)
	n.Payload.Initial()
	return n
}
