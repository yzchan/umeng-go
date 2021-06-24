package android

import (
	"encoding/json"
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

func (cast *Customizedcast) SetAliasType(aliasType string) *Customizedcast {
	cast.AliasType = aliasType
	return cast
}

func (cast *Customizedcast) SetAlias(alias string) *Customizedcast {
	if cast.TemplateName != "" {
		return cast
	}
	cast.Alias = alias
	return cast
}

func (cast *Customizedcast) SetFileId(fileId string) *Customizedcast {
	if cast.TemplateName != "" {
		return cast
	}
	cast.FileId = fileId
	return cast
}

func (cast *Customizedcast) SetAsTemplate(name string) {
	cast.TemplateName = name
	cast.Alias = "${alias}"
	cast.FileId = ""
}

func (cast *Customizedcast) Send() (string, error) {
	cast.SetPackageName(cast.App.PackageName)
	return cast.BaseSend(cast)
}

func (cast *Customizedcast) String() string {
	marshaled, _ := json.MarshalIndent(cast, "", "    ")
	return string(marshaled)
}
