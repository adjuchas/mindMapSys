package service

import (
	"backend/model/models"
	mysqlConn "backend/model/mysql"
	redisConn "backend/model/redis"
	"backend/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
)

type responseRecommend struct {
	DotID       string `json:"dot_id"`
	TITLE       string `json:"title"`
	KsID        string `json:"ks_id"`
	TAGS        string `json:"tags"`
	DESCRIPTION string `json:"description"`
	UpdateTime  string `json:"UpdateTime"`
}

func GetRecommends(c *gin.Context) {
	var data map[string]interface{}
	var requestInfo models.RequestInfo
	c.ShouldBindJSON(&requestInfo)
	_, err := redisConn.GetInfoMsg(requestInfo.AccessToken)
	if err == redis.Nil {
		c.AsciiJSON(http.StatusOK, data)
		return
	}
	recommends := mysqlConn.GetRecommends()
	var respRecommends []responseRecommend
	for _, item := range recommends {
		var respRecommend responseRecommend
		err := tools.CopyStructValue(&item, &respRecommend)
		if err != nil {
			fmt.Println(err)
		}
		respRecommends = append(respRecommends, respRecommend)
	}
	data = map[string]interface{}{
		"result": respRecommends,
	}
	c.AsciiJSON(http.StatusOK, data)
}
