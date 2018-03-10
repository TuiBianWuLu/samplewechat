package basic

import (
    "net/http"
    "sync"

    "github.com/TuiBianWuLu/samplewechat/core/interfaces"
)

type Wechat struct {
    AppId  string
    Secret string

    Token  string
    AESKey string

    MchId  string
    MchKey string

    Request  http.Request
    Response *http.ResponseWriter

    accessTokenLock sync.Mutex
    jsApiTicketLock sync.Mutex

    cache interfaces.Cacher
}
