package android

type MiPush struct {
	MiPush     string `json:"mipush,omitempty"`
	MiActivity string `json:"mi_activity,omitempty"`
}

// Deprecated: use Channel.SetChannelActivity instead.
func (m *MiPush) SetPackageName(activity string) *MiPush {
	m.MiPush = "true"
	m.MiActivity = activity + ".activity.MiPushActivity"
	return m
}
