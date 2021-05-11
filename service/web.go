package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"matchingServer/common"
	"matchingServer/request"
)

func StartServer() {
	engine := gin.Default()
	r := engine

	r.POST(common.RequestApi, func(context *gin.Context) {
		context.BindJSON(request.Message{})
		fmt.Println(request.Message{})
	})

	r.Run(":8081")
}
