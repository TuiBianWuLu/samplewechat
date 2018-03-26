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
    ADDMENUURL = "https://api.weixin.qq.com/cgi-bin/menu/create"
)

type Menu struct {
    *config.Config
}

type Item struct {
    Type      string `json:"type"`
    Name      string `json:"name"`
    Url       string `json:"url,omitempty"`
    MediaId   string `json:"media_id,omitempty"`
    Appid     string `json:"appid,omitempty"`
    Key       string `json:"key,omitempty"`
    PagePath  string `json:"pagepath,omitempty"`
    SubButton []Item `json:"sub_button,omitempty"`
}

var m = new(Menu)

func New(c *config.Config) *Menu {
    m.Config = c
    return m
}

func (m *Menu) CreateMenu(menu []Item) (resMsg response.CommonError, err error) {

    accessToken, err := token.NewAccessToken(m.Config).AccessToken()

    if err != nil {
        return
    }

    url := fmt.Sprintf("%s?access_token=%s", ADDMENUURL, accessToken)

    menuList := struct {
        Button []Item `json:"button"`
    }{
        Button: menu,
    }

    res, err := request.Post(url, menuList)

    if err != nil {
        return
    }

    err = json.Unmarshal(res, &resMsg)

    return
}
