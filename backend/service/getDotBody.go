package service

import (
	redisConn "backend/model/redis"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetDotBody(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	jsonData, _ := ioutil.ReadFile(data["nodePath"])
	dotid, _ := strconv.Atoi(data["dotId"])
	redisConn.SetHistory(data["id"], dotid)
	resData = map[string]interface{}{
		"result": string(jsonData),
	}
	c.AsciiJSON(http.StatusOK, resData)
}
