package request

import (
	"github.com/yzchan/umeng-go/push/notification"
)

type IOSBaseRequest struct {
	notification.Notification
	Payload notification.IOSPayload `json:"payload"`
}

func NewIOSBaseRequest(cast string) *IOSBaseRequest {
	n := &IOSBaseRequest{}
	n.SetNotificationType(cast)
	n.InitTimestamp()
	n.Payload.Initial()
	return n
}

func NewIOSUnicastRequest() *IOSBaseRequest {
	return NewIOSBaseRequest(notification.Unicast)
}

func NewIOSListcastRequest() *IOSBaseRequest {
	return NewIOSBaseRequest(notification.Listcast)
}

func NewIOSGroupcastRequest() *IOSBaseRequest {
	return NewIOSBaseRequest(notification.Groupcast)
}

func NewIOSBroadcastRequest() *IOSBaseRequest {
	return NewIOSBaseRequest(notification.Broadcast)
}

func NewIOSFilecastRequest() *IOSBaseRequest {
	return NewIOSBaseRequest(notification.Filecast)
}

func NewIOSCustomizedcastRequest() *IOSBaseRequest {
	return NewIOSBaseRequest(notification.Customizedcast)
}
