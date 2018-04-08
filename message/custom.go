package message

type Basic struct {
	ToUser	string `json:"touser"`
	MsgType	string `json:"msgtype"`
}

type Info struct {
	Content			string	`json:"content,omitempty"`
	MediaID			string	`json:"media_id,omitempty"`
	ThumbMediaID	string	`json:"thumb_media_id,omitempty"`
	Title			string	`json:"title,omitempty"`
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
}

type Image struct {
	Basic
	Image	Info	`json:"image"`
}

type Voice struct {
	Basic
	Voice	Info	`json:"voice"`
}

type Video struct {
	Basic
	Video	Info	`json:"video"`
}

type Music struct {
	Basic
	Music	Info	`json:"music"`
}

type News struct {
	Basic
	News 	Info	`json:"news"`
}


