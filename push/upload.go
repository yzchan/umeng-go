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
		FileId string `json:"file_id"`
	} `json:"data"`
}

func (c *Client) UploadFile(file string) (fileId string, err error) {
	var content []byte
	if content, err = ioutil.ReadFile(file); err != nil {
		return
	}
	return c.Upload(string(content))
}

func (c *Client) Upload(content string) (fileId string, err error) {
	var result []byte
	data := UploadReq{c.AppKey, time.Now().Unix(), content}
	if result, err = c.Request(Host+UploadPath, data); err != nil {
		return
	}

	r := UploadResp{}
	if err = json.Unmarshal(result, &r); err != nil {
		return
	}
	fileId = r.Data.FileId
	return
}
