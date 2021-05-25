package idcardauth

import (
	"encoding/hex"
	"fmt"
	"github.com/haxqer/gofunc"
	"github.com/haxqer/idcardauth/pkg/xhttp"
	"github.com/pquerna/ffjson/ffjson"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	AppId     string
	BizId     string
	SecureKey string
	TestCode  string
	mu        sync.RWMutex
}

// NewClient 初始化身份证认证器客户端
//	appId：应用标识
//	mchId：游戏备案识别码
//	ApiKey：用户密钥
func NewClient(appId, BizId, secureId, testCode string) (client *Client) {
	return &Client{
		AppId:     appId,
		BizId:     BizId,
		SecureKey: secureId,
		TestCode:  testCode,
		mu:        sync.RWMutex{},
	}
}

func (c *Client) doGet(basUrl string, bm BodyMap, timeout time.Duration) (bs []byte, err error) {

	return bs, nil
}

//func (c *Client) doPost(basUrl string, bm BodyMap, timeout time.Duration) (bs []byte, err error) {
//
//	httpClient := xhttp.NewClient().SetTimeout(timeout)
//
//
//	return bs, nil
//}

func (c *Client) AuthCheck(req *AuthCheckRequest) ([]byte, error) {
	marshal, err := ffjson.Marshal(req)
	if err != nil {
		return nil, err
	}
	key, err := hex.DecodeString(c.SecureKey)
	if err != nil {
		return nil, err
	}
	encode, err := gofunc.AesGCMNoPaddingEncryptBase64Encode(marshal, key)
	if err != nil {
		return nil, err
	}
	body := fmt.Sprintf("{\"data\":\"%s\"}", encode)
	nowTime := gofunc.Int64ToString(time.Now().UnixNano() / 1e6)
	signStr := fmt.Sprintf("%sappId%sbizId%stimestamps%s%s",
		c.SecureKey, c.AppId, c.BizId, nowTime, body)
	sign := gofunc.Sha256Lower(signStr)

	url := CheckUrl
	if c.TestCode != NULL {
		url = fmt.Sprintf("%s/%s", CheckTestUrl, c.TestCode)
	}

	httpClient := xhttp.NewClient()
	httpClient.Header.Set("appId", c.AppId)
	httpClient.Header.Set("bizId", c.BizId)
	httpClient.Header.Set("timestamps", nowTime)
	httpClient.Header.Set("sign", sign)
	res, bs, errs := httpClient.SetTimeout(5 * time.Second).Type(xhttp.TypeJSON).Post(url).SendString(body).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil
}

func (c *Client) AuthQuery(req *AuthQueryRequest) ([]byte, error) {
	nowTime := gofunc.Int64ToString(time.Now().UnixNano() / 1e6)

	signStr := fmt.Sprintf("%sai%sappId%sbizId%stimestamps%s",
		c.SecureKey, req.AuthId, c.AppId, c.BizId, nowTime)
	sign := gofunc.Sha256Lower(signStr)

	httpClient := xhttp.NewClient()
	httpClient.Header.Set("appId", c.AppId)
	httpClient.Header.Set("bizId", c.BizId)
	httpClient.Header.Set("timestamps", nowTime)
	httpClient.Header.Set("sign", sign)
	httpClient.SetTimeout(5 * time.Second)
	url := fmt.Sprintf("%s?ai=%s", QueryUrl, req.AuthId)
	if c.TestCode != NULL {
		url = fmt.Sprintf("%s/%s?ai=%s", QueryTestUrl, c.TestCode, req.AuthId)
	}
	res, bs, errs := httpClient.Type(xhttp.TypeJSON).Get(url).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil
}

func (c *Client) BehaviorCollect(req *BehaviorRequest) ([]byte, error) {
	marshal, err := ffjson.Marshal(req)
	if err != nil {
		return nil, err
	}
	key, err := hex.DecodeString(c.SecureKey)
	if err != nil {
		return nil, err
	}
	encode, err := gofunc.AesGCMNoPaddingEncryptBase64Encode(marshal, key)
	if err != nil {
		return nil, err
	}
	body := fmt.Sprintf("{\"data\":\"%s\"}", encode)
	nowTime := gofunc.Int64ToString(time.Now().UnixNano() / 1e6)
	signStr := fmt.Sprintf("%sappId%sbizId%stimestamps%s%s",
		c.SecureKey, c.AppId, c.BizId, nowTime, body)
	sign := gofunc.Sha256Lower(signStr)

	url := BehaviorUrl
	if c.TestCode != NULL {
		url = fmt.Sprintf("%s/%s", BehaviorTestUrl, c.TestCode)
	}
	httpClient := xhttp.NewClient()
	httpClient.Header.Set("appId", c.AppId)
	httpClient.Header.Set("bizId", c.BizId)
	httpClient.Header.Set("timestamps", nowTime)
	httpClient.Header.Set("sign", sign)
	res, bs, errs := httpClient.SetTimeout(5 * time.Second).Type(xhttp.TypeJSON).Post(url).SendString(body).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil
}
