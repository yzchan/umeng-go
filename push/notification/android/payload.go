package android

// Payload
// @link https://developer.umeng.com/docs/67966/detail/68343#h2-u8C03u7528u53C2u65703
type Payload struct {
	DisplayType string `json:"display_type"`    // 必填，消息类型: notification(通知)、message(消息)
	Body        Body   `json:"body"`            // 必填，消息体
	Extra       Extra  `json:"extra,omitempty"` // 可选，用户自定义key-value。
}

func (p *Payload) Initial() *Payload {
	p.Extra = make(Extra)
	p.DisplayType = "notification"
	return p
}

func (p *Payload) AddExtra(key string, val string) *Payload {
	p.Extra.AddKV(key, val)
	return p
}

// SetDisplayType 设置通知类型
// 可选值notification(通知)、message(消息)
// 当通知类型为message时，Body消息体中只有Custom字段生效
func (p *Payload) SetDisplayType(displayType string) *Payload {
	if displayType == "message" {
		p.DisplayType = displayType
	} else {
		p.DisplayType = "notification"
	}
	return p
}

// Extra 用户自定义key-value
type Extra map[string]string

func (e *Extra) AddKV(key string, val string) *Extra {
	(*e)[key] = val
	return e
}

// Body
// 当display_type=message时，body的内容只需填写custom字段
type Body struct { // 必填，消息体
	Title       string `json:"title"`                  // 必填，通知标题
	Text        string `json:"text"`                   // 必填，通知文字描述
	Icon        string `json:"icon,omitempty"`         // 可选，状态栏图标ID，R.drawable.[smallIcon]，如果没有，默认使用应用图标
	LargeIcon   string `json:"largeIcon,omitempty"`    // 可选，通知栏拉开后左侧图标ID，R.drawable.[largeIcon]
	Img         string `json:"img,omitempty"`          // 可选，通知栏大图标的URL链接。该字段的优先级大于largeIcon 只支持华为，链接需要以https开头
	ExpandImage string `json:"expand_image,omitempty"` // 可选，消息下方展示大图，支持自有通道消息展示 厂商通道展示大图目前仅支持小米
	Sound       string `json:"sound,omitempty"`        // 可选，通知声音，R.raw.[sound] 如果该字段为空，采用SDK默认的声音
	BuilderId   int    `json:"builder_id,omitempty"`   // 可选，默认为0，用于标识该通知采用的样式。使用该参数时 开发者必须在SDK里面实现自定义通知栏样式
	PlayVibrate string `json:"play_vibrate,omitempty"` // 可选，收到通知是否震动，默认为"true"
	PlayLights  string `json:"play_lights,omitempty"`  // 可选，收到通知是否闪灯，默认为"true"
	PlaySound   string `json:"play_sound,omitempty"`   // 可选，收到通知是否发出声音，默认为"true"
	AfterOpen   string `json:"after_open,omitempty"`   // 可选，点击"通知"的后续行为(默认为打开app):，可选值: go_app/go_url/go_activity/go_custom
	Url         string `json:"url,omitempty"`          // 配合after_open=go_url时使用。 通知栏点击后跳转的URL
	Activity    string `json:"activity,omitempty"`     // 配合after_open=go_activity时使用。当after_open=go_activity时，必填。 通知栏点击后打开的Activity
	Custom      string `json:"custom,omitempty"`       // 配合after_open=go_custom时使用。或配合display_type=message时使用
}

// SetTitle 设置通知标题
func (b *Body) SetTitle(title string) *Body {
	b.Title = title
	return b
}

// SetText 设置通知文字描述
func (b *Body) SetText(text string) *Body {
	b.Text = text
	return b
}

func (b *Body) SetIcon(icon string) *Body {
	b.Icon = icon
	return b
}

func (b *Body) SetLargeIcon(icon string) *Body {
	b.LargeIcon = icon
	return b
}

func (b *Body) SetImg(img string) *Body {
	b.Img = img
	return b
}

func (b *Body) SetExpandImage(img string) *Body {
	b.ExpandImage = img
	return b
}

func (b *Body) SetSound(sound string) *Body {
	b.Sound = sound
	return b
}

func (b *Body) SetBuilderId(id int) *Body {
	b.BuilderId = id
	return b
}

// SetPlayVibrate 收到通知是否震动，默认为"true"
func (b *Body) SetPlayVibrate(play bool) *Body {
	if !play {
		b.PlayVibrate = "false"
	}
	return b
}

// SetPlayLights 收到通知是否闪灯，默认为"true"
func (b *Body) SetPlayLights(play bool) *Body {
	if !play {
		b.PlayLights = "false"
	}
	return b
}

// SetPlaySound 收到通知是否发出声音，默认为"true"
func (b *Body) SetPlaySound(play bool) *Body {
	if !play {
		b.PlaySound = "false"
	}
	return b
}

func (b *Body) SetAfterOpen(action string) *Body {
	b.AfterOpen = action
	return b
}

func (b *Body) SetUrl(url string) *Body {
	b.Url = url
	return b
}

func (b *Body) SetActivity(text string) *Body {
	b.Activity = text
	return b
}

func (b *Body) SetCustom(custom string) *Body {
	b.Custom = custom
	return b
}
