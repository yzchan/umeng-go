package android

// Payload
// @link https://developer.umeng.com/docs/67966/detail/68343#h2-u8C03u7528u53C2u65703
type Payload struct {
	DisplayType string `json:"display_type"`
	Body        struct { // 必填，消息体
		// 当display_type=message时，body的内容只需填写custom字段
		Custom string `json:"custom,omitempty"` // 当display_type=message时,必填。当display_type=notification且after_open=go_custom时，必填。用户自定义内容，可以为字符串或者JSON格式。
		// 当display_type=notification时，body包含如下参数:
		Title string `json:"title"` // 必填，通知标题
		Text        string `json:"text"` // Text 必填，通知文字描述
		Icon        string `json:"icon,omitempty"`         // 可选，状态栏图标ID，R.drawable.[smallIcon]，如果没有，默认使用应用图标
		LargeIcon   string `json:"largeIcon,omitempty"`    // 可选，通知栏拉开后左侧图标ID，R.drawable.[largeIcon]
		Img         string `json:"img,omitempty"`          // 可选，通知栏大图标的URL链接。该字段的优先级大于largeIcon 只支持华为，链接需要以https开头
		ExpandImage string `json:"expand_image,omitempty"` // 可选，消息下方展示大图，支持自有通道消息展示 厂商通道展示大图目前仅支持小米
		Sound       string `json:"sound,omitempty"`        // 可选，通知声音，R.raw.[sound] 如果该字段为空，采用SDK默认的声音
		BuilderId   int    `json:"builder_id,omitempty"`   // 可选，默认为0，用于标识该通知采用的样式。使用该参数时 开发者必须在SDK里面实现自定义通知栏样式
		PlayVibrate string `json:"play_vibrate,omitempty"` // 可选，收到通知是否震动，默认为"true"
		PlayLights  string `json:"play_lights,omitempty"`  // 可选，收到通知是否闪灯，默认为"true"
		PlaySound   string `json:"play_sound,omitempty"`   // 可选，收到通知是否发出声音，默认为"true"
		AfterOpen   string `json:"after_open,omitempty"`   // 可选，默认为"go_app"，值可以为: "go_app":打开应用 "go_url":跳转到URL "go_activity":打开特定的activity "go_custom":用户自定义内容
		Url         string `json:"url,omitempty"`          // 当after_open=go_url时，必填。 通知栏点击后跳转的URL，要求以http或者https开头
		Activity    string `json:"activity,omitempty"`     //当after_open=go_activity时，必填。 通知栏点击后打开的Activity
	} `json:"body"`
	Extra map[string]string `json:"extra"`
}

func (p *Payload) Initial() *Payload {
	p.Extra = make(map[string]string)
	return p
}

func (p *Payload) SetDisplayType(displayType string) *Payload {
	p.DisplayType = displayType
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
