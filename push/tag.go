package push

import (
	"io/ioutil"
	"strings"
	"time"
)

func (u *Client) AddTag(device string, tags []string) (ret string, err error) {
	data := struct {
		Appkey       string `json:"appkey"`
		Timestamp    int64  `json:"timestamp"`
		DeviceTokens string `json:"device_tokens"`
		Tag          string `json:"tag"`
	}{
		Appkey:       u.Appkey,
		Timestamp:    time.Now().Unix(),
		DeviceTokens: device,
		Tag:          strings.Join(tags, ","),
	}

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
	data := struct {
		Appkey       string `json:"appkey"`
		Timestamp    int64  `json:"timestamp"`
		DeviceTokens string `json:"device_tokens"`
	}{
		Appkey:       u.Appkey,
		Timestamp:    time.Now().Unix(),
		DeviceTokens: device,
	}

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
	data := struct {
		Appkey       string `json:"appkey"`
		Timestamp    int64  `json:"timestamp"`
		DeviceTokens string `json:"device_tokens"`
		Tag          string `json:"tag"`
	}{
		Appkey:       u.Appkey,
		Timestamp:    time.Now().Unix(),
		DeviceTokens: device,
		Tag:          strings.Join(tags, ","),
	}

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
	data := struct {
		Appkey       string `json:"appkey"`
		Timestamp    int64  `json:"timestamp"`
		DeviceTokens string `json:"device_tokens"`
		Tag          string `json:"tag"`
	}{
		Appkey:       u.Appkey,
		Timestamp:    time.Now().Unix(),
		DeviceTokens: device,
		Tag:          strings.Join(tags, ","),
	}

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
	data := struct {
		Appkey       string `json:"appkey"`
		Timestamp    int64  `json:"timestamp"`
		DeviceTokens string `json:"device_tokens"`
	}{
		Appkey:       u.Appkey,
		Timestamp:    time.Now().Unix(),
		DeviceTokens: device,
	}

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
