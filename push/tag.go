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

func (c *Client) buildTagReq() (req *tagReq) {
	req = new(tagReq)
	req.Appkey = c.AppKey
	req.Timestamp = time.Now().Unix()
	return req
}

func (c *Client) ListTags(device string) (tags []string, err error) {
	data := c.buildTagReq().setDeviceToken(device)
	result, err := c.Request(Host+TagListPath, data)
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

func (c *Client) AddTags(device string, tags []string) (err error) {
	data := c.buildTagReq().setDeviceToken(device).setTags(tags)
	_, err = c.Request(Host+TagAddPath, data)
	return
}

func (c *Client) AddTag(device string, tag string) (err error) {
	return c.AddTags(device, []string{tag})
}

func (c *Client) SetTags(device string, tags []string) (err error) {
	data := c.buildTagReq().setDeviceToken(device).setTags(tags)
	_, err = c.Request(Host+TagSetPath, data)
	return
}

func (c *Client) DeleteTags(device string, tags []string) (err error) {
	data := c.buildTagReq().setDeviceToken(device).setTags(tags)
	_, err = c.Request(Host+TagDeletePath, data)
	return
}

func (c *Client) DeleteTag(device string, tag string) (err error) {
	return c.DeleteTags(device, []string{tag})
}

func (c *Client) ClearTags(device string) (err error) {
	data := c.buildTagReq().setDeviceToken(device)
	_, err = c.Request(Host+TagClearPath, data)
	return
}
