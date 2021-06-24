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

func (a *App) AddTemplate(template interface{}) (templateId string, err error) {
	var (
		buf []byte
		r   AddTemplateResp
	)

	if buf, err = a.Request(Host+TmplAddPath, template); err != nil {
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

func (a *App) GetTemplate(templateId string) (ret *Template, err error) {
	payload := TemplateReq{
		Appkey:     a.AppKey,
		Timestamp:  time.Now().Unix(),
		TemplateId: templateId,
	}

	var buf []byte
	if buf, err = a.Request(Host+TmplGetPath, payload); err != nil {
		return
	}

	var r TemplateResp
	if err = json.Unmarshal(buf, &r); err != nil {
		return
	}

	return &r.Data, nil
}

func (a *App) DeleteTemplate(templateId string) (err error) {
	payload := TemplateReq{
		Appkey:     a.AppKey,
		Timestamp:  time.Now().Unix(),
		TemplateId: templateId,
	}

	if _, err = a.Request(Host+TmplDeletePath, payload); err != nil {
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

func (a *App) ListTemplate(page int, limit int) (ret *ListTemplateData, err error) {
	payload := ListTemplateReq{
		Appkey:    a.AppKey,
		Timestamp: time.Now().Unix(),
		Index:     page,
		Len:       limit,
	}

	var buf []byte
	if buf, err = a.Request(Host+TmplListPath, payload); err != nil {
		return
	}

	var r ListTemplateResp
	if err = json.Unmarshal(buf, &r); err != nil {
		return
	}

	return &r.Data, nil
}

func (a *App) GetTemplateCount() int {
	ret, err := a.ListTemplate(1, 1)
	if err != nil {
		return 0
	}
	return ret.Total
}

type MsgResp struct {
	Ret  string `json:"ret"`
	Data struct {
		List Msg `json:"list"`
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

func (a *App) SendTemplateMsg(templateId string, data interface{}) (templateMsgId string, err error) {
	return
}

func (a *App) GetMsg(templateMsgId string) (ret []Msg, err error) {
	return
}
