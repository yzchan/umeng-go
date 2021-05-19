package ios

type APNs struct {
	Alert struct { // 当content-available=1时(静默推送)，可选; 否则必填
		Title    string `json:"title"`    // 标题
		Subtitle string `json:"subtitle"` // 副标题
		Body     string `json:"body"`     // 文本内容
	} `json:"alert"`
	Badge            string `json:"badge,omitempty"`             // 可选
	Sound            string `json:"sound,omitempty"`             // 可选
	ContentAvailable int64  `json:"content-available,omitempty"` // 可选，代表静默推送
	Category         string `json:"category,omitempty"`          // 可选，注意: ios8才支持该字段
}

func (p *Payload) Initial() {
	(*p)["aps"] = &APNs{
		ContentAvailable: 0,
	}
}

func (a *APNs) SetAlert(title string, subTitle string, body string) {
	a.Alert.Title = title
	a.Alert.Subtitle = subTitle
	a.Alert.Body = body
}

type Payload map[string]interface{}

//func (p *Payload) SetAlert(title string, subTitle string, body string) {
//	(*p)["aps"].(*APNs).SetAlert(title, subTitle, body)
//}

func (p *Payload) SetExtra(key string, val string) {
	if key == "aps" { // 防止自定义aps覆盖
		return
	}
	(*p)[key] = val
}

func (p *Payload) SetExtras(data map[string]string) {
	for key, val := range data {
		(*p).SetExtra(key, val)
	}
}
