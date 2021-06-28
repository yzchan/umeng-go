package android

type Channel struct {
	ChannelActivity    string `json:"channel_activity,omitempty"`
	XiaomiChannelId    string `json:"xiaomi_channel_id,omitempty"`
	VivoClassification string `json:"vivo_classification,omitempty"`
	OppoChannelId      string `json:"oppo_channel_id,omitempty"`
}

func (c *Channel) SetChannelActivity(activity string) *Channel {
	c.ChannelActivity = activity + ".activity.MiPushActivity"
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
