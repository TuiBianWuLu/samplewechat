package menu

import (
    "encoding/json"
    "fmt"

    "github.com/TuiBianWuLu/samplewechat/config"
    "github.com/TuiBianWuLu/samplewechat/token"
    "github.com/TuiBianWuLu/samplewechat/util/request"
    "github.com/TuiBianWuLu/samplewechat/util/response"
)

const (
    ADDMENUURL   = "https://api.weixin.qq.com/cgi-bin/menu/create"
    QUERYMENUURL = "https://api.weixin.qq.com/cgi-bin/menu/get"
    DELMENUURL   = "https://api.weixin.qq.com/cgi-bin/menu/delete"
)

type Menu struct {
    *config.Config
}

var m = new(Menu)

func New(c *config.Config) *Menu {
    m.Config = c
    return m
}

func (m *Menu) CreateMenu(menu CreateMenuButton) (resMsg response.CommonError, err error) {

    accessToken, err := token.NewAccessToken(m.Config).AccessToken()

    if err != nil {
        return
    }

    url := fmt.Sprintf("%s?access_token=%s", ADDMENUURL, accessToken)

    res, err := request.Post(url, menu)

    if err != nil {
        return
    }

    err = json.Unmarshal(res, &resMsg)

    return
}

func (m *Menu) QueryMenu() (queryMenuButton QueryMenuButton, err error) {

    accessToken, err := token.NewAccessToken(m.Config).AccessToken()

    if err != nil {
        return
    }

    url := fmt.Sprintf("%s?access_token=%s", QUERYMENUURL, accessToken)

    res, err := request.Get(url)

    if err != nil {
        return
    }

    err = json.Unmarshal(res, &queryMenuButton)

    return
}

func (m *Menu) DelMenu() (resMsg response.CommonError, err error) {

    accessToken, err := token.NewAccessToken(m.Config).AccessToken()

    if err != nil {
        return
    }

    url := fmt.Sprintf("%s?access_token=%s", DELMENUURL, accessToken)

    res, err := request.Get(url)

    if err != nil {
        return
    }

    err = json.Unmarshal(res, &resMsg)

    return
}
