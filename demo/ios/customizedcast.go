package ios

import (
	"github.com/yzchan/umeng-go/push"
	"github.com/yzchan/umeng-go/push/notification/ios"
	"strconv"
	"time"
)

// SendCustomizedcastWithFileId 自定义文件播属于任务类消息，返回task_id
func SendCustomizedcastWithFileId(u *push.Umeng, aliasType string, fileId string) (string, error) {
	n := ios.NewCustomizedcast()
	n.SetDescription("customizedcast-file-test")
	n.Payload.GetAPNs().SetTitle("customizedcast-file-title")
	n.Payload.GetAPNs().SetSubTitle("customizedcast-file-subtitle")
	n.Payload.GetAPNs().SetBody("customizedcast-file-body")
	n.Payload.AddExtra("extra1", "1").AddExtra("extra2", "2")
	n.Policy.SetOutBizNo("customizedcast-" + strconv.Itoa(int(time.Now().Unix()))) // 设置外部id，主要用于防止重复推送，只对任务类消息生效
	n.SetAliasType(aliasType)
	n.SetFileId(fileId)
	n.BindApp(u.IOS)
	return n.Send()
}

// SendCustomizedcastWithAlias 自定义alias播属于单播类消息，返回msg_id
func SendCustomizedcastWithAlias(u *push.Umeng, aliasType string, alias string) (string, error) {
	n := ios.NewCustomizedcast()
	n.SetDescription("customizedcast-alias-test")
	n.Payload.GetAPNs().SetTitle("customizedcast-alias-title")
	n.Payload.GetAPNs().SetSubTitle("customizedcast-alias-subtitle")
	n.Payload.GetAPNs().SetBody("customizedcast-alias-body")
	n.Payload.AddExtra("extra1", "1").AddExtra("extra2", "2")
	n.SetAliasType(aliasType)
	n.SetAlias(alias)
	n.BindApp(u.IOS)
	return n.Send()
}
