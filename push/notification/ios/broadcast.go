package ios

import (
	"encoding/json"
	"github.com/yzchan/umeng-go/push/notification"
	"time"
)

type Broadcast struct {
	notification.Cast
	Payload Payload `json:"payload"`
	Policy  Policy  `json:"policy,omitempty"`
}

func NewBroadcast() *Broadcast {
	cast := &Broadcast{}
	cast.Type = "broadcast"
	cast.Payload = make(Payload)
	cast.Payload.Initial()
	cast.Timestamp = time.Now().Unix()
	return cast
}

func (cast *Broadcast) Send() (string, error) {
	return cast.BaseSend(cast)
}

func (cast *Broadcast) String() string {
	marshaled, _ := json.MarshalIndent(cast, "", "    ")
	return string(marshaled)
}
