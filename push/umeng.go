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

const (
	Host        string = "http://msg.umeng.com" // https://msgapi.umeng.com
	SendPath    string = "/api/send"
	StatusPath  string = "/api/status"
	ChannelPath string = "/api/channel/data"
	QuotaPath   string = "/api/quota/query"
	CancelPath  string = "/api/cancel"
	UploadPath  string = "/upload"

	TagAddPath    string = "/api/tag/add"
	TagListPath   string = "/api/tag/list"
	TagSetPath    string = "/api/tag/set"
	TagDeletePath string = "/api/tag/delete"
	TagClearPath  string = "/api/tag/clear"

	//TmplAddPath    string = "/api/template/add"
	//TmplDeletePath string = "/api/template/delete"
	//TmplGetPath    string = "/api/template/get"
	//TmplListPath   string = "/api/template/list"
	//TmplSendPath   string = "/api/template/send"
	//TmplMsgPath    string = "/api/template/msg"

	IOS     int = 0
	Android int = 1
)

type Client struct {
	Appkey       string
	MasterSecret string
	PackageName  string
	Platform     int
}

type Umeng struct {
	Android Client
	IOS     Client
}

func NewUmeng() *Umeng {
	return &Umeng{
		Android: Client{Platform: Android},
		IOS:     Client{Platform: IOS},
	}
}

func (u *Umeng) InitAndroid(appkey string, secret string) *Umeng {
	u.Android.Appkey = appkey
	u.Android.MasterSecret = secret
	return u
}

func (u *Umeng) SetPackageName(name string) *Umeng {
	u.Android.PackageName = name
	return u
}

func (u *Umeng) InitIOS(appkey string, secret string) *Umeng {
	u.IOS.Appkey = appkey
	u.IOS.MasterSecret = secret
	return u
}

// PreView 预览推送消息体
func (u *Client) PreView(cast notification.Caster) {
	cast.SetAppkey(u.Appkey)
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
func (u *Client) Send(cast notification.Caster) (string, error) {
	var (
		buf []byte
		err error
	)
	cast.SetAppkey(u.Appkey)
	if buf, err = u.Request(Host+SendPath, cast); err != nil {
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

func (u *Client) Request(url string, reqBody interface{}) (content []byte, err error) {
	var (
		body []byte
		resp *http.Response
	)

	if body, err = json.Marshal(reqBody); err != nil {
		return
	}

	url = fmt.Sprintf("%s?sign=%s", url, u.Sign(url, string(body)))
	if resp, err = http.Post(url, "application/json", bytes.NewBuffer(body)); err != nil {
		return
	}
	defer resp.Body.Close()

	if content, err = ioutil.ReadAll(resp.Body); err != nil {
		return
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

func (u *Client) Sign(url string, body string) string {
	str := fmt.Sprintf("%s%s%s%s", "POST", url, body, u.MasterSecret)
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
