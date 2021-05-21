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
	Host           string = "http://msg.umeng.com"
	HttpsHost      string = "https://msgapi.umeng.com"
	SendPath       string = "/api/send"
	StatusPath     string = "/api/status"
	ChanDataPath   string = "/api/channel/data"
	QuotaQueryPath string = "/api/quota/query"
	CancelPath     string = "/api/cancel"
	UploadPath     string = "/upload"

	TagAddPath    string = "/api/tag/add"
	TagListPath   string = "/api/tag/list"
	TagSetPath    string = "/api/tag/set"
	TagDeletePath string = "/api/tag/delete"
	TagClearPath  string = "/api/tag/clear"

	TmplAddPath    string = "/api/template/add"
	TmplDeletePath string = "/api/template/delete"
	TmplGetPath    string = "/api/template/get"
	TmplListPath   string = "/api/template/list"
	TmplSendPath   string = "/api/template/send"
	TmplMsgPath    string = "/api/template/msg"

	PlatformAndroid int = 1
	PlatformIOS     int = 0
)

type Client struct {
	Appkey       string
	MasterSecret string
	Platform     int
}

type Umeng struct {
	Android Client
	IOS     Client
}

func NewUmeng() *Umeng {
	return &Umeng{
		Android: Client{Platform: PlatformAndroid},
		IOS:     Client{Platform: PlatformIOS},
	}
}

func (u *Umeng) InitAndroid(appkey string, secret string) *Umeng {
	u.Android.Appkey = appkey
	u.Android.MasterSecret = secret
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

// Check 检测消息是否合法
func (u *Client) Check(cast notification.Caster) {
	// TODO
}

// Send 发送消息
func (u *Client) Send(cast notification.Caster) ([]byte, error) {
	cast.SetAppkey(u.Appkey)
	result, err := u.Request(Host+SendPath, cast)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type FailResp struct {
	Ret  string     `json:"ret"`
	Data UmengError `json:"data"`
}

type UmengError struct {
	Code string `json:"error_code"`
	Msg  string `json:"error_msg"`
}

func (e *UmengError) Error() string {
	return fmt.Sprintf("Umeng Resonse Error:[%s]%s", e.Code, e.Msg)
}

func (u *Client) Request(url string, reqBody interface{}) ([]byte, error) {
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	url = fmt.Sprintf("%s?sign=%s", url, u.Sign(url, string(body)))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var content []byte
	content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(content))
	// 统一处理请求失败
	if resp.StatusCode == 400 {
		var fail FailResp
		err = json.Unmarshal(content, &fail)
		if err != nil {
			return nil, err
		}
		return nil, &fail.Data
	} else if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("umeng response code:%d. %s", resp.StatusCode, content))
	}
	return content, nil
}

func (u *Client) Sign(url string, body string) string {
	str := fmt.Sprintf("%s%s%s%s", "POST", url, body, u.MasterSecret)
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
