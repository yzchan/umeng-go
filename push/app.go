package push

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type App struct {
	AppKey       string
	MasterSecret string
	Platform     string
	PackageName  string // 安卓包名
}

func (a *App) Marshal(data interface{}) ([]byte, error) {
	if PrettyJson {
		return json.MarshalIndent(data, "", "    ")
	} else {
		return json.Marshal(data)
	}
}

func (a *App) Request(url string, reqBody interface{}) (content []byte, err error) {
	var (
		body []byte
		resp *http.Response
		req  *http.Request
	)

	if body, err = a.Marshal(reqBody); err != nil {
		return
	}

	url = fmt.Sprintf("%s?sign=%s", url, a.Sign(url, string(body)))

	if Debug {
		log.Println("==========Umeng Request==========")
		log.Printf("POST %s\n%s\n", url, string(body))
	}
	client := http.Client{Timeout: time.Second * 5}
	if req, err = http.NewRequest("POST", url, bytes.NewBuffer(body)); err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")
	if resp, err = client.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()

	if content, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	if Debug {
		log.Println("==========Umeng Response==========")
		log.Printf("Http Code:%d\n%s\n", resp.StatusCode, string(content))
	}
	// 统一处理非200响应
	if resp.StatusCode != 200 {
		var errResp UmengErrorResp
		if err = json.Unmarshal(content, &errResp); err != nil {
			return nil, errors.New(fmt.Sprintf("[%d]:%s", resp.StatusCode, string(content)))
		}
		return nil, errResp.Data
	}
	return
}

func (a *App) Sign(url string, body string) string {
	str := fmt.Sprintf("%s%s%s%s", "POST", url, body, a.MasterSecret)
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
