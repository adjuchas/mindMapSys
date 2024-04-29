package service

import (
	mysqlConn "backend/model/mysql"
	redisConn "backend/model/redis"
	"backend/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

type ResponseRecommend struct {
	DotID        int       `json:"dot_id"`
	TITLE        string    `json:"title"`
	KsID         int       `json:"ks_id"`
	Auth         string    `json:"auth"`
	TAGS         string    `json:"tags"`
	DESCRIPTION  string    `json:"description"`
	UpdateTime   time.Time `json:"UpdateTime"`
	NodeTreePath string    `json:"nodeTreePath"`
}

func GetRecommends(c *gin.Context) {
	var data map[string]interface{}
	var requestInfo RequestInfo
	c.ShouldBindJSON(&requestInfo)
	_, err := redisConn.GetInfoMsg(requestInfo.AccessToken)
	if err == redis.Nil {
		c.AsciiJSON(http.StatusOK, data)
		return
	}
	recommends := mysqlConn.GetRecommends()
	var respRecommends []ResponseRecommend
	for _, item := range recommends {
		var respRecommend ResponseRecommend
		name := mysqlConn.KsToAuth(item.KsID)
		err := tools.CopyStructValue(&item, &respRecommend)
		if err != nil {
			fmt.Println(err)
		}
		respRecommend.Auth = name
		respRecommends = append(respRecommends, respRecommend)
	}
	data = map[string]interface{}{
		"result": respRecommends,
	}
	c.AsciiJSON(http.StatusOK, data)
}
