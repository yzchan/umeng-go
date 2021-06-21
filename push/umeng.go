package push

import "strings"

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

type Umeng struct {
	Android Platform
	IOS     Platform
}

func NewUmeng() *Umeng {
	return &Umeng{
		Android: Platform{Platform: Android},
		IOS:     Platform{Platform: IOS},
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

func (u *Umeng) GetPlatform(platform string) *Platform {
	if strings.ToLower(platform) == IOS {
		return &u.IOS
	} else if strings.ToLower(platform) == Android {
		return &u.Android
	}
	return &Platform{}
}
