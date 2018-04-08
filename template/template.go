package template

import (
	"github.com/TuiBianWuLu/samplewechat/config"
	"github.com/TuiBianWuLu/samplewechat/token"
	"fmt"
	"github.com/TuiBianWuLu/samplewechat/util/request"
	"encoding/json"
	"github.com/TuiBianWuLu/samplewechat/util/response"
)

const (
	SendTemplateUrl		= "https://api.weixin.qq.com/cgi-bin/message/template/send"
	SetIndustryUrl		= "https://api.weixin.qq.com/cgi-bin/template/api_set_industry"
	GetIndustryUrl		= "https://api.weixin.qq.com/cgi-bin/template/get_industry"
	GetTmpIDUrl			= "https://api.weixin.qq.com/cgi-bin/template/api_add_template"
	GetTmpListUrl		= "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template"
	DelTemplateUrl		= "https://api.weixin.qq.com/cgi-bin/template/del_private_template"
)

type Template struct {
	*config.Config
}

var m = new(Template)

func New(c *config.Config) *Template {
	m.Config = c
	return m
}

func (t *Template) Send(content Template) (sendTemplateRes SendTemplateRes, err error) {
	accessToken, err := token.NewAccessToken(t.Config).AccessToken()

	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", SendTemplateUrl, accessToken)

	res, err := request.Post(url, content)

	if err != nil {
		return
	}

	json.Unmarshal(res, &sendTemplateRes)

	return
}

func (t *Template) SetIndustry(industry Industry) (resMsg response.CommonError, err error) {
	accessToken, err := token.NewAccessToken(t.Config).AccessToken()

	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", SetIndustryUrl, accessToken)

	res, err := request.Post(url, industry)

	if err != nil {
		return
	}

	json.Unmarshal(res, &resMsg)

	return
}

func (t *Template) GetIndustry() (industryRes GetIndustryRes, err error) {
	accessToken, err := token.NewAccessToken(t.Config).AccessToken()

	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", GetIndustryUrl, accessToken)

	res, err := request.Get(url)

	if err != nil {
		return
	}

	json.Unmarshal(res, &industryRes)

	return
}

func (t *Template) GetTmpID(templateID GetTemplateID) (templateIdRes GetTemplateIdRes, err error) {
	accessToken, err := token.NewAccessToken(t.Config).AccessToken()

	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", GetTmpIDUrl, accessToken)

	res, err := request.Post(url, templateID)

	if err != nil {
		return
	}

	json.Unmarshal(res, &templateIdRes)

	return
}

func (t *Template) GetTmpList() (tmpListRes GetTemplateListRes, err error) {
	accessToken, err := token.NewAccessToken(t.Config).AccessToken()

	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", GetTmpListUrl, accessToken)

	res, err := request.Get(url)

	if err != nil {
		return
	}

	json.Unmarshal(res, &tmpListRes)

	return
}

func (t *Template) DelTemplate(templateID DelTemplateID) (resMsg response.CommonError, err error) {
	accessToken, err := token.NewAccessToken(t.Config).AccessToken()

	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", DelTemplateUrl, accessToken)

	res, err := request.Post(url, templateID)

	if err != nil {
		return
	}

	json.Unmarshal(res, &resMsg)

	return
}