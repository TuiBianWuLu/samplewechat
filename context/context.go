package context

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"time"
)

const (
	AccessTokenUrl = "https://api.weixin.qq.com/cgi-bin/token"
)
type Context struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
}

type ResAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	ErrCode 	int64  `json:"errcode"`
	ErrMsg  	string `json:"errmsg"`
}

func (ctx *Context) GetAccessToken() (accessToken string, err error)  {
	var resAccessToken ResAccessToken
	resAccessToken, err = ctx.GetAccessTokenByService()
	if err != nil {
		return
	}

	accessToken = resAccessToken.AccessToken
	return
}

func (ctx *Context) GetAccessTokenByService() (resAccessToken ResAccessToken, err error)  {
	uri := fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", AccessTokenUrl, ctx.AppID, ctx.AppSecret)

	body, err := http.Get(uri)

	if err != nil {
		return
	}

	content, err := ioutil.ReadAll(body.Body)

	if err != nil {
		return
	}

	err = json.Unmarshal(content, &resAccessToken)

	if err != nil {
		return
	}

	if resAccessToken.ErrMsg != "" {
		err = fmt.Errorf("get access_token error : errcode=%v , errormsg=%v", resAccessToken.ErrCode, resAccessToken.ErrMsg)
		return
	}
	return
}