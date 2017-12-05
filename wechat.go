package main

import (
	"github.com/gin-gonic/gin"
	"weChatSdk/basic"
)

func main()  {

	router := gin.Default()

	router.HEAD("/", func(c *gin.Context) {
		c.Status(200)
	})

	router.GET("/", func(c *gin.Context) {
		c.Status(200)
	})

	basicRouter := router.Group("/basic")
	{
		basicRouter.GET("/access_token", basic.Token)
	}

	router.Run(":5213")
}


