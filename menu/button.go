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

type CreateMenuButton struct {
    Buttons []Button `json:"button"`
}

type QueryMenuButton struct {
    Menus struct{
        CreateMenuButton
        MenuID int64 `json:"menuid,omitempty"`
    } `json:"menu,omitempty"`
}