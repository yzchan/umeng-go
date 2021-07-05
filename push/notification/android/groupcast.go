package android

import (
	"github.com/yzchan/umeng-go/push/notification"
)

func NewGroupcast() *notification.AndroidNotification {
	n := notification.NewAndroidNotification()
	n.SetNotificationType(notification.Groupcast)
	n.Payload.Initial()
	return n
}
