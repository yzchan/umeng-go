package notification

type CastResp struct {
	Ret  string `json:"ret"`
	Data struct {
		MsgId  string `json:"msg_id"`
		TaskId string `json:"task_id"`
	} `json:"data"`
}
