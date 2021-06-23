package push

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type UploadReq struct {
	Appkey    string `json:"appkey"`
	Timestamp int64  `json:"timestamp"`
	Content   string `json:"content"`
}

type UploadResp struct {
	Ret  string `json:"ret"`
	Data struct {
		FileId    string `json:"file_id"`
	} `json:"data"`
}

func (a *App) UploadFile(file string) (fileId string, err error) {
	var content []byte
	if content, err = ioutil.ReadFile(file); err != nil {
		return
	}
	return a.Upload(string(content))
}

func (a *App) Upload(content string) (fileId string, err error) {
	var result []byte
	data := UploadReq{a.AppKey, time.Now().Unix(), content}
	if result, err = a.Request(Host+UploadPath, data); err != nil {
		return
	}

	r := UploadResp{}
	if err = json.Unmarshal(result, &r); err != nil {
		return
	}
	fileId = r.Data.FileId
	return
}
