package push

import (
	"encoding/json"
	"time"
)

type StatusReq struct {
	Appkey    string `json:"appkey"`
	Timestamp int64  `json:"timestamp"`
	TaskId    string `json:"task_id"`
}

type StatusResp struct {
	Ret  string     `json:"ret"`
	Data StatusData `json:"data"`
}

type StatusData struct {
	TaskId       string `json:"task_id"`
	OpenCount    int64  `json:"open_count"`
	SentCount    int64  `json:"sent_count"`
	DismissCount int64  `json:"dismiss_count"`
	Status       int64  `json:"status"`
	ErrorCode    int64  `json:"error_code"`
	ErrorMsg     string `json:"error_msg"`
}

func (u *Client) Status(taskId string) (ret StatusData, err error) {
	var result []byte
	data := StatusReq{u.Appkey, time.Now().Unix(), taskId}
	if result, err = u.Request(Host+StatusPath, data); err != nil {
		return
	}
	var s StatusResp
	if err = json.Unmarshal(result, &s); err != nil {
		return
	}
	ret = s.Data
	return
}
