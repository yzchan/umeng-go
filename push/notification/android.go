package notification

type AndroidPayload struct {
	DisplayType string `json:"display_type"`
	Body        Body   `json:"body"`
	Extra       Extra  `json:"extra,omitempty"`
}

func (p *AndroidPayload) Initial() *AndroidPayload {
	p.Extra = make(Extra)
	p.DisplayType = "notification"
	return p
}

func (p *AndroidPayload) AddExtra(key string, val string) *AndroidPayload {
	p.Extra.AddKV(key, val)
	return p
}

func (p *AndroidPayload) SetDisplayType(displayType string) *AndroidPayload {
	if displayType == "message" {
		p.DisplayType = displayType
	} else {
		p.DisplayType = "notification"
	}
	return p
}

type Extra map[string]string

func (e *Extra) AddKV(key string, val string) *Extra {
	(*e)[key] = val
	return e
}

type Body struct {
	Title       string `json:"title"`
	Text        string `json:"text"`
	Ticker      string `json:"ticker,omitempty"`
	Icon        string `json:"icon,omitempty"`
	LargeIcon   string `json:"largeIcon,omitempty"`
	Img         string `json:"img,omitempty"`
	ExpandImage string `json:"expand_image,omitempty"`
	Sound       string `json:"sound,omitempty"`
	BuilderId   int    `json:"builder_id,omitempty"`
	PlayVibrate string `json:"play_vibrate,omitempty"`
	PlayLights  string `json:"play_lights,omitempty"`
	PlaySound   string `json:"play_sound,omitempty"`
	AfterOpen   string `json:"after_open,omitempty"`
	Url         string `json:"url,omitempty"`
	Activity    string `json:"activity,omitempty"`
	Custom      string `json:"custom"`
}

func (b *Body) SetTitle(title string) *Body {
	b.Title = title
	return b
}

func (b *Body) SetTicker(ticker string) *Body {
	b.Ticker = ticker
	return b
}

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

func (b *Body) SetPlayVibrate(play bool) *Body {
	if !play {
		b.PlayVibrate = "false"
	}
	return b
}

func (b *Body) SetPlayLights(play bool) *Body {
	if !play {
		b.PlayLights = "false"
	}
	return b
}

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
