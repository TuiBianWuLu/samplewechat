package basic

import (
	"github.com/gin-gonic/gin"
	"weChatSdk/library/parmas"
	"fmt"
	"log"
)

func Token(c *gin.Context) {
	param := new (parmas.GetToken)

	if c.Bind(param) != nil {
		log.Println("编译出错")
	}

	fmt.Print(param)
}