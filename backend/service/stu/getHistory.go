package stu

import (
	mysqlConn "backend/model/mysql"
	redisConn "backend/model/redis"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strconv"
)

type ResponseHistory struct {
	DotId        int    `json:"dotId"`
	Title        string `json:"title"`
	NodeTreePath string `json:"nodeTreePath"`
}

func GetHistory(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	history, err := redisConn.GetHistory(data["Id"])
	if err == redis.Nil {
		resData = map[string]interface{}{
			"result": history,
		}
		c.AsciiJSON(http.StatusOK, resData)
		return
	}
	var resResults []ResponseHistory
	for _, his := range history {
		var res ResponseHistory
		dotid, _ := strconv.Atoi(his)
		res.DotId = dotid
		res.NodeTreePath = mysqlConn.GetDotNodeTreePath(dotid)
		res.Title = mysqlConn.GetTitle(dotid)
	}
	resData = map[string]interface{}{
		"result": resResults,
	}
	c.AsciiJSON(http.StatusOK, resData)
}
