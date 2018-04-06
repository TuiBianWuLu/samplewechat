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
    ADDMENUURL        = "https://api.weixin.qq.com/cgi-bin/menu/create"
    ADDCONDITIONALURL = "https://api.weixin.qq.com/cgi-bin/menu/addconditional"
    QUERYMENUURL      = "https://api.weixin.qq.com/cgi-bin/menu/get"
    DELMENUURL        = "https://api.weixin.qq.com/cgi-bin/menu/delete"
    TRYMATCHURL       = "https://api.weixin.qq.com/cgi-bin/menu/trymatch"
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

    var url string
    if menu.MatchRule == (MatchRule{}) {
        url = fmt.Sprintf("%s?access_token=%s", ADDMENUURL, accessToken)
    } else {
        url = fmt.Sprintf("%s?access_token=%s", ADDCONDITIONALURL, accessToken)
    }

    res, err := request.Post(url, menu)

    if err != nil {
        return
    }

    err = json.Unmarshal(res, &resMsg)

    return
}

func (m *Menu) QueryMenu() (queryMenuButton QueryMenuButtonRes, err error) {

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

func (m *Menu) TryMatch(user TryMatchUser) (tryMatchRes TryMatchRes, err error) {

    accessToken, err := token.NewAccessToken(m.Config).AccessToken()

    if err != nil {
        return
    }

    url := fmt.Sprintf("%s?access_token=%s", TRYMATCHURL, accessToken)

    res, err := request.Post(url, user)

    if err != nil {
        return
    }

    fmt.Println(string(res))
    err = json.Unmarshal(res, &tryMatchRes)

    return
}
