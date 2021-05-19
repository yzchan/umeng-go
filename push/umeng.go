package push

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/yzchan/umeng-go/push/notification"
	"net/http"
)

const (
	Host       string = "http://msg.umeng.com"
	HttpsHost  string = "https://msgapi.umeng.com"
	SendPath   string = "/api/send"
	StatusPath string = "/api/status"
	CancelPath string = "/api/cancel"
	UploadPath string = "/upload"

	PlatformAndroid int64 = 1
	PlatformIOS     int64 = 0
)

type Client struct {
	Appkey       string `json:"appkey"`
	MasterSecret string `json:"-"`
	Platform     int64  `json:"platform"`
}

type Umeng struct {
	Android Client
	IOS     Client
}

func NewUmeng(androidKey string, androidSecret string, iosKey string, iosSecret string) *Umeng {
	return &Umeng{
		Android: Client{
			Appkey:       androidKey,
			MasterSecret: androidSecret,
			Platform:     1,
		},
		IOS: Client{
			Appkey:       iosKey,
			MasterSecret: iosSecret,
			Platform:     0,
		},
	}
}

func (u *Client) Send(cast notification.Caster) (*http.Response, error) {
	cast.SetAppkey(u.Appkey)
	cast.SetTimestamp()
	return u.Request(Host+SendPath, cast)
}

func (u *Client) Request(url string, reqBody interface{}) (*http.Response, error) {
	body, err := json.MarshalIndent(reqBody, "", "    ")
	if err != nil {
		return nil, err
	}
	url = fmt.Sprintf("%s?sign=%s", url, u.Sign(url, string(body)))
	fmt.Println(url)
	fmt.Println(string(body))
	return http.Post(url, "application/json", bytes.NewBuffer(body))
}

func (u *Client) Sign(url string, body string) string {
	str := fmt.Sprintf("%s%s%s%s", "POST", url, body, u.MasterSecret)
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
