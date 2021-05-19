package ios

import (
	"github.com/yzchan/umeng-go/push/notification"
)

type Filecast struct {
	notification.Cast
	FileId  string  `json:"file_id"`
	Payload Payload `json:"payload"`
}

func NewFilecast() *Filecast {
	cast := &Filecast{}
	cast.Type = "filecast"
	cast.SetProductionMode(true)
	cast.Payload = make(Payload)
	cast.Payload.Initial()
	return cast
}

func (f *Filecast) SetFileId(fileId string) *Filecast {
	f.FileId = fileId
	return f
}
