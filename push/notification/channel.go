package notification

type Channel struct {
	ChannelActivity    string `json:"channel_activity,omitempty"`
	XiaomiChannelId    string `json:"xiaomi_channel_id,omitempty"`
	VivoClassification string `json:"vivo_classification,omitempty"`
	OppoChannelId      string `json:"oppo_channel_id,omitempty"`
	MainActivity       string `json:"main_activity,omitempty"`
}

func (c *Channel) SetChannelActivity(activity string) *Channel {
	c.ChannelActivity = activity
	return c
}

func (c *Channel) SetMainActivity(activity string) *Channel {
	c.MainActivity = activity
	return c
}

func (c *Channel) SetXiaomiChannelId(id string) *Channel {
	c.XiaomiChannelId = id
	return c
}

func (c *Channel) SetVivoClassification(val string) *Channel {
	c.VivoClassification = val
	return c
}

func (c *Channel) SetOppoChannelId(id string) *Channel {
	c.OppoChannelId = id
	return c
}
