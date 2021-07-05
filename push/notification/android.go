package notification

import (
	"encoding/json"
	. "github.com/yzchan/umeng-go/push/notification/payload"
)

type AndroidNotification struct {
	Notification
	Payload AndroidPayload `json:"payload"`
	Channel Channel        `json:"channel_properties,omitempty"`
}

func NewAndroidNotification() *AndroidNotification {
	cast := &AndroidNotification{}
	cast.Payload.Initial()
	cast.Policy = &Policy{}
	return cast
}

func (n *AndroidNotification) Send() (string, error) {
	n.Channel.SetChannelActivity(n.App.PackageName)
	return n.BaseSend(n)
}

func (n *AndroidNotification) String() string {
	marshaled, _ := json.Marshal(n)
	return string(marshaled)
}
