package redisConn

import (
	"context"
	"fmt"
	"reflect"
)

func SetStarDot(stuId string, dotId string) bool {
	rdb := RedisConn.Conn()
	defer func() {
		fmt.Println("is dafer")
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
		fmt.Println("is dafer")
		rdb.Close()
	}()
	result := rdb.SMembers(context.Background(), stuId+"dot")
	fmt.Println("is GetStarDotsInfo start")
	fmt.Println(result)
	fmt.Println(result.Val())
	fmt.Println(reflect.TypeOf(result))
	fmt.Println("is GetStartDotsInfo end")
	_, err := result.Result()
	if err != nil {
		fmt.Println("set start dot err ")
		fmt.Println(err)
		return nil
	}
	return result.Val()
}
