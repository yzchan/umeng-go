package android

import (
	"github.com/yzchan/umeng-go/push"
	"github.com/yzchan/umeng-go/push/notification/android"
)

// SendListcast 列播推送必须指定device_tokens。可以使用 SetDeviceTokenList 或者 SetDeviceTokens方法设置
func SendListcast(u *push.Umeng, deviceTokens []string) (string, error) {
	n := android.NewListcast()
	n.SetDescription("listcast-test")
	n.Payload.Body.SetTitle("listcast-title").SetText("listcast-text")
	n.Payload.Extra.AddKV("extra1", "1").AddKV("extra2", "2")
	n.SetDeviceTokenList(deviceTokens)
	//n.SetDeviceTokens("token1,token2") // 也可以自行拼接DeviceToken字符串，以英文逗号分隔
	n.BindApp(u.Android)
	return n.Send()
}
