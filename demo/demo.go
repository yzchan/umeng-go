package demo

import (
	"fmt"
	"github.com/yzchan/umeng-go/push"
	"github.com/yzchan/umeng-go/push/notification/android"
	"github.com/yzchan/umeng-go/push/notification/ios"
)

const (
	iosToken1     = "token111"
	iosToken2     = "token222"
	androidToken1 = "token333"
	androidToken2 = "token444"
	packageName   = "{package name}"
)

var u *push.Umeng

func init() {
	u = push.NewUmeng().
		InitAndroid("aaa", "bbb").SetPackageName(packageName).
		InitIOS("ccc", "ddd")

}

func SendIOSUnicast() {
	cast := ios.NewUnicast()
	cast.SetDescription("单播测试")
	cast.Payload.GetAPNs().SetAlert("单播测试-title", "单播测试-subtitle", "单播测试-body")
	cast.Payload.AddExtra("show_alert", "1").AddExtra("type", "2")
	cast.SetDeviceToken(iosToken1)

	u.IOS.PreView(cast)
	ret, err := u.IOS.Send(cast)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}

func SendIOSListcast() {
	cast := ios.NewListcast()
	cast.SetDescription("列播测试")
	cast.Payload.GetAPNs().SetAlert("列播测试-title", "列播测试-subtitle", "列播测试-body")
	cast.Payload.AddExtra("show_alert", "1").AddExtra("type", "2")
	cast.SetDeviceTokens([]string{iosToken1, iosToken2})

	u.IOS.PreView(cast)
	ret, err := u.IOS.Send(cast)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}

func SendIOSBroadcast() {
	cast := ios.NewBroadcast()
	cast.SetDescription("广播测试")
	cast.Payload.GetAPNs().SetAlert("广播测试-title", "广播测试-subtitle", "广播测试-body")
	cast.Payload.AddExtra("show_alert", "1").AddExtra("type", "2")

	u.IOS.PreView(cast)
	ret, err := u.IOS.Send(cast)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}

func SendIOSGroupcast() {
	cast := ios.NewGroupcast()
	cast.SetDescription("组播测试")
	cast.Payload.GetAPNs().SetAlert("组播测试-title", "组播测试-subtitle", "组播测试-body")
	cast.Payload.AddExtra("show_alert", "1").AddExtra("type", "2")
	cast.SetFilter(`{
	"where":
	{
	"and": [{"tag": "group-1"}]
	}
	}`)

	u.IOS.PreView(cast)
	ret, err := u.IOS.Send(cast)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}

func SendIOSFilecast() {
	//fileId, err := u.IOS.Upload(iosToken1)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(fileId)
	fileId := "*****"

	cast := ios.NewFilecast()
	cast.SetDescription("文件播测试")
	cast.Payload.GetAPNs().SetAlert("文件播测试-title", "文件播测试-subtitle", "文件播测试-body")
	cast.Payload.AddExtra("show_alert", "1").AddExtra("type", "2")
	cast.SetFileId(fileId)

	u.IOS.PreView(cast)
	ret, err := u.IOS.Send(cast)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}

func SendIOSCustomizedcast() {
	cast := ios.NewCustomizedcast()
	cast.SetDescription("自定义播测试")
	cast.Payload.GetAPNs().SetAlert("自定义播测试-title", "自定义播测试-subtitle", "自定义播测试-body")
	cast.Payload.AddExtra("show_alert", "1").AddExtra("type", "2")
	cast.SetAliasType("AliasType")
	cast.SetAlias("1234")

	u.IOS.PreView(cast)
	ret, err := u.IOS.Send(cast)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}

func SendAndroidUnicast() {
	cast := android.NewUnicast()
	cast.SetPackageName(packageName)
	cast.SetDescription("单播测试")
	cast.Payload.Body.SetTitle("单播测试-title").SetText("单播测试-text")
	cast.Payload.Extra.AddKV("show_alert", "1").AddKV("type", "2")
	cast.SetDeviceToken(androidToken1)

	u.Android.PreView(cast)
	ret, err := u.Android.Send(cast)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}

func SendAndroidListcast() {
	cast := android.NewListcast()
	cast.SetPackageName(packageName)
	cast.SetDescription("列播测试")
	cast.Payload.Body.SetTitle("列播测试-title").SetText("列播测试-text")
	cast.Payload.Extra.AddKV("show_alert", "1").AddKV("type", "2")
	cast.SetDeviceTokens([]string{androidToken1, androidToken2})

	u.Android.PreView(cast)
	ret, err := u.Android.Send(cast)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}

func SendAndroidBroadcast() {
	cast := android.NewBroadcast()
	cast.SetPackageName(packageName)
	cast.SetDescription("广播测试")
	cast.Payload.Body.SetTitle("广播测试-title").SetText("广播测试-text")
	cast.Payload.Extra.AddKV("show_alert", "1").AddKV("type", "2")
	//cast.Policy.SetStartTime(time.Unix(time.Now().Unix()+3600, 0)) // 延后一小时

	u.Android.PreView(cast)
	ret, err := u.Android.Send(cast)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}

func SendAndroidGroupcast() {
	cast := android.NewGroupcast()
	cast.SetPackageName(packageName)
	cast.SetDescription("组播测试")
	cast.Payload.Body.SetTitle("组播测试-title").SetText("组播测试-text")
	cast.Payload.Extra.AddKV("show_alert", "1").AddKV("type", "2")
	cast.SetFilter(`{
	"where":
	{
	"and": [{"tag": "group-1"}]
	}
	}`)

	u.Android.PreView(cast)
	ret, err := u.Android.Send(cast)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}

func SendAndroidFilecast() {
	//fileId, err := u.Android.Upload(androidToken1)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(fileId)
	fileId := "*****"

	cast := android.NewFilecast()
	cast.SetPackageName(packageName)
	cast.SetDescription("文件播测试")
	cast.Payload.Body.SetTitle("文件播测试-title").SetText("文件播测试-text")
	cast.Payload.Extra.AddKV("show_alert", "1").AddKV("type", "2")
	cast.SetFileId(fileId)

	u.Android.PreView(cast)
	ret, err := u.Android.Send(cast)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}

func SendAndroidCustomizedcast() {
	cast := android.NewCustomizedcast()
	cast.SetPackageName(packageName)
	cast.SetDescription("自定义播测试")
	cast.Payload.Body.SetTitle("自定义播测试-title").SetText("自定义播测试-text")
	cast.Payload.Extra.AddKV("show_alert", "1").AddKV("type", "2")
	cast.SetAliasType("AliasType")
	cast.SetAlias("4321")

	u.Android.PreView(cast)
	ret, err := u.Android.Send(cast)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}
