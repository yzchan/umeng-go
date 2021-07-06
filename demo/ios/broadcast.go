package ios

import (
	"github.com/yzchan/umeng-go/push"
	"github.com/yzchan/umeng-go/push/notification/ios"
	"strconv"
	"time"
)

func SendBroadcast(u *push.Umeng) (string, error) {
	n := ios.NewBroadcast()
	n.SetDescription("broadcast-test")
	n.Payload.GetAPNs().SetTitle("broadcast-title").SetSubTitle("broadcast-text").SetBody("broadcast-body")
	n.Payload.GetAPNs().SetImg("https://xxx.com/xxx") // 设置图片
	//n.SetImage("https://xxx.com/xxx") // 设置图片
	n.Payload.AddExtra("extra1", "1").AddExtra("extra2", "2")
	n.Policy.SetOutBizNo(strconv.Itoa(int(time.Now().Unix()))) // 设置外部id，主要用于防止重复推送，只对任务类消息生效
	//n.Policy.SetStartTime(time.Unix(time.Now().Unix()+3600, 0)) // 延迟1小时
	//n.Policy.SetStartTimeStr("2022-12-12 12:12:12") // 直接设置有效的时间字符串
	n.BindApp(u.IOS)
	return n.Send()
}
