package push

import (
	"encoding/json"
	"time"
)

type QuotaReq struct {
	Appkey    string `json:"appkey"`
	Timestamp int64  `json:"timestamp"`
}

type QuotaResp struct {
	Ret  string    `json:"ret"`
	Data QuotaData `json:"data"`
}

type QuotaData struct {
	VivoSysMsgCount    string `json:"vivoSysMsgCount"`
	XmAckedCount       string `json:"xmAckedCount"`
	OppoTotalCount     string `json:"oppoTotalCount"`
	XmQuotaCount       string `json:"xmQuotaCount"`
	OppoPushCount      string `json:"oppoPushCount"`
	VivoMarketMsgCount string `json:"vivoMarketMsgCount"`
	OppoRemainCount    string `json:"oppoRemainCount"`
}

func (u *Client) Quota() (ret QuotaData, err error) {
	var result []byte
	data := QuotaReq{u.Appkey, time.Now().Unix()}
	if result, err = u.Request(Host+QuotaPath, data); err != nil {
		return
	}
	var q QuotaResp
	if err = json.Unmarshal(result, &q); err != nil {
		return
	}
	ret = q.Data
	return
}
