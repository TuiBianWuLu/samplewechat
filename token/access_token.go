package token

import (
    "encoding/json"
    "errors"
    "fmt"
    "time"

    "github.com/TuiBianWuLu/samplewechat/config"
    "github.com/TuiBianWuLu/samplewechat/util/request"
    "github.com/TuiBianWuLu/samplewechat/util/response"
)

const (
    ACCESSTOKENURL = "https://api.weixin.qq.com/cgi-bin/token"
)

type AccessToken struct {
    *config.Config
}

type AccessTokenRes struct {
    response.CommonError
    AccessToken string `json:"access_token"`
    ExpiresIn   int64  `json:"expires_in"`
}

var a = new(AccessToken)

func NewAccessToken(c *config.Config) *AccessToken {
    a.Config = c
    return a
}

func (a *AccessToken) refreshToken() (accessTokenRes AccessTokenRes, err error) {

    url := fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", ACCESSTOKENURL, a.AppID, a.Secret)

    res, err := request.Get(url)

    if err != nil {
        return
    }

    err = json.Unmarshal(res, &accessTokenRes)

    return
}

func (a *AccessToken) AccessToken() (token string, err error) {

    a.AccessTokenLock.Lock()
    defer a.AccessTokenLock.Unlock()

    var key = fmt.Sprintf("access_token_%s", a.AppID)

    if a.PrefixAccessTokenKey != "" {
        key = a.PrefixAccessTokenKey
    }

    isExists, err := a.Cache.IsExists(key)

    if err != nil {
        return
    }

    // 自动刷token
    if !isExists && a.PrefixAccessTokenKey == "" {
        res, err := a.refreshToken()

        if err != nil {
            return "", err
        }

        if err != nil || res.ErrorCode > 0 {
            err = errors.New(fmt.Sprintf("%d %s", res.ErrorCode, res.ErrorMessage))
            return "", err
        }

        if _, err = a.Cache.Set(key, res.AccessToken, time.Second * 300); err != nil {
            return "", err
        }
    }

    token, err = a.Cache.Get(key)

    return
}
