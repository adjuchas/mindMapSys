package redisConn

import (
	"backend/conf"
	"github.com/redis/go-redis/v9"
)

var RedisConn *redis.Client

func InitConnRedis(conf conf.ServerConfig) {
	RedisConn = redis.NewClient(&redis.Options{
		Addr:     conf.REDISDB.HOST,
		Password: conf.REDISDB.PASSWORD,
		DB:       0,
	})
}
