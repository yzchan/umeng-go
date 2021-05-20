package ios

type Payload map[string]interface{}

func (p *Payload) Initial() {
	(*p)["aps"] = &APNs{
		ContentAvailable: 0,
	}
}

func (p *Payload) GetAPNs() *APNs {
	return (*p)["aps"].(*APNs)
}

func (p *Payload) AddExtra(key string, val string) *Payload {
	if key == "aps" { // 防止自定义aps覆盖
		return p
	}
	(*p)[key] = val
	return p
}

type APNs struct {
	Alert struct { // 当content-available=1时(静默推送)，可选; 否则必填
		Title    string `json:"title"`    // 标题
		Subtitle string `json:"subtitle"` // 副标题
		Body     string `json:"body"`     // 文本内容
	} `json:"alert"`
	Badge            string `json:"badge,omitempty"`             // 可选
	Sound            string `json:"sound,omitempty"`             // 可选
	ContentAvailable int    `json:"content-available,omitempty"` // 可选，代表静默推送
	Category         string `json:"category,omitempty"`          // 可选，注意: ios8才支持该字段
}

func (a *APNs) SetAlert(title string, subTitle string, body string) *APNs {
	a.Alert.Title = title
	a.Alert.Subtitle = subTitle
	a.Alert.Body = body
	return a
}

func (a *APNs) SetBadge(badge string) *APNs {
	a.Badge = badge
	return a
}

func (a *APNs) SetSound(sound string) *APNs {
	a.Sound = sound
	return a
}

func (a *APNs) SetContentAvailable(val int) *APNs {
	a.ContentAvailable = val
	return a
}

func (a *APNs) SetCategory(category string) *APNs {
	a.Category = category
	return a
}
