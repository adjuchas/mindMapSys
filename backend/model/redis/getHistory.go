package redisConn

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func GetHistory(stuId string) ([]string, error) {
	rdb := RedisConn.Conn()
	defer func() {
		rdb.Close()
	}()
	val, err := rdb.LRange(context.Background(), stuId+"history", int64(0), int64(-1)).Result()
	if err == redis.Nil || err != nil {
		return []string{}, redis.Nil
	}
	return val, nil
}

func SetHistory(stuId string, dotId int) {
	rdb := RedisConn.Conn()
	defer func() {
		rdb.Close()
	}()
	rdb.LPush(context.Background(), stuId+"history", string(dotId))
}
