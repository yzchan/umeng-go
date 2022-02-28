package notification

import "time"

const TimeFormat = "2006-01-02 15:04:05"

type Policy struct {
	StartTime      string `json:"start_time,omitempty"`
	ExpireTime     string `json:"expire_time,omitempty"`
	MaxSendNum     int    `json:"max_send_num,omitempty"`
	OutBizNo       string `json:"out_biz_no,omitempty"`
	ApnsCollapseId string `json:"apns_collapse_id,omitempty"` // 只对ios生效
}

func (p *Policy) SetStartAt(time time.Time) *Policy {
	p.StartTime = time.Format(TimeFormat)
	return p
}

func (p *Policy) SetStartTime(time string) *Policy {
	p.StartTime = time
	return p
}

func (p *Policy) SetExpireAt(time time.Time) *Policy {
	p.ExpireTime = time.Format(TimeFormat)
	return p
}

func (p *Policy) SetExpireTime(time string) *Policy {
	p.ExpireTime = time
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

func (p *Policy) SetApnsCollapseId(id string) *Policy {
	p.ApnsCollapseId = id
	return p
}
