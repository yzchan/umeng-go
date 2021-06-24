package ios

import (
	"encoding/json"
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

func (cast *Customizedcast) Send() (string, error) {
	return cast.BaseSend(cast)
}

func (cast *Customizedcast) String() string {
	marshaled, _ := json.MarshalIndent(cast, "", "    ")
	return string(marshaled)
}

func (cast *Customizedcast) SetTemplateName(name string) {
	cast.TemplateName = name
}

func (cast *Customizedcast) AddToTemplate(name string) (string, error) {
	cast.SetAppKey(cast.App.AppKey)
	cast.SetAlias("${alias}")
	cast.SetFileId("")
	cast.SetTemplateName(name)
	cast.SetTimestamp()
	return cast.App.AddTemplate(cast)
}
