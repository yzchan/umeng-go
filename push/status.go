package push

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type StatusResp struct {
	Ret  string `json:"ret"`
	Data struct {
		TaskId       string `json:"task_id"`
		OpenCount    int64  `json:"open_count"`
		SentCount    int64  `json:"sent_count"`
		DismissCount int64  `json:"dismiss_count"`
		Status       int64  `json:"status"`
		ErrorCode    int64  `json:"error_code"`
		ErrorMsg     string `json:"error_msg"`
	} `json:"data"`
}

func (u *Client) Status(taskId string) (ret StatusResp, err error) {

	data := struct {
		Appkey    string `json:"appkey"`
		Timestamp int64  `json:"timestamp"`
		TaskId    string `json:"task_id"`
	}{
		Appkey:    u.Appkey,
		Timestamp: time.Now().Unix(),
		TaskId:    taskId,
	}

	resp, err := u.Request(Host+StatusPath, data)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	retStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(retStr, &ret)
	if err != nil {
		return
	}
	return
}
