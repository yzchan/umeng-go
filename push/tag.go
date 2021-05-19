package push

import (
	"io/ioutil"
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

func (u *Client) AddTag(device string, tags []string) (ret string, err error) {
	data := u.buildTagReq().setDeviceToken(device).setTags(tags)

	resp, err := u.Request(Host+TagAddPath, data)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	retStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	ret = string(retStr)
	return
}

// ListTag 查询设备标签
func (u *Client) ListTag(device string) (ret string, err error) {
	data := u.buildTagReq().setDeviceToken(device)

	resp, err := u.Request(Host+TagListPath, data)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	retStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	ret = string(retStr)
	return
}

// SetTag 该方法会清掉原来设置的tag
func (u *Client) SetTag(device string, tags []string) (ret string, err error) {
	data := u.buildTagReq().setDeviceToken(device).setTags(tags)

	resp, err := u.Request(Host+TagSetPath, data)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	retStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	ret = string(retStr)
	return
}

// DeleteTag 删除设备标签
func (u *Client) DeleteTag(device string, tags []string) (ret string, err error) {
	data := u.buildTagReq().setDeviceToken(device).setTags(tags)

	resp, err := u.Request(Host+TagDeletePath, data)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	retStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	ret = string(retStr)
	return
}

// ClearTag 清除设备标签
func (u *Client) ClearTag(device string) (ret string, err error) {
	data := u.buildTagReq().setDeviceToken(device)

	resp, err := u.Request(Host+TagClearPath, data)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	retStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	ret = string(retStr)
	return
}
