package request

import (
	"github.com/yzchan/umeng-go/push"
	"github.com/yzchan/umeng-go/push/notification"
)

type IOSBaseRequest struct {
	notification.Notification
	Platform string                  `json:"-"`
	Payload  notification.IOSPayload `json:"payload"`
}

func (r *IOSBaseRequest) GetPlatform() string {
	return r.Platform
}

func NewIOSBaseRequest(cast string) *IOSBaseRequest {
	n := &IOSBaseRequest{Platform: push.IOS}
	n.Payload = make(notification.IOSPayload)
	n.Payload.Initial()
	n.Policy = &notification.Policy{}
	n.SetNotificationType(cast)
	n.InitTimestamp()
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
