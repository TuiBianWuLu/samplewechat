package parmas

type (
	GetToken struct {
		AppId	string	`form:"app_id" json:"app_id" binding:"required"`
		Secret	string	`form:"secret" json:"secret" binding:"required"`
	}
)
