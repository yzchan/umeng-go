package notification

type Policy struct {
	StartTime  string `json:"start_time,omitempty"`
	ExpireTime string `json:"expire_time,omitempty"`
	MaxSendNum int    `json:"max_send_num,omitempty"`
	OutBizNo   string `json:"out_biz_no,omitempty"`
}

func (p *Policy) SetStartTime(time string) *Policy {
	p.StartTime = time
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
