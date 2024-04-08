package redisConn

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

func CheckUserByStuid(accessToken string, stuId string) bool {
	info, err := GetInfoMsg(accessToken)
	if err == redis.Nil || info.Yb_studentid != stuId || err != nil {
		fmt.Println(err)
		fmt.Println(info.Yb_studentid)
		fmt.Println(stuId)
		return false
	}
	return true
}

func CheckTeaUserByStuid(accessToken string, Id string) bool {
	info, err := GetInfoMsg(accessToken)
	if err == redis.Nil || info.Yb_employid != Id || err != nil {
		return false
	}
	return true
}
