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

func (g *Groupcast) SetFilter(condition string) *Groupcast {
	var v interface{}
	json.Unmarshal([]byte(condition), &v)
	g.Filter = v
	return g
}
