package ios

import (
	"github.com/yzchan/umeng-go/push/notification"
)

func NewUnicast() *notification.IOSNotification {
	n := notification.NewIOSNotification()
	n.SetNotificationType(notification.Unicast)
	n.Payload.Initial()
	return n
}
