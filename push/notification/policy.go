package notification

type Policy struct {
	StartTime      string `json:"start_time,omitempty"`
	ExpireTime     string `json:"expire_time,omitempty"`
	MaxSendNum     int    `json:"max_send_num,omitempty"`
	OutBizNo       string `json:"out_biz_no,omitempty"`
}
