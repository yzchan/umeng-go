package android

import (
	"github.com/yzchan/umeng-go/push/notification"
)

type Filecast struct {
	notification.Cast
	FileId  string  `json:"file_id"`
	Payload Payload `json:"payload"`
	MiPush
}

func NewFilecast() *Filecast {
	cast := &Filecast{}
	cast.Type = "filecast"
	cast.SetProductionMode(true)
	cast.Payload.Initial()
	cast.Payload.SetDisplayType("notification")
	return cast
}

func (f *Filecast) SetFileId(fileId string) *Filecast {
	f.FileId = fileId
	return f
}
