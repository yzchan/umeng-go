package push

import (
	"encoding/json"
	"errors"
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

	defer resp.Body.Close()
	retStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	result := UploadResp{}
	err = json.Unmarshal(retStr, &result)
	if err != nil {
		return
	}
	if result.Ret == "FAIL" {
		return "", errors.New("upload err:" + result.Data.ErrorMsg)
	}
	fileId = result.Data.FileId
	return
}
