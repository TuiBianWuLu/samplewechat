package custom

import (
	"github.com/TuiBianWuLu/samplewechat/config"
	"github.com/TuiBianWuLu/samplewechat/token"
	"fmt"
	"github.com/TuiBianWuLu/samplewechat/util/request"
	"github.com/TuiBianWuLu/samplewechat/util/response"
	"encoding/json"
)

const (
	AddAccountUrl 		= "https://api.weixin.qq.com/customservice/kfaccount/add"
	UpdateAccountUrl	= "https://api.weixin.qq.com/customservice/kfaccount/update"
	DelAccountUrl 		= "https://api.weixin.qq.com/customservice/kfaccount/del"
	InviteWorkerUrl		= "https://api.weixin.qq.com/customservice/kfaccount/inviteworker"
	KfListUrl			= "https://api.weixin.qq.com/cgi-bin/customservice/getkflist"
	GetOnlineKfListUrl	= "https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist"

	IsUpdate = 1
)

type Custom struct {
	*config.Config
}

var kf = new(Custom)

func New(c *config.Config) *Custom {
	kf.Config = c
	return kf
}

func (kf *Custom) AddAccount(account AddKfAccount, isUp int) (resMsg response.CommonError, err error) {

	accessToken, err := token.NewAccessToken(kf.Config).AccessToken()

	if err != nil {
		return
	}

	var url string
	if isUp == IsUpdate {
		url = fmt.Sprintf("%s?access_token=%s", UpdateAccountUrl, accessToken)
	} else {
		url = fmt.Sprintf("%s?access_token=%s", AddAccountUrl, accessToken)
	}

	res, err := request.Post(url, account)

	if err != nil {
		return
	}

	json.Unmarshal(res, &resMsg)

	return
}

func (kf *Custom) DelAccount(account string) (resMsg response.CommonError, err error) {
	accessToken, err := token.NewAccessToken(kf.Config).AccessToken()

	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s&kf_account=%s", DelAccountUrl, accessToken, account)


	res, err := request.Get(url)

	json.Unmarshal(res, &resMsg)

	return
}

func (kf *Custom) InviteWorker(account AddKfAccount) (resMsg response.CommonError, err error) {
	accessToken, err := token.NewAccessToken(kf.Config).AccessToken()

	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", InviteWorkerUrl, accessToken)


	res, err := request.Post(url, account)

	json.Unmarshal(res, &resMsg)

	return
}

func (kf *Custom) KfList() (kfList KfListRes, err error) {

	accessToken, err := token.NewAccessToken(kf.Config).AccessToken()

	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", KfListUrl, accessToken)

	res, err := request.Get(url)

	json.Unmarshal(res, &kfList)

	return
}

func (kf *Custom) GetOnlineKfList() (onlineKfList OnlineKfListRes, err error) {

	accessToken, err := token.NewAccessToken(kf.Config).AccessToken()

	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", GetOnlineKfListUrl, accessToken)

	res, err := request.Get(url)

	json.Unmarshal(res, &onlineKfList)

	return
}

