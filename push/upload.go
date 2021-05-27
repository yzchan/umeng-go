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
		ErrorCode string `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
	} `json:"data"`
}

func (u *Client) UploadFile(file string) (fileId string, err error) {
	var content []byte
	if content, err = ioutil.ReadFile(file); err != nil {
		return
	}
	return u.Upload(string(content))
}

func (u *Client) Upload(content string) (fileId string, err error) {
	var result []byte
	data := UploadReq{u.Appkey, time.Now().Unix(), content}
	if result, err = u.Request(Host+UploadPath, data); err != nil {
		return
	}

	r := UploadResp{}
	if err = json.Unmarshal(result, &r); err != nil {
		return
	}
	fileId = r.Data.FileId
	return
}
