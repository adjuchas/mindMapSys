package redisConn

import (
	"backend/tools"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type TeaInfo struct {
	Name       string `json:"name"`
	KsID       string `json:"ks_id"`
	Teacher_ID string `json:"tea_id"`
	Identity   string `json:"identity"`
	Pic        string `json:"pic"`
	Yb_ID      string `json:"yb_ID"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

func GetTeaInfoMsg(accessToken string) (TeaInfo, error) {
	rdb := RedisConn.Conn()
	defer func() {
		rdb.Close()
	}()
	val, err := rdb.HVals(context.Background(), accessToken).Result()
	if err != nil || len(val) == 0 {
		return TeaInfo{}, redis.Nil
	}
	valMap := make(map[string]string)
	byteSlice := make([]byte, 0)
	for _, str := range val {
		byteSlice = append(byteSlice, []byte(str)...)
	}
	json.Unmarshal(byteSlice, &valMap)
	var info TeaInfo
	err = tools.MapToStruct(valMap, &info)

	if err != nil {
		fmt.Println(err)
	}
	return info, nil
}
