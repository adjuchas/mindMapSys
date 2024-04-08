package redisConn

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func GetHistory(stuId string) ([]string, error) {
	rdb := RedisConn.Conn()
	defer func() {
		fmt.Println("is dafer")
		rdb.Close()
	}()
	val, err := rdb.LRange(context.Background(), stuId, int64(0), int64(-1)).Result()
	if err == redis.Nil || err != nil {
		return []string{}, redis.Nil
	}
	return val, nil
}
