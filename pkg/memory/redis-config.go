package memory

import (
	"github.com/go-ozzo/ozzo-config"
	"github.com/go-redis/redis"
	"github.com/qkgo/scaff/pkg/cfg"
	"strconv"
)


func InitializationStorageRedis(c *config.Config) {
	addr := c.GetString("db.redis.storage.host")
	cfg.Log.Info(addr)
	port := c.GetInt("db.redis.storage.port")
	cfg.Log.Info(port)
	password := c.GetString("db.redis.storage.password")
	selectDb := c.GetInt("db.redis.storage.select")
	cfg.Log.Info(selectDb)
	portInt := strconv.Itoa(port)
	StorageRedisClient = *redis.NewClient(&redis.Options{
		Addr:     addr + ":" + portInt,
		Password: password,
		DB:       selectDb,
	})
}
