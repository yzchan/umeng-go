package android

type Payload struct {
	DisplayType string            `json:"display_type"`
	Body        map[string]string `json:"body"`
	Extra       map[string]string `json:"extra"`
}

func (p *Payload) Initial() *Payload {
	p.Body = make(map[string]string)
	p.Extra = make(map[string]string)
	return p
}

func (p *Payload) SetDisplayType(displayType string) *Payload {
	p.DisplayType = displayType
	return p
}

func (p *Payload) SetBody(key string, val string) *Payload {
	p.Body[key] = val
	return p
}

func (p *Payload) SetBodies(data map[string]string) *Payload {
	for key, val := range data {
		p.Body[key] = val
	}
	return p
}

func (p *Payload) SetExtra(key string, val string) *Payload {
	p.Extra[key] = val
	return p
}

func (p *Payload) SetExtras(data map[string]string) *Payload {
	for key, val := range data {
		p.Extra[key] = val
	}
	return p
}
