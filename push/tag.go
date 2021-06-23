package push

import (
	"encoding/json"
	"strings"
	"time"
)

type tagReq struct {
	Appkey       string `json:"appkey"`
	Timestamp    int64  `json:"timestamp"`
	DeviceTokens string `json:"device_tokens"`
	Tag          string `json:"tag,omitempty"`
}

func (t *tagReq) setTags(tags []string) *tagReq {
	t.Tag = strings.Join(tags, ",")
	return t
}
func (t *tagReq) setDeviceToken(device string) *tagReq {
	t.DeviceTokens = device
	return t
}

func (a *App) buildTagReq() (req *tagReq) {
	req = new(tagReq)
	req.Appkey = a.AppKey
	req.Timestamp = time.Now().Unix()
	return req
}

// ListTags 查询设备标签
func (a *App) ListTags(device string) (tags []string, err error) {
	data := a.buildTagReq().setDeviceToken(device)
	result, err := a.Request(Host+TagListPath, data)
	if err != nil {
		return
	}
	var v struct {
		Ret  string `json:"ret"`
		Data struct {
			Data struct {
				Tags string `json:"tags"`
			} `json:"data"`
		} `json:"data"`
	}
	err = json.Unmarshal(result, &v)
	if err != nil {
		return
	}
	if v.Data.Data.Tags == "" {
		return []string{}, nil
	}
	return strings.Split(v.Data.Data.Tags, ","), nil
}

// AddTags 给设备添加标签
func (a *App) AddTags(device string, tags []string) (err error) {
	data := a.buildTagReq().setDeviceToken(device).setTags(tags)
	_, err = a.Request(Host+TagAddPath, data)
	return
}

func (a *App) AddTag(device string, tag string) (err error) {
	return a.AddTags(device, []string{tag})
}

// SetTags 给设备重设标签，该方法会清掉原来设置的tag
func (a *App) SetTags(device string, tags []string) (err error) {
	data := a.buildTagReq().setDeviceToken(device).setTags(tags)
	_, err = a.Request(Host+TagSetPath, data)
	return
}

// DeleteTags 删除设备标签
func (a *App) DeleteTags(device string, tags []string) (err error) {
	data := a.buildTagReq().setDeviceToken(device).setTags(tags)
	_, err = a.Request(Host+TagDeletePath, data)
	return
}

// DeleteTag 删除单个标签
func (a *App) DeleteTag(device string, tag string) (err error) {
	return a.DeleteTags(device, []string{tag})
}

// ClearTags 清除设备标签
func (a *App) ClearTags(device string) (err error) {
	data := a.buildTagReq().setDeviceToken(device)
	_, err = a.Request(Host+TagClearPath, data)
	return
}
