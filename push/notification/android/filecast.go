package android

import (
	"encoding/json"
	"github.com/yzchan/umeng-go/push/notification"
)

type Filecast struct {
	notification.Cast
	FileId  string  `json:"file_id"`
	Payload Payload `json:"payload"`
	Policy  Policy  `json:"policy,omitempty"`
	MiPush
}

func NewFilecast() *Filecast {
	cast := &Filecast{}
	cast.Type = "filecast"
	cast.Payload.Initial()
	return cast
}

func (cast *Filecast) SetFileId(fileId string) *Filecast {
	cast.FileId = fileId
	return cast
}

func (cast *Filecast) Send() (string, error) {
	cast.SetPackageName(cast.App.PackageName)
	return cast.BaseSend(cast)
}

func (cast *Filecast) String() string {
	marshaled, _ := json.MarshalIndent(cast, "", "    ")
	return string(marshaled)
}
