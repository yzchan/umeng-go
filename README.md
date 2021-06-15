# umeng-go

[友盟](https://www.umeng.com/) 相关接口的封装

### About 友盟推送 U-Push

[友盟推送接口](https://developer.umeng.com/docs/67966/detail/68343) 封装。其中部分接口需要U-Push Pro版本的支持。

- 友盟推送：消息的推送、任务查询、撤销推送、文件上传
- 标签接口：标签的查询、新增、重设、删除、清除
- 模板消息(TODO)
- 安卓厂商数据透出和额度查询

消息类型包括：单播(unicast)、列播(listcast)、广播(broadcast)、组播(groupcast)、文件播(filecast)、自定义播(customizedcast)
。其中自定义播又可以分为自定义文件播和自定义alias播。 这7种类型整体可以分为两大类：单播类消息和任务类消息。

> unicast、listcast、customizedcast(不带file_id)统称为单播类消息，Web后台不会展示此类消息详细信息，仅展示前一天的汇总数据。单播类消息推送响应体中返回msg_id。

> broadcast、filecast、groupcast、customizedcast(带file_id)统称为任务类消息，任务类消息支持API查询、撤销操作，Web后台会展示任务类消息详细信息。任务类消息推送响应体中返回task_id。

| 消息类别 | 消息类型 | 用户范围 | msg_id | task_id | API支持 |
| ------ | ------ | ------ | ------ | ------ | ------ |
| 单播类 | 单播 | 独立用户 | uu****** | | |
| 单播类 | 列播 | 独立用户 | ul****** | | |
| 单播类 | 自定义alias播 | 特定用户 | ua****** | | |
| 任务类 | 广播 | 全部用户 | | um****** | |
| 任务类 | 组播 | 部分用户 | | us****** | |
| 任务类 | 文件播 | 部分用户 | | uf****** | 只支持API |
| 任务类 | 自定义文件播 | 特定用户 | | uc****** | 只支持API |

### Quickstart

```go
package main

import (
	"fmt"
	"github.com/yzchan/umeng-go/push"
	"github.com/yzchan/umeng-go/push/notification/android"
)

func main() {
	u := push.NewUmeng()
	u.InitAndroid("*****", "*****").SetPackageName("com.***.forum")
	u.InitIOS("*****", "*****")

	// 查询推送任务结果
	ret1, err := u.Android.Status("us*****")
	fmt.Println(ret1, err)

	// 撤销推送
	ret2, err := u.Android.Cancel("um*****")
	fmt.Println(ret2, err)

	// 新建一条单播推送
	cast := android.NewUnicast()
	cast.SetDescription("单播测试")
	cast.Payload.Body.SetTitle("push title").SetText("push text")
	cast.Payload.Extra.AddKV("extra1", "1").AddKV("extra2", "2")
	cast.SetDeviceToken("{input your device-token}") // 设置推送目标的设备

	//u.Android.PreView(cast) // 预览json格式的消息体
	ret, err := u.Android.Send(cast)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(ret))
}
```

> 更多示例程序见demo
