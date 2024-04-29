package service

import (
	mysqlConn "backend/model/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DotDel(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	mysqlConn.DelDot(data["dotId"])
	mysqlConn.DelDraft(data["dotId"])
	resData = map[string]interface{}{
		"result": true,
	}
	c.AsciiJSON(http.StatusOK, resData)
}
