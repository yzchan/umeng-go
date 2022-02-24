package push

import (
	"github.com/yzchan/umeng-go/v2/push/notification"
)

type IOSRequest struct {
	notification.Notification
	Platform string                  `json:"-"`
	Payload  notification.IOSPayload `json:"payload"`
}

func (r *IOSRequest) GetPlatform() string {
	return r.Platform
}

func (r *IOSRequest) GetRequestUri() string {
	return Host + SendPath
}

func NewIOSRequest(cast string) *IOSRequest {
	n := &IOSRequest{Platform: IOS}
	n.Payload = make(notification.IOSPayload)
	n.Payload.Initial()
	n.Policy = &notification.Policy{}
	n.SetNotificationType(cast)
	n.InitTimestamp()
	return n
}

func NewIOSUnicastRequest() *IOSRequest {
	return NewIOSRequest(notification.Unicast)
}

func NewIOSListcastRequest() *IOSRequest {
	return NewIOSRequest(notification.Listcast)
}

func NewIOSGroupcastRequest() *IOSRequest {
	return NewIOSRequest(notification.Groupcast)
}

func NewIOSBroadcastRequest() *IOSRequest {
	return NewIOSRequest(notification.Broadcast)
}

func NewIOSFilecastRequest() *IOSRequest {
	return NewIOSRequest(notification.Filecast)
}

func NewIOSCustomizedcastRequest() *IOSRequest {
	return NewIOSRequest(notification.Customizedcast)
}
