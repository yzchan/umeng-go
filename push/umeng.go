package push

import (
	"strings"
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

	IOS     string = "ios"
	Android string = "android"
)

var (
	Debug       bool
	PrettyJson  bool
)

type Umeng struct {
	Android *App
	IOS     *App
}

func NewUmeng() *Umeng {
	return &Umeng{
		Android: &App{Platform: Android},
		IOS:     &App{Platform: IOS},
	}
}

func (u *Umeng) Debug() *Umeng {
	Debug = true
	return u
}

func (u *Umeng) InitAndroid(appkey string, secret string) *Umeng {
	u.Android.AppKey = appkey
	u.Android.MasterSecret = secret
	return u
}
func (u *Umeng) SetPackageName(packageName string) *Umeng {
	u.Android.PackageName = packageName
	return u
}

func (u *Umeng) InitIOS(appkey string, secret string) *Umeng {
	u.IOS.AppKey = appkey
	u.IOS.MasterSecret = secret
	return u
}

func (u *Umeng) GetApp(platform string) *App {
	if strings.ToLower(platform) == IOS {
		return u.IOS
	} else if strings.ToLower(platform) == Android {
		return u.Android
	}
	return &App{}
}
