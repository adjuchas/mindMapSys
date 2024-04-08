package service

import (
	mysqlConn "backend/model/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetDotState(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)

	result := mysqlConn.SetDotState(data["dotId"], data["state"])
	resData = map[string]interface{}{
		"result": result,
	}
	c.AsciiJSON(http.StatusOK, resData)
}
