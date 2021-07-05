package notification

import (
	"encoding/json"
	"github.com/yzchan/umeng-go/push/notification/ios"
)

type IOSNotification struct {
	Notification
	Payload ios.Payload `json:"payload"`
}

func NewIOSNotification() *IOSNotification {
	cast := &IOSNotification{}
	cast.Payload = make(ios.Payload)
	cast.Payload.Initial()
	cast.Policy = &Policy{}
	return cast
}

func (n *IOSNotification) Send() (string, error) {
	return n.BaseSend(n)
}

func (n *IOSNotification) String() string {
	marshaled, _ := json.Marshal(n)
	return string(marshaled)
}
