package samplewechat

import (
    "github.com/TuiBianWuLu/samplewechat/config"
    "github.com/TuiBianWuLu/samplewechat/token"
)

type Wechat struct {
    *config.Config
}

var w = new(Wechat)

func New(c *config.Config) *Wechat {
    w.Config = c
    return w
}

func (w *Wechat) AccessToken() *token.AccessToken {
    return token.NewAccessToken(w.Config)
}
