package service

import (
	mysqlConn "backend/model/mysql"
	redisConn "backend/model/redis"
	"backend/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func StarDot(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	Id := data["Id"]
	dotId := data["dotId"]
	result := redisConn.SetStarDot(Id, dotId)
	resData = map[string]interface{}{
		"result": result,
	}
	c.AsciiJSON(http.StatusOK, resData)
}

func CheckStar(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	Id := data["id"]
	dotId := data["dotId"]
	result := redisConn.CheckDotStat(Id, dotId)
	resData = map[string]interface{}{
		"result": result,
	}
	c.AsciiJSON(http.StatusOK, resData)
}

func QuitStar(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	Id := data["id"]
	dotId := data["dotId"]
	result := redisConn.QuitDotStat(Id, dotId)
	resData = map[string]interface{}{
		"result": result,
	}
	c.AsciiJSON(http.StatusOK, resData)
}

func GetStarDots(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	Id := data["Id"]
	var respResults []ResponseRecommend
	result := redisConn.GetStarDots(Id)
	fmt.Println(result)
	for _, item := range result {
		var respResult ResponseRecommend
		dotId, _ := strconv.Atoi(item)
		dot := mysqlConn.IdToDot(dotId)
		name := mysqlConn.KsToAuth(dot.KsID)
		tools.CopyStructValue(&dot, &respResult)
		respResult.Auth = name
		respResults = append(respResults, respResult)
	}

	resData = map[string]interface{}{
		"result": respResults,
	}
	c.AsciiJSON(http.StatusOK, resData)
}
