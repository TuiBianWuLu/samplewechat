package menu

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
    Sex                int64  `json:"sex,omitempty"`
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

type QueryMenuButton struct {
    Menus struct {
        CreateMenuButton
        MenuID int64 `json:"menuid,omitempty"`
    } `json:"menu,omitempty"`
}

type TryMatch struct {
    UserID      string  `json:"user_id"`
}

type QueryTryMatch struct {
    Menus struct {
        Buttons []Button `json:"button"`
    } `json:"menu,omitempty"`
}
