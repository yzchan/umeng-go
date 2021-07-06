package ios

import (
	"github.com/yzchan/umeng-go/push"
	"github.com/yzchan/umeng-go/push/notification/ios"
)

// SendUnicast 单播推送必须指定device_tokens。可以使用SetDeviceToken方法设置
func SendUnicast(u *push.Umeng, deviceToken string) (string, error) {
	n := ios.NewUnicast()
	n.SetDescription("unicast-test")
	n.Payload.GetAPNs().SetTitle("unicast-title").SetSubTitle("unicast-subtitle").SetBody("unicast-body")
	n.Payload.AddExtra("extra1", "1").AddExtra("extra2", "2")
	n.SetDeviceToken(deviceToken)
	n.BindApp(u.IOS)
	return n.Send()
}
