package notification

import "time"

const TimeFormat = "2006-01-02 15:04:05"

type Policy struct {
	StartTime  string `json:"start_time,omitempty"`
	ExpireTime string `json:"expire_time,omitempty"`
	MaxSendNum int    `json:"max_send_num,omitempty"`
	OutBizNo   string `json:"out_biz_no,omitempty"`
}

func (p *Policy) SetStartTime(time time.Time) *Policy {
	p.StartTime = time.Format(TimeFormat)
	return p
}

func (p *Policy) SetExpireTime(time time.Time) *Policy {
	p.ExpireTime = time.Format(TimeFormat)
	return p
}

func (p *Policy) SetMaxSendNum(num int) *Policy {
	p.MaxSendNum = num
	return p
}

func (p *Policy) SetOutBizNo(id string) *Policy {
	p.OutBizNo = id
	return p
}
