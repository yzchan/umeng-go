package android

import (
	"github.com/yzchan/umeng-go/push"
	"github.com/yzchan/umeng-go/push/notification/android"
	"strconv"
	"time"
)

func SendBroadcast(u *push.Umeng) (string, error) {
	n := android.NewBroadcast()
	n.SetDescription("broadcast-test")
	n.Payload.Body.SetTitle("broadcast-title").SetText("broadcast-text")
	n.Payload.Body.SetImg("https://xxx.com/xxx") // 设置图片
	//n.SetImage("https://xxx.com/xxx") // 设置图片
	n.Payload.Extra.AddKV("extra1", "1").AddKV("extra2", "2")
	n.Policy.SetOutBizNo(strconv.Itoa(int(time.Now().Unix()))) // 设置外部id，主要用于防止重复推送，只对任务类消息生效
	//n.Policy.SetStartTime(time.Unix(time.Now().Unix()+3600, 0)) // 延迟1小时
	//n.Policy.SetStartTimeStr("2022-12-12 12:12:12") // 直接设置有效的时间字符串
	n.BindApp(u.Android)
	return n.Send()
}
