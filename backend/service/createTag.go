package service

import (
	mysqlConn "backend/model/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTag(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	leaderId := mysqlConn.TeaidToKs(data["id"])
	mysqlConn.CreateTag(leaderId, data["tagName"])
	resData = map[string]interface{}{
		"result": true,
	}
	c.AsciiJSON(http.StatusOK, resData)
}
