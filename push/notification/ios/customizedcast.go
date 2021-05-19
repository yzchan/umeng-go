package ios

import (
	"github.com/yzchan/umeng-go/push/notification"
)

type Customizedcast struct {
	notification.Cast
	AliasType string  `json:"alias_type"`        // 必选，别名类型 Alias 和 FileId 必选其一
	Alias     string  `json:"alias,omitempty"`   // 可选，别名，多个别名用,隔开
	FileId    string  `json:"file_id,omitempty"` // 可选，将别名上传之后的文件Id
	Payload   Payload `json:"payload"`
}

func NewCustomizedcast() *Customizedcast {
	cast := &Customizedcast{}
	cast.Type = "customizedcast"
	cast.SetProductionMode(true)
	cast.Payload = make(Payload)
	cast.Payload.Initial()
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
