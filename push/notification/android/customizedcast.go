package android

import (
	"github.com/yzchan/umeng-go/push/notification"
)

type Customizedcast struct {
	notification.Cast
	AliasType string  `json:"alias_type"`
	Alias     string  `json:"alias,omitempty"`
	FileId    string  `json:"file_id,omitempty"`
	Payload   Payload `json:"payload"`
	Policy    Policy  `json:"policy,omitempty"`
	MiPush
	TemplateName string `json:"template_name,omitempty"`
}

func NewCustomizedcast() *Customizedcast {
	cast := &Customizedcast{}
	cast.Type = "customizedcast"
	cast.Payload.Initial()
	return cast
}

func (c *Customizedcast) SetAliasType(aliasType string) *Customizedcast {
	c.AliasType = aliasType
	return c
}

func (c *Customizedcast) SetAlias(alias string) *Customizedcast {
	if c.TemplateName != "" {
		return c
	}
	c.Alias = alias
	return c
}

func (c *Customizedcast) SetFileId(fileId string) *Customizedcast {
	if c.TemplateName != "" {
		return c
	}
	c.FileId = fileId
	return c
}

func (c *Customizedcast) SetAsTemplate(name string) {
	c.TemplateName = name
	c.Alias = "${alias}"
	c.FileId = ""
}

func (c *Customizedcast) Send() (string, error) {
	c.SetPackageName(c.App.PackageName)
	return c.BaseSend(c)
}
