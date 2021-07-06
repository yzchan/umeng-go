package ios

import (
	"github.com/yzchan/umeng-go/push/notification"
)

func NewGroupcast() *notification.IOSNotification {
	n := notification.NewIOSNotification()
	n.SetNotificationType(notification.Groupcast)
	n.Payload.Initial()
	return n
}
