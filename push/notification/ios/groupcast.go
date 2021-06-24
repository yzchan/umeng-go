package ios

import (
	"encoding/json"
	"github.com/yzchan/umeng-go/push/notification"
	"time"
)

type Groupcast struct {
	notification.Cast
	DeviceTokens string      `json:"device_tokens"`
	Payload      Payload     `json:"payload"`
	Filter       interface{} `json:"filter"`
	Policy       Policy      `json:"policy,omitempty"`
}

func NewGroupcast() *Groupcast {
	cast := &Groupcast{}
	cast.Type = "groupcast"
	cast.Payload = make(Payload)
	cast.Payload.Initial()
	cast.Timestamp = time.Now().Unix()
	return cast
}

func (cast *Groupcast) SetFilter(condition string) *Groupcast {
	var v interface{}
	_ = json.Unmarshal([]byte(condition), &v)
	cast.Filter = v
	return cast
}

func (cast *Groupcast) Send() (string, error) {
	return cast.BaseSend(cast)
}

func (cast *Groupcast) String() string {
	marshaled, _ := json.MarshalIndent(cast, "", "    ")
	return string(marshaled)
}
