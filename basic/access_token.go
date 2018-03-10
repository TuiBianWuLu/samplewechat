package basic

import (
    "encoding/json"
    "fmt"

    "github.com/TuiBianWuLu/samplewechat/core/request"
)

const refreshTokenUrl = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential"

func (w *Wechat) GetAccessToken(key string) string {
    return w.cache.Get(key)
}

func (w *Wechat) SetAccessToken(key, value string) bool {
    return w.cache.Set(key, value)
}

func (w *Wechat) RefreshToken() (responseToken ResponseAccessToken, err error) {

    uri := fmt.Sprintf("%s&appid=%s&secret=%s", refreshTokenUrl, w.AppId, w.Secret)

    body, err := request.Get(uri)

    if err != nil {
        return
    }

    err = json.Unmarshal(body, &responseToken)

    return
}
