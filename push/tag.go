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

func (u *Client) buildTagReq() (req *tagReq) {
	req = new(tagReq)
	req.Appkey = u.Appkey
	req.Timestamp = time.Now().Unix()
	return req
}

// ListTags 查询设备标签
func (u *Client) ListTags(device string) (tags []string, err error) {
	data := u.buildTagReq().setDeviceToken(device)
	result, err := u.Request(Host+TagListPath, data)
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
	return strings.Split(v.Data.Data.Tags, ","), nil
}

// AddTags 添加标签
func (u *Client) AddTags(device string, tags []string) (err error) {
	data := u.buildTagReq().setDeviceToken(device).setTags(tags)
	_, err = u.Request(Host+TagAddPath, data)
	return
}

func (u *Client) AddTag(device string, tag string) (err error) {
	return u.AddTags(device, []string{tag})
}

// SetTags 该方法会清掉原来设置的tag
func (u *Client) SetTags(device string, tags []string) (err error) {
	data := u.buildTagReq().setDeviceToken(device).setTags(tags)
	_, err = u.Request(Host+TagSetPath, data)
	return
}

// DeleteTags 删除设备标签
func (u *Client) DeleteTags(device string, tags []string) (err error) {
	data := u.buildTagReq().setDeviceToken(device).setTags(tags)
	_, err = u.Request(Host+TagDeletePath, data)
	return
}

// DeleteTag 删除单个标签
func (u *Client) DeleteTag(device string, tag string) (err error) {
	return u.DeleteTags(device, []string{tag})
}

// ClearTags 清除设备标签
func (u *Client) ClearTags(device string) (err error) {
	data := u.buildTagReq().setDeviceToken(device)
	_, err = u.Request(Host+TagClearPath, data)
	return
}
