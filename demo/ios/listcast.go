package ios

import (
	"github.com/yzchan/umeng-go/push"
	"github.com/yzchan/umeng-go/push/notification/ios"
)

// SendListcast 列播推送必须指定device_tokens。可以使用 SetDeviceTokenList 或者 SetDeviceTokens方法设置
func SendListcast(u *push.Umeng, deviceTokens []string) (string, error) {
	n := ios.NewListcast()
	n.SetDescription("listcast-test")
	n.Payload.GetAPNs().SetTitle("listcast-title").SetSubTitle("listcast-subtitle").SetBody("listcast-body")
	n.Payload.AddExtra("extra1", "1").AddExtra("extra2", "2")
	n.SetDeviceTokenList(deviceTokens)
	//n.SetDeviceTokens("token1,token2") // 也可以自行拼接DeviceToken字符串，以英文逗号分隔
	n.BindApp(u.IOS)
	return n.Send()
}
