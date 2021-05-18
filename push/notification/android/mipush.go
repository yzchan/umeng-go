package android

type MiPush struct {
	MiPush     string `json:"mipush"`
	MiActivity string `json:"mi_activity"`
}

func (m *MiPush) SetPackageName(activity string) *MiPush {
	m.MiPush = "true"
	m.MiActivity = activity + ".activity.MiPushActivity"
	return m
}
