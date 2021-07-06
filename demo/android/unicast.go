package android

import (
	"github.com/yzchan/umeng-go/push"
	"github.com/yzchan/umeng-go/push/notification/android"
)

// SendUnicast 单播推送必须指定device_tokens。可以使用SetDeviceToken方法设置
func SendUnicast(u *push.Umeng, deviceToken string) (string, error) {
	n := android.NewUnicast()
	n.SetDescription("unicast-test")
	n.Payload.Body.SetTitle("unicast-title").SetText("unicast-text")
	n.Payload.Extra.AddKV("extra1", "1").AddKV("extra2", "2")
	n.SetDeviceToken(deviceToken)
	n.BindApp(u.Android)
	return n.Send()
}
