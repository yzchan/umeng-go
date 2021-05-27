package push

import (
	"encoding/json"
	"time"
)

type ChannelReq struct {
	Appkey    string `json:"appkey"`
	Timestamp int64  `json:"timestamp"`
	TaskId    string `json:"task_id"`
}

type ChannelResp struct {
	Ret  string      `json:"ret"`
	Data ChannelData `json:"data"`
}

type ChannelData struct {
	Stats []struct {
		Channel            string        `json:"channel"`
		ChannelArriveCount int           `json:"channel_arrive_count"`
		ChannelClick       int           `json:"channel_click"`
		ChannelSentCount   int           `json:"channel_sent_count"`
		Errors             []interface{} `json:"errors"`
	} `json:"stats"`
}

func (u *Client) Channel(taskId string) (ret ChannelData, err error) {
	var result []byte
	data := ChannelReq{u.Appkey, time.Now().Unix(), taskId}
	if result, err = u.Request(Host+ChannelPath, data); err != nil {
		return
	}
	var c ChannelResp
	if err = json.Unmarshal(result, &c); err != nil {
		return
	}
	ret = c.Data
	return
}