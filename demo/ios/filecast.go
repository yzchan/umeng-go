package ios

import (
	"github.com/yzchan/umeng-go/push"
	"github.com/yzchan/umeng-go/push/notification/ios"
	"strconv"
	"time"
)

// SendFilecast 文件播推送需要先上传文件得到file_id。同样内容的file_id可以缓存使用（友盟建议缓存不超过3个月）。
func SendFilecast(u *push.Umeng, fileId string) (string, error) {
	n := ios.NewFilecast()
	n.SetDescription("filecast-test")
	n.Payload.GetAPNs().SetTitle("filecast-title").SetSubTitle("filecast-subtitle").SetBody("filecast-body")
	n.Payload.AddExtra("extra1", "1").AddExtra("extra2", "2")
	n.Policy.SetOutBizNo("filecast-" + strconv.Itoa(int(time.Now().Unix()))) // 设置外部id，主要用于防止重复推送，只对任务类消息生效
	n.SetFileId(fileId)
	n.BindApp(u.IOS)
	return n.Send()
}
