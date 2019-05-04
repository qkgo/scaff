package cfg

import (
	"fmt"
	"log"
)

func SetRabbitServer() {
	if ConfigParam == nil {
		return
	}
	userName := ConfigParam.Get("rabbit.username")
	pass := ConfigParam.Get("rabbit.password")
	host := ConfigParam.Get("rabbit.host")
	port := ConfigParam.Get("rabbit.port")
	log.Println(userName, pass,
		host,
		port)
	if userName == nil || host == nil || port == nil {
		return
	}
	if userName.(string) != "" && host.(string) != "" && port.(int) > 100 {
		RabbitServerUrl = fmt.Sprintf("amqp://%s:%s@%s:%d/",
			userName,
			pass,
			host,
			port)
	}
}
