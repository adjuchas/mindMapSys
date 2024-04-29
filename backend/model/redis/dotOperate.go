package redisConn

import (
	"context"
	"fmt"
)

func SetStarDot(stuId string, dotId string) bool {
	rdb := RedisConn.Conn()
	defer func() {
		rdb.Close()
	}()
	result := rdb.SAdd(context.Background(), stuId+"dot", dotId)
	_, err := result.Result()
	if err != nil {
		fmt.Println("set star dot err ")
		fmt.Println(err)
		return false
	}
	return true
}

func GetStarDots(stuId string) []string {
	rdb := RedisConn.Conn()
	defer func() {
		rdb.Close()
	}()
	result := rdb.SMembers(context.Background(), stuId+"dot")
	_, err := result.Result()
	if err != nil {
		fmt.Println("set start dot err ")
		fmt.Println(err)
		return nil
	}
	return result.Val()
}

func CheckDotStat(stuId string, dotId string) bool {
	rdb := RedisConn.Conn()
	defer func() {
		rdb.Close()
	}()
	result := rdb.SIsMember(context.Background(), stuId+"dot", dotId)
	return result.Val()
}

func QuitDotStat(stuId string, dotId string) bool {
	rdb := RedisConn.Conn()
	defer func() {
		rdb.Close()
	}()
	result := rdb.SRem(context.Background(), stuId+"dot", dotId)
	_, err := result.Result()
	if err != nil {
		fmt.Println("set star dot err ")
		fmt.Println(err)
		return false
	}
	return true
}
