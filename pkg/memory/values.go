package memory

import (
	"github.com/go-redis/redis"
)

var StorageRedisClient redis.Client

type TopicData struct {
	productId string
	roomId    string
	event     string
	cmd       string
	body      interface{}
	uuid      string
}
