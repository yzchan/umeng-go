package android

import (
	"encoding/json"
	"github.com/yzchan/umeng-go/push/notification"
)

type Unicast struct {
	notification.Cast
	DeviceTokens string  `json:"device_tokens"`
	Payload      Payload `json:"payload"`
	Policy       Policy  `json:"policy,omitempty"`
	MiPush
	TemplateName string `json:"template_name,omitempty"`
}

func NewUnicast() *Unicast {
	cast := &Unicast{}
	cast.Type = "unicast"
	cast.Payload.Initial()
	return cast
}

func (cast *Unicast) SetDeviceToken(token string) *Unicast {
	if cast.TemplateName != "" { // 模板消息 禁止设置DeviceTokens
		return cast
	}
	cast.DeviceTokens = token
	return cast
}

func (cast *Unicast) Send() (string, error) {
	cast.SetPackageName(cast.App.PackageName)
	return cast.BaseSend(cast)
}

func (cast *Unicast) String() string {
	marshaled, _ := json.MarshalIndent(cast, "", "    ")
	return string(marshaled)
}

func (cast *Unicast) SetAsTemplate(name string) {
	cast.TemplateName = name
	cast.DeviceTokens = "${device_tokens}"
}
