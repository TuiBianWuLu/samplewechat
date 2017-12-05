package basic

import (
	"github.com/gin-gonic/gin"
	"samplewechat/library/params"
	"fmt"
	"log"
)

func Token(c *gin.Context) {
	param := new (params.GetToken)

	if c.Bind(param) != nil {
		log.Println("编译出错")
	}

	fmt.Print(param)
}