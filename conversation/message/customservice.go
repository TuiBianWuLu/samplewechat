package message

import (
	"github.com/TuiBianWuLu/samplewechat/basic"
	"github.com/TuiBianWuLu/samplewechat/core/request"
	"fmt"
	"encoding/json"
)

const (
	AddAccount = "https://api.weixin.qq.com/customservice/kfaccount/add"
	UpdateAccount = "https://api.weixin.qq.com/customservice/kfaccount/update"
	DeleteAccount = "https://api.weixin.qq.com/customservice/kfaccount/delete"
	UploadHeadImg = "https://api.weixin.qq.com/customservice/kfaccount/uploadheadimg"
	GetKfList = "https://api.weixin.qq.com/cgi-bin/customservice/getkflist"
)
//a = add u = update d = delete
var a, u, d = 1, 2, 3

type Message struct {
	*basic.WeChat
}

type reqAccount struct {
	Account []*Account `json:"account"`
}

type Account struct {
	KfAccount	string 	`json:"kf_account"`
	NickName	string	`json:"nick_name"`
	Password	string	`json:"pass_word"`
}

type KfList struct {
	*basic.CommonError
	KfAccount		string	`json:"kf_account"`
	KfNick			string	`json:"kf_nick"`
	KfId			string	`json:"kf_id"`
	KfHeadImgUrl	string	`json:"kf_headimgurl"`
}

type Avatar struct {
	KfAccount	string	`json:"kf_account"`
}

//客服账号
func (m *Message) Account (account []*Account, apiType int) error {

	req := &reqAccount{account}

	res, err := m.RefreshToken()

	if err != nil {
		return  err
	}

	var uri string

	switch apiType {
		case a:
			uri = fmt.Sprintf("%s?access_token=%s", AddAccount, res.AccessToken)
		case u:
			uri = fmt.Sprintf("%s?access_token=%s", UpdateAccount, res.AccessToken)
		case d:
			uri = fmt.Sprintf("%s?access_token=%s", DeleteAccount, res.AccessToken)
	}

	body, err := request.Post(uri, req)

	if err != nil {
		return err
	}

	var result basic.CommonError

	err = json.Unmarshal(body, result)

	if err != nil {
		return err
	}

	if result.ErrCode != '0' {
		return fmt.Errorf("add/update service account fail : %s", result.ErrMsg)
	}
	return nil
}

//设置头像
func (m *Message) SetAvatar (avatar Avatar) error {
	req := &avatar

	res, err := m.RefreshToken()

	if err != nil {
		return  err
	}

	uri := fmt.Sprintf("%s?access_token=%s", UploadHeadImg, res.AccessToken)

	body, err := request.Post(uri, req)

	if err != nil {
		return err
	}

	var result basic.CommonError

	err = json.Unmarshal(body, result)

	if err != nil {
		return err
	}

	if result.ErrCode != '0' {
		return fmt.Errorf("set avatar fail : %s", result.ErrMsg)
	}
	return nil
}

func (m *Message) GetKfList () (KfList, error) {

	var req byte

	var resp KfList

	res, err := m.RefreshToken()

	if err != nil {
		return resp, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", GetKfList, res.AccessToken)

	body, err := request.Post(uri, req)

	if err != nil {
		return resp, err
	}

	var result KfList

	err = json.Unmarshal(body, result)

	if err != nil {
		return resp, err
	}

	if result.ErrCode != '0' {
		return resp, err
	}
	return result, nil
}