package message

import (
	"github.com/TuiBianWuLu/samplewechat/config"
)

const (
	SendCustomMsgUrl	= "https://api.weixin.qq.com/cgi-bin/message/custom/send"
	SendTemplateUrl		= "https://api.weixin.qq.com/cgi-bin/message/template/send"
)

type Message struct {
	*config.Config
}

var m = new(Message)

func New(c *config.Config) *Message {
	m.Config = c
	return m
}
