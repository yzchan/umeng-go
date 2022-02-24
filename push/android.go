package push

import (
	"github.com/yzchan/umeng-go/v2/push/notification"
)

type AndroidRequest struct {
	notification.Notification
	Platform string                      `json:"-"`
	Payload  notification.AndroidPayload `json:"payload"`
	Channel  notification.Channel        `json:"channel_properties,omitempty"`
}

func (r *AndroidRequest) GetPlatform() string {
	return r.Platform
}

func (r *AndroidRequest) GetRequestUri() string {
	return Host + SendPath
}

func NewAndroidRequest(cast string) *AndroidRequest {
	n := &AndroidRequest{Platform: Android}
	n.SetNotificationType(cast)
	n.InitTimestamp()
	n.Payload.Initial()
	n.Policy = &notification.Policy{}
	return n
}

func NewAndroidUnicastRequest() *AndroidRequest {
	return NewAndroidRequest(notification.Unicast)
}

func NewAndroidListcastRequest() *AndroidRequest {
	return NewAndroidRequest(notification.Listcast)
}

func NewAndroidGroupcastRequest() *AndroidRequest {
	return NewAndroidRequest(notification.Groupcast)
}

func NewAndroidBroadcastRequest() *AndroidRequest {
	return NewAndroidRequest(notification.Broadcast)
}

func NewAndroidFilecastRequest() *AndroidRequest {
	return NewAndroidRequest(notification.Filecast)
}

func NewAndroidCustomizedcastRequest() *AndroidRequest {
	return NewAndroidRequest(notification.Customizedcast)
}
