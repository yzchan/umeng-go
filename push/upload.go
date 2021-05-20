package push

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type UploadResp struct {
	Ret  string `json:"ret"`
	Data struct {
		FileId    string `json:"file_id"`
		ErrorCode string `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
	} `json:"data"`
}

func (u *Client) UploadFile(file string) (fileId string, err error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	return u.Upload(string(content))
}

func (u *Client) Upload(content string) (fileId string, err error) {

	data := struct {
		Appkey    string `json:"appkey"`
		Timestamp int64  `json:"timestamp"`
		Content   string `json:"content"`
	}{
		Appkey:    u.Appkey,
		Timestamp: time.Now().Unix(),
		Content:   content,
	}

	resp, err := u.Request(Host+UploadPath, data)
	if err != nil {
		return
	}

	result := UploadResp{}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	fileId = result.Data.FileId
	return
}
