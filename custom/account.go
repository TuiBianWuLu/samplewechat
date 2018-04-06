package custom

import "github.com/TuiBianWuLu/samplewechat/util/response"

type AddKfAccount struct {
	KfAccount	string	`json:"kf_account,omitempty"`
	NickName	string 	`json:"nickname,omitempty"`
	InviteWX	string	`json:"invite_wx,omitempty"`
}

type ListBasic struct {
	KfAccount	string `json:"kf_account"`
	KfNick		string `json:"kf_nick"`
	KfID		string `json:"kf_id"`
	KfHeadImgUrl	string `json:"kf_headimgurl"`
	KfWX		string `json:"kf_wx"`
}

type KfListRes struct {
	response.CommonError
	KfList	[]ListBasic `json:"kf_list"`
}

type DelAccount struct {
	KfAccount	string	`json:"kf_account"`
}

type OnlineKfListRes struct {
	KfOnlineList	[]OnlineKfList	`json:"kf_online_list"`
}

type OnlineKfList struct {
	KfAccount	string	`json:"kf_account"`
	Status		int		`json:"status"`
	KfID		int		`json:"ke_id"`
	AcceptedCase	string	`json:"accepted_case"`
}