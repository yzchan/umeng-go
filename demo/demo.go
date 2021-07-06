package demo

import (
	"fmt"
	"github.com/yzchan/umeng-go/push"
	"github.com/yzchan/umeng-go/push/notification"
)

var u *push.Umeng

func init() {
	u = push.NewUmeng().
		InitAndroid("aaa", "bbb").SetPackageName("{package name}").
		InitIOS("ccc", "ddd")
	u.Debug(true)
}

// 通用方法
func otherDemo() {
	// 查询任务结果，只能查询任务类消息task_id
	fmt.Println(u.Android.Status("usw68v916214979512****"))
	// 安卓厂商数据透出
	fmt.Println(u.Android.Channel("usw68v916214979512****"))
	// 安卓厂商额度查询
	fmt.Println(u.Android.Quota())
	// 上传文件，结合文件播和自定义文件播使用
	fmt.Println(u.Android.Upload("content..."))
	// 撤销推送，只能撤销任务类消息
	fmt.Println(u.Android.Cancel("umrr338162149966Demo**"))
}

// 通用notification，使用本sdk封装接口时使用，平台和消息类型可以作为参数从外部传入，避免代码中编写过多if-else
func notificationDemo() {
	n := notification.NewUmengNotification(push.Android)
	n.GetNotification().SetNotificationType(notification.Unicast)
	n.GetNotification().SetFileId("title")
	n.GetNotification().SetAlias("alias")
	n.GetNotification().SetAliasType("aliasType")
	n.GetNotification().SetDescription("desc")
	n.GetNotification().SetDeviceTokens("xxx")
	n.GetNotification().SetFilter("...")
	n.GetNotification().BindApp(u.GetApp(push.Android))

	n.GetNotification().Policy.SetOutBizNo("UniqueId")
	n.GetNotification().Policy.SetStartTimeStr("2022-12-12 12:12:12")

	n.SetTitle("title")
	n.SetText("text")
	n.SetExtras(map[string]string{
		"extra1": "1",
		"extra2": "2",
	})

	fmt.Println(n.Send())
}

// 标签管理
func tagsDemo() {
	token1 := "xxxxxxxxxxx"
	fmt.Println(u.Android.AddTag(token1, "group-1"))
	fmt.Println("SetTag:\t", u.Android.SetTags(token1, []string{"a", "b"}))
	fmt.Println(u.Android.ListTags(token1))
	fmt.Println("AddTag:\t", u.Android.AddTags(token1, []string{"c", "d"}))
	fmt.Println(u.Android.ListTags(token1))
	fmt.Println("DeleteTag:\t", u.Android.DeleteTags(token1, []string{"a", "c", "d"}))
	fmt.Println(u.Android.ListTags(token1))
	fmt.Println("ClearTag:\t", u.Android.ClearTags(token1))
	fmt.Println(u.Android.ListTags(token1))
}

// TODO
func templateDemo() {

}
