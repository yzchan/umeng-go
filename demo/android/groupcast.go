package android

import (
	"fmt"
	"github.com/yzchan/umeng-go/push"
	"github.com/yzchan/umeng-go/push/notification/android"
	"strconv"
	"time"
)

// SendGroupcast 分组推送的条件过滤功能十分强大，可以按照各种维度进行筛选
// 这里以过滤标签维度作为示例
// 更多条件过滤设置请参考友盟官方文档
func SendGroupcast(u *push.Umeng, tag string) (string, error) {
	n := android.NewGroupcast()
	n.SetDescription("groupcast-test")
	n.Payload.Body.SetTitle("groupcast-title").SetText("groupcast-text" + ",tag:" + tag)
	n.Payload.Extra.AddKV("extra1", "1").AddKV("extra2", "2")
	n.Policy.SetOutBizNo("groupcast-" + strconv.Itoa(int(time.Now().Unix()))) // 设置外部id，主要用于防止重复推送，只对任务类消息生效
	condition := fmt.Sprintf(`{"where":{"and": [{"tag": "%s"}]}}`, tag)
	n.SetFilter(condition)
	n.BindApp(u.Android)
	return n.Send()
}
