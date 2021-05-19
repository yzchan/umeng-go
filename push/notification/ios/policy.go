package ios

import "github.com/yzchan/umeng-go/push/notification"

type Policy struct {
	notification.Policy
	ApnsCollapseId string `json:"apns_collapse_id,omitempty"`
}

func (p *Policy) SetApnsCollapseId(id string) *Policy {
	p.ApnsCollapseId = id
	return p
}
