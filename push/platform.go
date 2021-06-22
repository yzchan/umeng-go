package push

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yzchan/umeng-go/push/notification"
	"io/ioutil"
	"net/http"
	"reflect"
)

type Platform struct {
	Appkey       string
	MasterSecret string
	PackageName  string
	Platform     string
	Debug        bool
}

func (p *Platform) SetDebugMode() {
	p.Debug = true
}

// PreView 预览推送消息体
func (p *Platform) PreView(cast notification.Caster) {
	cast.SetAppkey(p.Appkey)
	t := reflect.TypeOf(cast)
	fmt.Printf("===============[%s]===============\n", t.String())
	encoded, err := json.MarshalIndent(cast, "", "    ")
	if err != nil {
		fmt.Println("json marshal err:", err.Error())
	} else {
		fmt.Println(string(encoded))
	}
	fmt.Printf("===============[%s]===============\n", t.String())
}

// Send 发送消息
func (p *Platform) Send(cast notification.Caster) (string, error) {
	var (
		buf []byte
		err error
	)
	cast.SetAppkey(p.Appkey)
	if buf, err = p.Request(Host+SendPath, cast); err != nil {
		return "", err
	}

	var r struct {
		Ret  string `json:"ret"`
		Data struct {
			MsgId  string `json:"msg_id"`
			TaskId string `json:"task_id"`
		} `json:"data"`
	}
	if err = json.Unmarshal(buf, &r); err != nil {
		return "", err
	}

	return r.Data.MsgId + r.Data.TaskId, nil
}

func (p *Platform) Request(url string, reqBody interface{}) (content []byte, err error) {
	var (
		body []byte
		resp *http.Response
	)

	if body, err = json.Marshal(reqBody); err != nil {
		return
	}

	if p.Debug {
		fmt.Printf("\n%s\n", string(body))
	}

	url = fmt.Sprintf("%s?sign=%s", url, p.Sign(url, string(body)))
	if resp, err = http.Post(url, "application/json", bytes.NewBuffer(body)); err != nil {
		return
	}
	defer resp.Body.Close()

	if content, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	if p.Debug {
		fmt.Printf("\n%s\n", string(content))
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

func (p *Platform) Sign(url string, body string) string {
	str := fmt.Sprintf("%s%s%s%s", "POST", url, body, p.MasterSecret)
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
