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
    AddMenuUrl          = "https://api.weixin.qq.com/cgi-bin/menu/create"
    QueryMenuUrl        = "https://api.weixin.qq.com/cgi-bin/menu/get"
    DelMenuUrl          = "https://api.weixin.qq.com/cgi-bin/menu/delete"
    AddConditionalUrl   = "https://api.weixin.qq.com/cgi-bin/menu/addconditional"
    TryMatchUrl         = "https://api.weixin.qq.com/cgi-bin/menu/trymatch"
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

    url := fmt.Sprintf("%s?access_token=%s", AddMenuUrl, accessToken)

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

    url := fmt.Sprintf("%s?access_token=%s", QueryMenuUrl, accessToken)

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

    url := fmt.Sprintf("%s?access_token=%s", DelMenuUrl, accessToken)

    res, err := request.Get(url)

    if err != nil {
        return
    }

    err = json.Unmarshal(res, &resMsg)

    return
}

func (m *Menu) AddConditional(menu CreateMenuButton) (resMsg response.CommonError, err error) {
    accessToken, err := token.NewAccessToken(m.Config).AccessToken()

    if err != nil {
        return
    }

    url := fmt.Sprintf("%s?access_token=%s", AddConditionalUrl, accessToken)

    res, err := request.Post(url, menu)

    if err != nil {
        return
    }

    err = json.Unmarshal(res, &resMsg)

    return
}

func (m *Menu) TestTryMatch(menu TryMatch) (queryTryMatch QueryTryMatch, err error) {
    accessToken, err := token.NewAccessToken(m.Config).AccessToken()

    if err != nil {
        return
    }

    url := fmt.Sprintf("%s?access_token=%s", TryMatchUrl, accessToken)

    res, err := request.Post(url, menu)

    if err != nil {
        return
    }

    err = json.Unmarshal(res, &queryTryMatch)

    return
}
