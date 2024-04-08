package stu

import (
	"backend/model/models"
	redisConn "backend/model/redis"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHistory(c *gin.Context) {

	var data map[string]interface{}
	var resp models.RequestInfo
	c.ShouldBindJSON(&resp)
	info, _ := redisConn.GetInfoMsg(resp.AccessToken)
	val, err := redisConn.GetHistory(info.Yb_studentid)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	data = map[string]interface{}{
		"history": val,
	}
	c.AsciiJSON(http.StatusOK, data)
}
