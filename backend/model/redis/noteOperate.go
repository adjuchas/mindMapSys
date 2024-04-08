package redisConn

import (
	"context"
	"fmt"
	"reflect"
)

func SetStarNote(stuId string, noteId string) bool {
	rdb := RedisConn.Conn()
	defer func() {
		fmt.Println("is dafer")
		rdb.Close()
	}()
	result := rdb.SAdd(context.Background(), stuId+"note", noteId)
	_, err := result.Result()
	if err != nil {
		fmt.Println("set star note err ")
		fmt.Println(err)
		return false
	}
	return true
}

func GetStarNotes(stuId string) []string {
	rdb := RedisConn.Conn()
	defer func() {
		fmt.Println("is dafer")
		rdb.Close()
	}()
	result := rdb.SMembers(context.Background(), stuId+"note")
	fmt.Println("is GetStartDotsInfo start")
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
