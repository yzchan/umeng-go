package android

import "github.com/yzchan/umeng-go/push/notification"

type Customizedcast struct {
	notification.Cast
	AliasType string  `json:"alias_type"`
	Alias     string  `json:"alias,omitempty"`
	FileId    string  `json:"file_id,omitempty"`
	Payload   Payload `json:"payload"`
	MiPush
}

func NewCustomizedcast() *Customizedcast {
	cast := &Customizedcast{}
	cast.Type = "customizedcast"
	cast.SetProductionMode(true)
	cast.Payload.Initial()
	cast.Payload.SetDisplayType("notification")
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
