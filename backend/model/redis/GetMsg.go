package redisConn

import (
	"backend/model/models"
	"backend/tools"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func GetRedisUserInfo(accessToken string) (models.RedisUserInfo, error) {
	rdb := RedisConn.Conn()
	defer func() {
		rdb.Close()
	}()

	val, err := rdb.HVals(context.Background(), accessToken).Result()
	if err != nil || len(val) == 0 {
		return models.RedisUserInfo{}, redis.Nil
	}

	valMap := make(map[string]string)
	byteSlice := make([]byte, 0)
	for _, str := range val {
		byteSlice = append(byteSlice, []byte(str)...)
	}
	json.Unmarshal(byteSlice, &valMap)

	var info models.RedisUserInfo
	err = tools.MapToStruct(valMap, &info)
	if err != nil {
		fmt.Println(err)
	}
	return info, nil
}
