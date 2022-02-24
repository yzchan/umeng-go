package request

import (
	"github.com/yzchan/umeng-go/push"
	"github.com/yzchan/umeng-go/push/notification"
)

type AndroidBaseRequest struct {
	notification.Notification
	Platform string                      `json:"-"`
	Payload  notification.AndroidPayload `json:"payload"`
	Channel  notification.Channel        `json:"channel_properties,omitempty"`
}

func (r *AndroidBaseRequest) GetPlatform() string {
	return r.Platform
}

func NewAndroidBaseRequest(cast string) *AndroidBaseRequest {
	n := &AndroidBaseRequest{Platform: push.Android}
	n.SetNotificationType(cast)
	n.InitTimestamp()
	n.Payload.Initial()
	n.Policy = &notification.Policy{}
	return n
}

func NewAndroidUnicastRequest() *AndroidBaseRequest {
	return NewAndroidBaseRequest(notification.Unicast)
}

func NewAndroidListcastRequest() *AndroidBaseRequest {
	return NewAndroidBaseRequest(notification.Listcast)
}

func NewAndroidGroupcastRequest() *AndroidBaseRequest {
	return NewAndroidBaseRequest(notification.Groupcast)
}

func NewAndroidBroadcastRequest() *AndroidBaseRequest {
	return NewAndroidBaseRequest(notification.Broadcast)
}

func NewAndroidFilecastRequest() *AndroidBaseRequest {
	return NewAndroidBaseRequest(notification.Filecast)
}

func NewAndroidCustomizedcastRequest() *AndroidBaseRequest {
	return NewAndroidBaseRequest(notification.Customizedcast)
}
