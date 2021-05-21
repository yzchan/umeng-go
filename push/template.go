package push

import (
	"github.com/yzchan/umeng-go/push/notification"
	"time"
)

func (u *Client) AddTemplate(cast notification.Caster) ([]byte, error) {
	cast.SetAppkey(u.Appkey)
	result, err := u.Request(Host+TmplAddPath, cast)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *Client) ListTemplate(page int, limit int) ([]byte, error) {
	data := struct {
		Appkey    string `json:"appkey"`
		Timestamp int64  `json:"timestamp"`
		Index     int    `json:"index"`
		Len       int    `json:"len"`
	}{
		Appkey:    u.Appkey,
		Timestamp: time.Now().Unix(),
		Index:     page,  // 分页
		Len:       limit, // 每页数量
	}
	result, err := u.Request(Host+TmplListPath, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *Client) GetTemplate(id string) ([]byte, error) {
	data := struct {
		Appkey     string `json:"appkey"`
		Timestamp  int64  `json:"timestamp"`
		TemplateId string `json:"template_id"`
	}{
		Appkey:     u.Appkey,
		Timestamp:  time.Now().Unix(),
		TemplateId: id,
	}
	result, err := u.Request(Host+TmplGetPath, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *Client) DeleteTemplate(id string) ([]byte, error) {
	data := struct {
		Appkey     string `json:"appkey"`
		Timestamp  int64  `json:"timestamp"`
		TemplateId string `json:"template_id"`
	}{
		Appkey:     u.Appkey,
		Timestamp:  time.Now().Unix(),
		TemplateId: id,
	}
	result, err := u.Request(Host+TmplDeletePath, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *Client) SendTemplateMessages(data interface{}) ([]byte, error) {
	return nil, nil
}

func (u *Client) GetTemplateMessageId(msgId string) ([]byte, error) {
	return nil, nil
}
