package android

import (
	"github.com/yzchan/umeng-go/push/notification"
	"time"
)

type Customizedcast struct {
	notification.Cast
	AliasType string  `json:"alias_type"`
	Alias     string  `json:"alias,omitempty"`
	FileId    string  `json:"file_id,omitempty"`
	Payload   Payload `json:"payload"`
	Policy    Policy  `json:"policy,omitempty"`
	MiPush
}

func NewCustomizedcast() *Customizedcast {
	cast := &Customizedcast{}
	cast.Type = "customizedcast"
	cast.Payload.Initial()
	cast.Timestamp = time.Now().Unix()
	return cast
}

func (c *Customizedcast) SetAliasType(aliasType string) *Customizedcast {
	c.AliasType = aliasType
	return c
}

func (c *Customizedcast) SetAlias(alias string) *Customizedcast {
	c.Alias = alias
	return c
}

func (c *Customizedcast) SetFileId(fileId string) *Customizedcast {
	c.FileId = fileId
	return c
}
