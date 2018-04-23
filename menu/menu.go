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
    AddMenuUrl        = "https://api.weixin.qq.com/cgi-bin/menu/create"
    QueryMenuUrl      = "https://api.weixin.qq.com/cgi-bin/menu/get"
    DelMenuUrl        = "https://api.weixin.qq.com/cgi-bin/menu/delete"
    AddConditionalUrl = "https://api.weixin.qq.com/cgi-bin/menu/addconditional"
    TryMatchUrl       = "https://api.weixin.qq.com/cgi-bin/menu/trymatch"
    DelConditional    = "https://api.weixin.qq.com/cgi-bin/menu/delconditional"
)

type Menu struct {
    *config.Config
}

var m = new(Menu)

func New(c *config.Config) *Menu {
    m.Config = c
    return m
}

func (m *Menu) CreateMenu(createMenuButton CreateMenuButton) (createMenuRes CreateMenuRes, err error) {

    accessToken, err := token.NewAccessToken(m.Config).AccessToken()

    if err != nil {
        return
    }

    var url string
    if createMenuButton.MatchRule == (MatchRule{}) {
        url = fmt.Sprintf("%s?access_token=%s", AddMenuUrl, accessToken)
    } else {
        url = fmt.Sprintf("%s?access_token=%s", AddConditionalUrl, accessToken)
    }

    res, err := request.Post(url, createMenuButton)

    if err != nil {
        return
    }

    err = json.Unmarshal(res, &createMenuRes)

    return
}

func (m *Menu) QueryMenu() (queryMenuButtonRes QueryMenuButtonRes, err error) {

    accessToken, err := token.NewAccessToken(m.Config).AccessToken()

    if err != nil {
        return
    }

    url := fmt.Sprintf("%s?access_token=%s", QueryMenuUrl, accessToken)

    res, err := request.Get(url)

    if err != nil {
        return
    }

    fmt.Println(string(res))
    err = json.Unmarshal(res, &queryMenuButtonRes)

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

func (m *Menu) TryMatch(tryMatchUser TryMatchUser) (tryMatchRes TryMatchRes, err error) {

    accessToken, err := token.NewAccessToken(m.Config).AccessToken()

    if err != nil {
        return
    }

    url := fmt.Sprintf("%s?access_token=%s", TryMatchUrl, accessToken)

    res, err := request.Post(url, tryMatchUser)

    if err != nil {
        return
    }

    err = json.Unmarshal(res, &tryMatchRes)

    return
}


func (m *Menu) DelConditional (delMenu DelMenu) (resMsg response.CommonError, err error) {

    accessToken, err := token.NewAccessToken(m.Config).AccessToken()

    if err != nil {
        return
    }
    url := fmt.Sprintf("%s?access_token=%s", DelConditional, accessToken)

    res, err := request.Post(url, delMenu)

    if err != nil {
        return
    }

    err = json.Unmarshal(res, &resMsg)

    return
}