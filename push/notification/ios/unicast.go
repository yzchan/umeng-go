package ios

import (
	"github.com/yzchan/umeng-go/push/notification"
	"time"
)

type Unicast struct {
	notification.Cast
	DeviceTokens string  `json:"device_tokens"`
	Payload      Payload `json:"payload"`
	Policy       Policy  `json:"policy,omitempty"`
	TemplateName string  `json:"template_name,omitempty"`
}

func NewUnicast() *Unicast {
	cast := &Unicast{}
	cast.Type = "unicast"
	cast.Payload = make(Payload)
	cast.Payload.Initial()
	cast.Timestamp = time.Now().Unix()
	return cast
}

func (u *Unicast) SetDeviceToken(token string) *Unicast {
	if u.TemplateName != "" { // 模板消息 禁止设置DeviceTokens
		return u
	}
	u.DeviceTokens = token
	return u
}

func (u *Unicast) Send() (string, error) {
	return u.BaseSend(u)
}

func (u *Unicast) SetAsTemplate(name string) {
	u.TemplateName = name
	u.DeviceTokens = "${device_tokens}"
}
