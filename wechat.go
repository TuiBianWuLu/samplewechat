package samplewechat

import (
    "github.com/TuiBianWuLu/samplewechat/config"
    "github.com/TuiBianWuLu/samplewechat/material"
    "github.com/TuiBianWuLu/samplewechat/menu"
    "github.com/TuiBianWuLu/samplewechat/token"
    "github.com/TuiBianWuLu/samplewechat/custom"
    "github.com/TuiBianWuLu/samplewechat/message"
    "github.com/TuiBianWuLu/samplewechat/template"
)

type WeChat struct {
    *config.Config
}

var w = new(WeChat)

func New(c *config.Config) *WeChat {
    w.Config = c
    return w
}

func (w *WeChat) AccessToken() *token.AccessToken {
    return token.NewAccessToken(w.Config)
}

func (w *WeChat) Material() *material.Material {
    return material.New(w.Config)
}

func (w *WeChat) Menu() *menu.Menu {
    return menu.New(w.Config)
}

func (w *WeChat) Custom() *custom.Custom {
    return custom.New(w.Config)
}

func (w *WeChat) Template() *template.Template {
    return template.New(w.Config)
}

func (w *WeChat) Message() *message.Message {
    return message.New(w.Config)
}


