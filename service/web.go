package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"matchingServer/common"
	"matchingServer/request"
)

func StartServer() {
	engine := gin.Default()
	server := engine

	server.POST(common.RequestApi, func(context *gin.Context) {
		message := request.Message{}
		getOrder(message)
		err := context.BindJSON(&message)
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	server.Run(":8081")
}
