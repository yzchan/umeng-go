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
	"net/url"
	"time"
)

type CastRequester interface {
	SetAppKey(key string)
	GetUri() string
}

type Client struct {
	AppKey       string
	MasterSecret string
	Platform     string
	PackageName  string // 包名（安卓离线推送需要）
	Debug        bool
}

func (c *Client) Request(uri string, reqBody interface{}) (content []byte, err error) {
	var (
		body  []byte
		resp  *http.Response
		req   *http.Request
		proxy func(*http.Request) (*url.URL, error)
	)

	body, err = json.Marshal(reqBody)

	if err != nil {
		return
	}

	uri = fmt.Sprintf("%s?sign=%s", uri, c.Sign(uri, string(body)))

	if c.Debug {
		log.Println("==================== Umeng Debug Start ====================")
		log.Printf("[Request] POST %s\n%s\n", uri, string(body))
	}

	if Proxy != "" {
		proxy = func(_ *http.Request) (*url.URL, error) {
			return url.Parse(Proxy)
		}
	}

	client := &http.Client{
		Transport: &http.Transport{Proxy: proxy},
		Timeout:   time.Second * 5,
	}
	if req, err = http.NewRequest("POST", uri, bytes.NewBuffer(body)); err != nil {
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
	if c.Debug {
		log.Printf("[Response] Http Code:%d\n%s\n", resp.StatusCode, string(content))
		log.Println("==================== Umeng Debug End ====================")
		log.Println()
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

func (c *Client) Send(n CastRequester) (content []byte, err error) {

	n.SetAppKey(c.AppKey)

	return c.Request(n.GetUri(), n)
}

func (c *Client) Sign(url string, body string) string {
	str := fmt.Sprintf("%s%s%s%s", "POST", url, body, c.MasterSecret)
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
