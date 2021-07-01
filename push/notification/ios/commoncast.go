package ios

import (
	"encoding/json"
	"github.com/yzchan/umeng-go/push/notification"
)

type Commoncast struct {
	notification.Cast
	Payload      Payload     `json:"payload"`
	Policy       Policy      `json:"policy,omitempty"`
	AliasType    string      `json:"alias_type,omitempty"`
	Alias        string      `json:"alias,omitempty"`
	FileId       string      `json:"file_id,omitempty"`
	DeviceTokens string      `json:"device_tokens,omitempty"`
	Filter       interface{} `json:"filter,omitempty"`
}

func NewCommoncast() *Commoncast {
	cast := &Commoncast{}
	cast.Payload = make(Payload)
	cast.Payload.Initial()
	return cast
}

func (cast *Commoncast) SetAliasType(aliasType string) *Commoncast {
	cast.AliasType = aliasType
	return cast
}

func (cast *Commoncast) SetAlias(alias string) *Commoncast {
	cast.Alias = alias
	return cast
}

func (cast *Commoncast) SetFileId(fileId string) *Commoncast {
	cast.FileId = fileId
	return cast
}

func (cast *Commoncast) SetFilter(condition string) *Commoncast {
	var v interface{}
	_ = json.Unmarshal([]byte(condition), &v)
	cast.Filter = v
	return cast
}

func (cast *Commoncast) SetDeviceTokens(tokens string) *Commoncast {
	cast.DeviceTokens = tokens
	return cast
}

func (cast *Commoncast) Send() (string, error) {
	return cast.BaseSend(cast)
}

func (cast *Commoncast) String() string {
	marshaled, _ := json.Marshal(cast)
	return string(marshaled)
}
