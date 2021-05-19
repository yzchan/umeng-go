package push

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/yzchan/umeng-go/push/notification"
	"log"
	"net/http"
)

const (
	Host       string = "http://msg.umeng.com"
	HttpsHost  string = "https://msgapi.umeng.com"
	SendPath   string = "/api/send"
	StatusPath string = "/api/status"
	CancelPath string = "/api/cancel"
	UploadPath string = "/upload"

	TagAddPath    string = "/api/tag/add"
	TagListPath   string = "/api/tag/list"
	TagSetPath    string = "/api/tag/set"
	TagDeletePath string = "/api/tag/delete"
	TagClearPath  string = "/api/tag/clear"

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

func (u *Client) Send(cast notification.Caster) (*http.Response, error) {
	cast.SetAppkey(u.Appkey)
	cast.SetTimestamp()
	return u.Request(Host+SendPath, cast)
}

func (u *Client) Request(url string, reqBody interface{}) (*http.Response, error) {
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	url = fmt.Sprintf("%s?sign=%s", url, u.Sign(url, string(body)))
	log.Println(url)
	log.Println(string(body))
	return http.Post(url, "application/json", bytes.NewBuffer(body))
}

func (u *Client) Sign(url string, body string) string {
	str := fmt.Sprintf("%s%s%s%s", "POST", url, body, u.MasterSecret)
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
