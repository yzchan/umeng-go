package notification

import (
	"encoding/json"
	. "github.com/yzchan/umeng-go/push/notification/payload"
	"strconv"
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

func (n *IOSNotification) GetNotification() *Notification {
	return &n.Notification
}

func (n *IOSNotification) SetTitle(title string) Notificationer {
	n.Payload.GetAPNs().SetTitle(title)
	return n
}

func (n *IOSNotification) SetText(text string) Notificationer {
	n.Payload.GetAPNs().SetBody(text)
	return n
}

func (n *IOSNotification) SetExtras(extras map[string]string) Notificationer {
	for k, v := range extras {
		n.Payload.AddExtra(k, v)
	}
	return n
}

func (n *IOSNotification) SetImage(img string) Notificationer {
	n.Payload.GetAPNs().SetImg(img)
	return n
}

func (n *IOSNotification) SetSilent() Notificationer {
	n.Payload.GetAPNs().Alert = nil
	n.Payload.GetAPNs().SetSound("")
	n.Payload.GetAPNs().SetContentAvailable(1)
	return n
}

func (n *IOSNotification) SetBadge(badge string) Notificationer {
	b, err := strconv.Atoi(badge)
	if err != nil {
		return n
	}
	n.Payload.GetAPNs().Badge = b
	return n
}
