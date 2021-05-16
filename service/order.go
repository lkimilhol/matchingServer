package service

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"matchingServer/domain"
	"matchingServer/request"
)

func getOrder(message request.Message) {
	conn := getRedisConn()
	defer conn.Close()

	reply, err := redis.Values(conn.Do("hgetall", "com.lkimilhol.matchingProject.domain.OrderSave:4940414585990594970"))

	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
	}
	fmt.Println("complete")
	orderSave := domain.OrderSave{}
	err = redis.ScanStruct(reply, &orderSave)

	fmt.Printf("%+v", orderSave)
}
