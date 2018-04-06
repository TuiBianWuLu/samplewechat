package menu

import "github.com/TuiBianWuLu/samplewechat/util/response"

type MenuID int64

type Button struct {
    Type      string   `json:"type"`
    Name      string   `json:"name"`
    Url       string   `json:"url,omitempty"`
    MediaID   string   `json:"media_id,omitempty"`
    AppID     string   `json:"appid,omitempty"`
    Key       string   `json:"key,omitempty"`
    PagePath  string   `json:"pagepath,omitempty"`
    SubButton []Button `json:"sub_button,omitempty"`
}

type MatchRule struct {
    TagID              int64  `json:"tag_id,omitempty"`
    Sex                string `json:"sex,omitempty"`
    ClientPlatformType int64  `json:"client_platform_type,omitempty"`
    Country            string `json:"country,omitempty"`
    Province           string `json:"province,omitempty"`
    City               string `json:"city,omitempty"`
    Language           string `json:"language,omitempty"`
}

type CreateMenuButton struct {
    Buttons []Button `json:"button"`
    MatchRule        `json:"matchrule,omitempty"`
}

type CreateMenuRes struct {
    response.CommonError
    MenuID `json:"menuid"`
}

type QueryMenuButtonRes struct {
    response.CommonError
    Menus struct {
        Buttons []Button `json:"button"`
        MenuID           `json:"menuid"`
    } `json:"menu"`
    ConditionalMenu []struct {
        Buttons []Button `json:"button"`
        MenuID           `json:"menuid"`
        MatchRule        `json:"matchrule"`
    } `json:"conditionalmenu"`
}

type TryMatchUser struct {
    UserId string `json:"user_id"`
}

type TryMatchRes struct {
    response.CommonError
    Menus struct{
        Buttons []Button `json:"button"`
    } `json:"menu"`
}

type DelMenu struct {
    MenuID `json:"menuid"`
}
