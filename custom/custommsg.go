package custom

type Basic struct {
	ToUser	string `json:"touser"`
	MsgType	string `json:"msgtype"`
}

type CService struct {
	KfAccount	string	`json:"kf_account,omitempty"`
}

type Info struct {
	Content			string	`json:"content,omitempty"`
	MediaID			string	`json:"media_id,omitempty"`
	ThumbMediaID	string	`json:"thumb_media_id,omitempty"`
	Title			string	`json:"title,omitempty"`
	AppID			string	`json:"appid,omitempty"`
	PagePath		string	`json:"pagepath,omitempty"`
	Url				string	`json:"url,omitempty"`
	Description		string	`json:"description,omitempty"`
	MusicUrl		string	`json:"musicurl,omitempty"`
	HqMusicUrl		string	`json:"hqmusicurl,omitempty"`
	PicUrl			string	`json:"picurl,omitempty"`
	CardID			string	`json:"card_id,omitempty"`
	Articles		[]Info	`json:"articles,omitempty"`
}

type Text struct {
	Basic
	Text	Info	`json:"text"`
	CustomService	CService	`json:"customservice,omitempty"`
}

type Image struct {
	Basic
	Image	Info	`json:"image"`
	CustomService	CService	`json:"customservice,omitempty"`
}

type Voice struct {
	Basic
	Voice	Info	`json:"voice"`
	CustomService	CService	`json:"customservice,omitempty"`
}

type Video struct {
	Basic
	Video	Info	`json:"video"`
	CustomService	CService	`json:"customservice,omitempty"`
}

type Music struct {
	Basic
	Music	Info	`json:"music"`
	CustomService	CService	`json:"customservice,omitempty"`
}

type News struct {
	Basic
	News 	Info	`json:"news"`
	CustomService	CService	`json:"customservice,omitempty"`
}

type MpNews struct {
	Basic
	MpNews	Info	`json:"mpnews"`
	CustomService	CService	`json:"customservice,omitempty"`
}

type WxCard struct {
	Basic
	WxCard	Info	`json:"wxcard"`
	CustomService	CService	`json:"customservice,omitempty"`
}