package message

import (
	"github.com/TuiBianWuLu/samplewechat/config"
	"github.com/TuiBianWuLu/samplewechat/util/response"
	"github.com/TuiBianWuLu/samplewechat/token"
	"fmt"
	"github.com/TuiBianWuLu/samplewechat/util/request"
	"encoding/json"
)

const (
	SendCustomMsgUrl = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
)

type Message struct {
	*config.Config
}

var m = new(Message)

func New(c *config.Config) *Message {
	m.Config = c
	return m
}

func (m *Message) SendCustom(content interface{}) (resMsg response.CommonError, err error) {

	accessToken, err := token.NewAccessToken(m.Config).AccessToken()

	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", SendCustomMsgUrl, accessToken)

	res, err := request.Post(url, content)

	if err != nil {
		return
	}

	json.Unmarshal(res, &resMsg)

	return
}