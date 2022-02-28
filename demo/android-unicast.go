package main

import (
	"fmt"
	"github.com/yzchan/umeng-go/v2/push"
)

func main() {
	u := push.NewUmeng()
	u.InitAndroid("*****", "*****")
	u.InitIOS("*****", "*****")
	u.Debug(true) // 开启调试模式会输出http请求和响应内容

	// 查询推送任务结果
	fmt.Println(u.Android.Status("us*****"))

	// 撤销推送
	fmt.Println(u.GetClient(push.IOS).Cancel("um*****"))

	// 新建一条android单播消息
	request := push.NewAndroidUnicastRequest()
	request.SetDescription("unicast-test")   // 消息描述，友盟后台显示
	request.SetDeviceToken("<device_token>") // 需要推送的设备
	request.Payload.Body.SetTitle("unicast-title").SetText("unicast-text")
	request.Payload.Extra.AddKV("extra1", "1").AddKV("extra2", "2")
	request.Channel.SetChannelActivity("xxx") // 离线通道需要用

	result, err := u.Send(request)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
