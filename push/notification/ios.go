package notification

import (
	"encoding/json"
	. "github.com/yzchan/umeng-go/push/notification/payload"
)

type IOSNotification struct {
	Notification
	Payload IOSPayload `json:"payload"`
}

func NewIOSNotification() *IOSNotification {
	cast := &IOSNotification{}
	cast.Payload = make(IOSPayload)
	cast.Payload.Initial()
	cast.Policy = &Policy{}
	return cast
}

func (n *IOSNotification) Send() (string, error) {
	return n.BaseSend(n)
}

func (n *IOSNotification) AddToTemplate(name string) (string, error) {
	return n.InitTemplate(name)
}

func (n *IOSNotification) String() string {
	marshaled, _ := json.Marshal(n)
	return string(marshaled)
}
