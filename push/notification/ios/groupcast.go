package ios

import (
	"encoding/json"
	"github.com/yzchan/umeng-go/push/notification"
)

type Groupcast struct {
	notification.Cast
	DeviceTokens string      `json:"device_tokens"`
	Payload      Payload     `json:"payload"`
	Filter       interface{} `json:"filter"`
	Policy       Policy      `json:"policy"`
}

func NewGroupcast() *Groupcast {
	cast := &Groupcast{}
	cast.Type = "groupcast"
	cast.SetProductionMode(true)
	cast.Payload = make(Payload)
	cast.Payload.Initial()
	return cast
}

func (g *Groupcast) SetFilter(condition string) *Groupcast {
	var v interface{}
	json.Unmarshal([]byte(condition), &v)
	g.Filter = v
	return g
}
