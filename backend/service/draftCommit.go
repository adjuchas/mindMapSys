package service

import (
	mysqlConn "backend/model/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DraftCommit(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)

	ok := mysqlConn.DraftCommit(data["draftId"])
	if !ok {
		resData = map[string]interface{}{
			"result": false,
		}
		c.AsciiJSON(http.StatusOK, resData)
		return
	}
	resData = map[string]interface{}{
		"result": true,
	}
	c.AsciiJSON(http.StatusOK, resData)
}
