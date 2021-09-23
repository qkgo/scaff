package serialize

import (
	config "github.com/go-ozzo/ozzo-config"
	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
	"github.com/jinzhu/gorm"
//	"github.com/mongodb/mongo-go-driver/mongo"
	"strconv"
)

var DB *gorm.DB

var SecondDB *gorm.DB

// var MonGoDB *mongo.Client

var XDB *xorm.Engine

var SecondXDB *xorm.Engine

var ChannelRcsClient redis.Client

func InitializationRcsRedis(c *config.Config) {
	addr := c.GetString("db.redis.rcs.host")
	port := c.GetInt("db.redis.rcs.port")
	password := c.GetString("db.redis.rcs.password")
	selectDb := c.GetInt("db.redis.rcs.select")
	portInt := strconv.Itoa(port)
	ChannelRcsClient = *redis.NewClient(&redis.Options{
		Addr:      addr + ":" + portInt,
		Password:  password,
		DB:        selectDb,
		OnConnect: redisConnectEvent,
	})
}

func redisConnectEvent(cnn *redis.Conn) error {
	println("[Redis] connected. ", cnn)
	return nil
}
