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
	n := &AndroidNotification{}
	n.Payload.Initial()
	n.Policy = &Policy{}
	return n
}

func (n *AndroidNotification) Send() (string, error) {
	n.Channel.SetChannelActivity(n.App.PackageName)
	return n.BaseSend(n)
}

func (n *AndroidNotification) AddToTemplate(name string) (string, error) {
	n.Channel.SetChannelActivity(n.App.PackageName)
	return n.InitTemplate(name)
}

func (n *AndroidNotification) String() string {
	marshaled, _ := json.Marshal(n)
	return string(marshaled)
}

func (n *AndroidNotification) GetNotification() *Notification {
	return &n.Notification
}

func (n *AndroidNotification) SetTitle(title string) Notificationer {
	n.Payload.Body.SetTitle(title)
	return n
}

func (n *AndroidNotification) SetText(text string) Notificationer {
	n.Payload.Body.SetText(text)
	return n
}

func (n *AndroidNotification) SetExtras(extras map[string]string) Notificationer {
	for k, v := range extras {
		n.Payload.AddExtra(k, v)
	}
	return n
}

func (n *AndroidNotification) SetImage(img string) Notificationer {
	n.Payload.Body.SetImg(img)
	return n
}

func (n *AndroidNotification) SetSilent() Notificationer {
	n.Payload.SetDisplayType("message")
	return n
}
