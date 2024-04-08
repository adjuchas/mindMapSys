package service

import (
	"backend/model/models"
	redisConn "backend/model/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StarDot(c *gin.Context) {
	var resData map[string]interface{}
	var data models.RequestGetDot

	result := redisConn.SetStarDot(data.Id, data.DotId)
	resData = map[string]interface{}{
		"result": result,
	}

	c.AsciiJSON(http.StatusOK, resData)
}
