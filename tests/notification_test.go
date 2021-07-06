package tests

import (
	"fmt"
	iosDemo "github.com/yzchan/umeng-go/demo/IOS"
	androidDemo "github.com/yzchan/umeng-go/demo/android"
	"github.com/yzchan/umeng-go/push"
	"testing"
)

var u *push.Umeng

// 使用本测试脚本之前需要通过config.go文件配置友盟基础参数，参见config.go.template
func init() {
	u = push.NewUmeng().
		InitAndroid(AndroidAppKey, AndroidMasterSecret).SetPackageName(AndroidPackageName).
		InitIOS(IOSAppKey, IOSMasterSecret)
	u.Debug(false)
}

func checkUmeng(t *testing.T) (ret bool) {
	ret = true
	if u.Android.AppKey == "" {
		t.Log("请设置环境变量[export ANDROID_APPKEY=]")
		ret = false
	}
	if u.Android.MasterSecret == "" {
		t.Log("请设置环境变量[export ANDROID_MASTER_SECRET=]")
		ret = false
	}
	if u.Android.PackageName == "" {
		t.Log("请设置环境变量[export ANDROID_PACKAGE_NAME=]")
		ret = false
	}
	if u.IOS.AppKey == "" {
		t.Log("请设置环境变量[export IOS_APPKEY=]")
		ret = false
	}
	if u.IOS.MasterSecret == "" {
		t.Log("请设置环境变量[export IOS_MASTER_SECRET=]")
		ret = false
	}
	if !ret {
		t.Fatal("初始化失败")
	}
	return ret
}

func TestAndroidUnicast(t *testing.T) {
	checkUmeng(t)

	t.Log("send android unicast...")
	msgId, err := androidDemo.SendUnicast(u, AndroidToken1)

	if err != nil {
		t.Log(err.Error())
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Logf("send success. msg_id=%s\n", msgId)
	t.Log("\x1b[32mtest ok \x1b[0m")
}

func TestAndroidListcast(t *testing.T) {
	checkUmeng(t)

	t.Log("send android listcast...")
	msgId, err := androidDemo.SendListcast(u, []string{
		AndroidToken1,
		AndroidToken2,
	})

	if err != nil {
		t.Log(err.Error())
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Logf("send success. msg_id=%s\n", msgId)
	t.Log("\x1b[32mtest ok \x1b[0m")
}

func TestAndroidBroadcast(t *testing.T) {
	checkUmeng(t)

	t.Log("send android broadcast...")
	taskId, err := androidDemo.SendBroadcast(u)

	if err != nil {
		t.Log(err.Error())
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Logf("send success. task_id=%s\n", taskId)
	t.Log("\x1b[32mtest ok \x1b[0m")
}

func TestAndroidGroupcast(t *testing.T) {
	checkUmeng(t)

	t.Log("send android groupcast...")
	taskId, err := androidDemo.SendGroupcast(u, "game3")

	if err != nil {
		t.Log(err.Error())
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Logf("send success. task_id=%s\n", taskId)
	t.Log("\x1b[32mtest ok \x1b[0m")
}

func TestAndroidFilecast(t *testing.T) {
	checkUmeng(t)

	t.Log("send android filecast...")

	t.Log(" -> upload first.")
	content := fmt.Sprintf("%s\n%s", AndroidToken1, AndroidToken2)
	fileId, err := u.Android.Upload(content)
	if err != nil {
		t.Log("upload error.", err.Error())
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Logf(" -> upload ok. file_id=%s\n", fileId)

	taskId, err := androidDemo.SendFilecast(u, fileId)

	if err != nil {
		t.Log(err.Error())
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Logf("send success. task_id=%s\n", taskId)
	t.Log("\x1b[32mtest ok \x1b[0m")
}

func TestAndroidCustomizedcastWithFileId(t *testing.T) {
	checkUmeng(t)

	t.Log("send android customizedcast with file_id...")

	t.Log(" -> upload first.")
	fileId, err := u.Android.Upload("2\n3\n4")
	if err != nil {
		t.Log("upload error.", err.Error())
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Logf(" -> upload ok. file_id=%s\n", fileId)

	taskId, err := androidDemo.SendCustomizedcastWithFileId(u, "kUMessageAliasTypeUserId", fileId)

	if err != nil {
		t.Log(err.Error())
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Logf("send success. task_id=%s\n", taskId)
	t.Log("\x1b[32mtest ok \x1b[0m")
}

func TestAndroidCustomizedcastWithAlias(t *testing.T) {
	checkUmeng(t)

	t.Log("send android customizedcast with alias...")

	msgId, err := androidDemo.SendCustomizedcastWithAlias(u, "kUMessageAliasTypeUserId", "2")

	if err != nil {
		t.Log(err.Error())
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Logf("send success. msg_id=%s\n", msgId)
	t.Log("\x1b[32mtest ok \x1b[0m")
}

func TestIOSUnicast(t *testing.T) {
	checkUmeng(t)

	t.Log("send ios unicast...")
	msgId, err := iosDemo.SendUnicast(u, IOSToken1)

	if err != nil {
		t.Log(err.Error())
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Logf("send success. msg_id=%s\n", msgId)
	t.Log("\x1b[32mtest ok \x1b[0m")
}

func TestIOSListcast(t *testing.T) {
	checkUmeng(t)

	t.Log("send ios listcast...")
	msgId, err := iosDemo.SendListcast(u, []string{
		IOSToken1,
		IOSToken2,
	})

	if err != nil {
		t.Log(err.Error())
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Logf("send success. msg_id=%s\n", msgId)
	t.Log("\x1b[32mtest ok \x1b[0m")
}

func TestIOSBroadcast(t *testing.T) {
	checkUmeng(t)

	t.Log("send ios broadcast...")
	taskId, err := iosDemo.SendBroadcast(u)

	if err != nil {
		t.Log(err.Error())
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Logf("send success. task_id=%s\n", taskId)
	t.Log("\x1b[32mtest ok \x1b[0m")
}

func TestIOSGroupcast(t *testing.T) {
	checkUmeng(t)

	t.Log("send ios groupcast...")
	taskId, err := iosDemo.SendGroupcast(u, "game3")

	if err != nil {
		t.Log(err.Error())
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Logf("send success. task_id=%s\n", taskId)
	t.Log("\x1b[32mtest ok \x1b[0m")
}

func TestIOSFilecast(t *testing.T) {
	checkUmeng(t)

	t.Log("send ios filecast...")

	t.Log(" -> upload first.")
	content := fmt.Sprintf("%s\n%s", IOSToken1, IOSToken2)
	fileId, err := u.IOS.Upload(content)
	if err != nil {
		t.Log("upload error.", err.Error())
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Logf(" -> upload ok. file_id=%s\n", fileId)

	taskId, err := iosDemo.SendFilecast(u, fileId)

	if err != nil {
		t.Log(err.Error())
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Logf("send success. task_id=%s\n", taskId)
	t.Log("\x1b[32mtest ok \x1b[0m")
}

func TestIOSCustomizedcastWithFileId(t *testing.T) {
	checkUmeng(t)

	t.Log("send ios customizedcast with file_id...")

	t.Log(" -> upload first.")
	fileId, err := u.IOS.Upload("2\n3\n4")
	if err != nil {
		t.Log("upload error.", err.Error())
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Logf(" -> upload ok. file_id=%s\n", fileId)

	taskId, err := iosDemo.SendCustomizedcastWithFileId(u, "kUMessageAliasTypeUserId", fileId)

	if err != nil {
		t.Log(err.Error())
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Logf("send success. task_id=%s\n", taskId)
	t.Log("\x1b[32mtest ok \x1b[0m")
}

func TestIOSCustomizedcastWithAlias(t *testing.T) {
	checkUmeng(t)

	t.Log("send ios customizedcast with alias...")

	msgId, err := iosDemo.SendCustomizedcastWithAlias(u, "kUMessageAliasTypeUserId", "2")

	if err != nil {
		t.Log(err.Error())
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Logf("send success. msg_id=%s\n", msgId)
	t.Log("\x1b[32mtest ok \x1b[0m")
}
