package template

import (
	"github.com/TuiBianWuLu/samplewechat/util/response"
)

type Tpl struct {
	ToUser		string				`json:"touser"`
	TemplateID	string				`json:"template_id"`
	Data		map[string]Item		`json:"data"`
	Url			string				`json:"url,omitempty"`
	MiniProgram	WeApp				`json:"miniprogram,omitempty"`
	Color		string				`json:"color,omitempty"`
}

type WeApp struct {
	AppID		string	`json:"appid"`
	PagePath	string	`json:"pagepath"`
}

type SendTemplateRes struct {
	response.CommonError
	MsgID		int64	`json:"msgid"`
}

type Item struct {
	Value		string	`json:"value"`
	Color 		string	`json:"color,omitempty"`
}

type Industry struct {
	IndustryID1	int		`json:"industry_id1"`
	IndustryID2	int		`json:"industry_id2"`
}

type GetTemplateID struct {
	TemplateIDShort	string	`json:"template_id_short"`
}

type DelTemplateID struct {
	TemplateID		string	`json:"template_id"`
}

type GetTemplateListRes struct {
	TemplateList	[]ListDetailRes	`json:"template_list"`
}

type ListDetailRes struct {
	TemplateID			string	`json:"template_id"`
	Title				string	`json:"title"`
	PrimaryIndustry		string	`json:"primary_industry"`
	DeputyIndustry		string	`json:"deputy_industry"`
	Content				string	`json:"content"`
	Example				string	`json:"example"`
}

type GetIndustryRes struct {
	PrimaryIndustry		IndustryDetail	`json:"primary_industry"`
	SecondaryIndustry	IndustryDetail	`json:"secondary_industry"`
}

type GetTemplateIdRes struct {
	response.CommonError
	TemplateID	string	`json:"template_id"`
}

type IndustryDetail struct {
	FirstClass	string	`json:"first_class"`
	SecondClass	string	`json:"second_class"`
}