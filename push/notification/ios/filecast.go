package ios

import (
	"github.com/yzchan/umeng-go/push/notification"
	"time"
)

type Filecast struct {
	notification.Cast
	FileId  string  `json:"file_id"`
	Payload Payload `json:"payload"`
	Policy  Policy  `json:"policy,omitempty"`
}

func NewFilecast() *Filecast {
	cast := &Filecast{}
	cast.Type = "filecast"
	cast.Payload = make(Payload)
	cast.Payload.Initial()
	cast.Timestamp = time.Now().Unix()
	return cast
}

func (f *Filecast) SetFileId(fileId string) *Filecast {
	f.FileId = fileId
	return f
}
