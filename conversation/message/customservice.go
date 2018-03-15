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
	UploadHeadImg = "http://api.weixin.qq.com/customservice/kfaccount/uploadheadimg"
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

type Avatar struct {
	KfAccount	string	`json:"kf_account"`
}

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