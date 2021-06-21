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
	Status       int64  `json:"status"`        // 消息状态：0排队中 1发送中 2发送完成 3发送失败 4消息被撤销 5消息过期 6筛选结果为空 7定时任务尚未开始处理
	OpenCount    int64  `json:"open_count"`    // android/ios共有：打开数
	SentCount    int64  `json:"sent_count"`    // android/ios共有：消息收到数(android) APNs返回SUCCESS的设备数(ios)
	DismissCount int64  `json:"dismiss_count"` // android特有：忽略数
	TotalCount   int64  `json:"total_count"`   // ios特有：投递APNs设备数
}

func (p *Platform) Status(taskId string) (ret StatusData, err error) {
	var result []byte
	data := StatusReq{p.Appkey, time.Now().Unix(), taskId}
	if result, err = p.Request(Host+StatusPath, data); err != nil {
		return
	}
	var s StatusResp
	if err = json.Unmarshal(result, &s); err != nil {
		return
	}
	ret = s.Data
	return
}
