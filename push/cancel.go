package push

import (
	"encoding/json"
	"time"
)

type CancelResp struct {
	Ret  string `json:"ret"`
	Data struct {
		TaskId    string `json:"task_id"`
		ErrorCode string `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
	} `json:"data"`
}

func (u *Client) Cancel(taskId string) (ret CancelResp, err error) {

	data := struct {
		Appkey    string `json:"appkey"`
		Timestamp int64  `json:"timestamp"`
		TaskId    string `json:"task_id"`
	}{
		Appkey:    u.Appkey,
		Timestamp: time.Now().Unix(),
		TaskId:    taskId,
	}

	resp, err := u.Request(Host+CancelPath, data)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return
	}
	return
}
