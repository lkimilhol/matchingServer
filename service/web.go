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
		r := request.Message{}

		err := context.BindJSON(&r)
		if err != nil {
			fmt.Println(err)
			fmt.Println("json parsing error")
			return
		}
		fmt.Println(r)
	})

	r.Run(":8081")
}
