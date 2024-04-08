package redisConn

import (
	"backend/model/models"
	"context"
	"encoding/json"
	"fmt"
)

func RedisSetUserInfo(token string, info models.RedisUserInfo) {
	rdb := RedisConn.Conn()
	rval, _ := json.Marshal(info)
	result := rdb.HSet(context.Background(), token, "info", rval)
	_, err := result.Result()
	if err != nil {
		fmt.Println(err)
	}
}
