package push

import (
	"encoding/json"
	"time"
)

type AddTemplateResp struct {
	Ret  string `json:"ret"`
	Data struct {
		TemplateKeys []string `json:"template_keys"`
		TemplateId   string   `json:"template_id"`
	} `json:"data"`
}

func (c *Client) AddTemplate(template interface{}) (templateId string, err error) {
	var (
		buf []byte
		r   AddTemplateResp
	)

	if buf, err = c.Request(Host+TmplAddPath, template); err != nil {
		return
	}

	if err = json.Unmarshal(buf, &r); err != nil {
		return
	}

	return r.Data.TemplateId, nil
}

type TemplateReq struct {
	Appkey     string `json:"appkey"`
	Timestamp  int64  `json:"timestamp"`
	TemplateId string `json:"template_id"`
}

type TemplateResp struct {
	Ret  string   `json:"ret"`
	Data Template `json:"data"`
}

type Template struct {
	TemplateName string `json:"template_name"`
	TemplateInfo string `json:"template_info"`
	TemplateKeys string `json:"template_keys"`
	Appkey       string `json:"appkey"`
	TemplateId   string `json:"template_id"`
	Id           int    `json:"id"`
}

func (c *Client) GetTemplate(templateId string) (ret *Template, err error) {
	payload := TemplateReq{
		Appkey:     c.AppKey,
		Timestamp:  time.Now().Unix(),
		TemplateId: templateId,
	}

	var buf []byte
	if buf, err = c.Request(Host+TmplGetPath, payload); err != nil {
		return
	}

	var r TemplateResp
	if err = json.Unmarshal(buf, &r); err != nil {
		return
	}

	return &r.Data, nil
}

func (c *Client) DeleteTemplate(templateId string) (err error) {
	payload := TemplateReq{
		Appkey:     c.AppKey,
		Timestamp:  time.Now().Unix(),
		TemplateId: templateId,
	}

	if _, err = c.Request(Host+TmplDeletePath, payload); err != nil {
		return NewUmengError(404, "模板不存在")
	}

	return nil
}

type ListTemplateReq struct {
	Appkey    string `json:"appkey"`
	Timestamp int64  `json:"timestamp"`
	Index     int    `json:"index"`
	Len       int    `json:"len"`
}
type ListTemplateResp struct {
	Ret  string           `json:"ret"`
	Data ListTemplateData `json:"data"`
}

type ListTemplateData struct {
	Next  bool       `json:"next"`
	Pre   bool       `json:"pre"`
	Total int        `json:"total"`
	Len   int        `json:"len"`
	Index int        `json:"index"`
	List  []Template `json:"list"`
}

func (c *Client) ListTemplate(page int, limit int) (ret *ListTemplateData, err error) {
	payload := ListTemplateReq{
		Appkey:    c.AppKey,
		Timestamp: time.Now().Unix(),
		Index:     page,
		Len:       limit,
	}

	var buf []byte
	if buf, err = c.Request(Host+TmplListPath, payload); err != nil {
		return
	}

	var r ListTemplateResp
	if err = json.Unmarshal(buf, &r); err != nil {
		return
	}

	return &r.Data, nil
}

func (c *Client) GetTemplateCount() int {
	ret, err := c.ListTemplate(1, 1)
	if err != nil {
		return 0
	}
	return ret.Total
}

type SendTemplateMsgReq struct {
	Appkey     string        `json:"appkey"`
	Timestamp  int64         `json:"timestamp"`
	TemplateId string        `json:"template_id"`
	ParamsData []interface{} `json:"params_data"`
}

type SendTemplateMsgResp struct {
	Ret  string `json:"ret"`
	Data struct {
		TemplateMsgId string `json:"template_msg_id"`
	} `json:"data"`
}

func (c *Client) SendTemplateMsg(templateId string, data []interface{}) (templateMsgId string, err error) {
	payload := SendTemplateMsgReq{
		Appkey:     c.AppKey,
		Timestamp:  time.Now().Unix(),
		TemplateId: templateId,
		ParamsData: data,
	}

	var buf []byte
	if buf, err = c.Request(Host+TmplSendPath, payload); err != nil {
		return
	}

	var r SendTemplateMsgResp
	if err = json.Unmarshal(buf, &r); err != nil {
		return
	}

	return r.Data.TemplateMsgId, nil
}

type MsgReq struct {
	Appkey        string `json:"appkey"`
	Timestamp     int64  `json:"timestamp"`
	TemplateMsgId string `json:"template_msg_id"`
	StartKey      string `json:"start_key,omitempty"`
	Len           int    `json:"len"`
}

type MsgResp struct {
	Ret  string `json:"ret"`
	Data struct {
		List []Msg `json:"list"`
	} `json:"data"`
}

type Msg struct {
	Appkey        string `json:"appkey"`
	Index         int    `json:"index"`
	MsgId         string `json:"msgId"`
	Params        string `json:"params"`
	TemplateId    string `json:"templateId"`
	TemplateMsgId string `json:"templateMsgId"`
}

func (c *Client) GetMsg(templateMsgId string, limit int, cursor string) (ret []Msg, last string, err error) {
	payload := MsgReq{
		Appkey:        c.AppKey,
		Timestamp:     time.Now().Unix(),
		TemplateMsgId: templateMsgId,
		Len:           limit,
		StartKey:      cursor,
	}

	var buf []byte
	if buf, err = c.Request(Host+TmplMsgPath, payload); err != nil {
		return
	}

	var r MsgResp
	if err = json.Unmarshal(buf, &r); err != nil {
		return
	}

	if len(r.Data.List) > 0 {
		lastMsg := r.Data.List[len(r.Data.List)-1]
		last = lastMsg.MsgId
	}

	return r.Data.List, last, nil
}
