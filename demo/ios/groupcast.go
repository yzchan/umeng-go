package ios

import (
	"fmt"
	"github.com/yzchan/umeng-go/push"
	"github.com/yzchan/umeng-go/push/notification/ios"
	"strconv"
	"time"
)

// SendGroupcast 分组推送的条件过滤功能十分强大，可以按照各种维度进行筛选
// 这里以过滤标签维度作为示例
// 更多条件过滤设置请参考友盟官方文档
func SendGroupcast(u *push.Umeng, tag string) (string, error) {
	n := ios.NewGroupcast()
	n.SetDescription("groupcast-test")
	n.Payload.GetAPNs().SetTitle("groupcast-title").SetSubTitle("groupcast-subtitle").SetBody("groupcast-body" + ",tag:" + tag)
	n.Payload.AddExtra("extra1", "1").AddExtra("extra2", "2")
	n.Policy.SetOutBizNo("groupcast-" + strconv.Itoa(int(time.Now().Unix()))) // 设置外部id，主要用于防止重复推送，只对任务类消息生效
	condition := fmt.Sprintf(`{"where":{"and": [{"tag": "%s"}]}}`, tag)
	n.SetFilter(condition)
	n.BindApp(u.IOS)
	return n.Send()
}
