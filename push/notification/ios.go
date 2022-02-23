package notification

import (
	"encoding/json"
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

//func (n *IOSNotification) Send() (string, error) {
//	return n.BaseSend(n)
//}
//
//func (n *IOSNotification) AddToTemplate(name string) (string, error) {
//	return n.InitTemplate(name)
//}

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

type IOSPayload map[string]interface{}

func (p *IOSPayload) Initial() {
	(*p)["aps"] = &APNs{
		ContentAvailable: 0,
	}
}

func (p *IOSPayload) GetAPNs() *APNs {
	return (*p)["aps"].(*APNs)
}

func (p *IOSPayload) AddExtra(key string, val string) *IOSPayload {
	if key == "aps" { // 防止自定义aps覆盖
		return p
	}
	(*p)[key] = val
	return p
}

type Alert struct {
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
	Body     string `json:"body"`
}

type APNs struct {
	Alert            *Alert `json:"alert"`
	Badge            string `json:"badge,omitempty"`
	Sound            string `json:"sound"`
	ContentAvailable int    `json:"content-available,omitempty"`
	MutableContent   int    `json:"mutable-content,omitempty"`
	Category         string `json:"category,omitempty"`
	QFAttach         string `json:"QFAttach,omitempty"`
}

func (a *APNs) SetTitle(title string) *APNs {
	if a.Alert == nil {
		a.Alert = &Alert{}
	}
	a.Alert.Title = title
	return a
}

func (a *APNs) SetSubTitle(subTitle string) *APNs {
	if a.Alert == nil {
		a.Alert = &Alert{}
	}
	a.Alert.SubTitle = subTitle
	return a
}

func (a *APNs) SetBody(body string) *APNs {
	if a.Alert == nil {
		a.Alert = &Alert{}
	}
	a.Alert.Body = body
	return a
}

func (a *APNs) SetBadge(badge string) *APNs {
	a.Badge = badge
	return a
}

func (a *APNs) SetSound(sound string) *APNs {
	a.Sound = sound
	return a
}

func (a *APNs) SetContentAvailable(val int) *APNs {
	a.ContentAvailable = val
	return a
}

func (a *APNs) SetMutableContent(val int) *APNs {
	a.MutableContent = val
	return a
}

func (a *APNs) SetCategory(category string) *APNs {
	a.Category = category
	return a
}

func (a *APNs) SetImg(img string) *APNs {
	a.QFAttach = img
	return a
}
