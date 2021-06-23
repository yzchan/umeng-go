package ios

import (
	"github.com/yzchan/umeng-go/push/notification"
	"time"
)

type Customizedcast struct {
	notification.Cast
	AliasType    string  `json:"alias_type"`        // 必选，别名类型 Alias 和 FileId 必选其一
	Alias        string  `json:"alias,omitempty"`   // 可选，别名，多个别名用,隔开
	FileId       string  `json:"file_id,omitempty"` // 可选，将别名上传之后的文件Id
	Payload      Payload `json:"payload"`
	Policy       Policy  `json:"policy,omitempty"`
	TemplateName string  `json:"template_name,omitempty"`
}

func NewCustomizedcast() *Customizedcast {
	cast := &Customizedcast{}
	cast.Type = "customizedcast"
	cast.Payload = make(Payload)
	cast.Payload.Initial()
	cast.Timestamp = time.Now().Unix()
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
	return c.BaseSend(c)
}
