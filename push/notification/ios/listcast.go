package ios

import (
	"encoding/json"
	"github.com/yzchan/umeng-go/push/notification"
	"strings"
)

type Listcast struct {
	notification.Cast
	DeviceTokens string  `json:"device_tokens"`
	Payload      Payload `json:"payload"`
	Policy       Policy  `json:"policy,omitempty"`
}

func NewListcast() *Listcast {
	cast := &Listcast{}
	cast.Type = "listcast"
	cast.Payload = make(Payload)
	cast.Payload.Initial()
	return cast
}

func (cast *Listcast) SetDeviceTokens(tokens []string) *Listcast {
	if len(tokens) > 500 {
		tokens = tokens[:500]
	}
	cast.DeviceTokens = strings.Join(tokens, ",")
	return cast
}

func (cast *Listcast) Send() (string, error) {
	return cast.BaseSend(cast)
}

func (cast *Listcast) String() string {
	marshaled, _ := json.MarshalIndent(cast, "", "    ")
	return string(marshaled)
}
