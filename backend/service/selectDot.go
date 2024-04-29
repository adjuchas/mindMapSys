package service

import (
	mysqlConn "backend/model/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DotSelect(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	fuzzy, _ := strconv.ParseBool(data["fuzzy"])
	if c.GetString("Identity") == "学生" {
		drafts := mysqlConn.DotFuzzyFind(fuzzy, data["state"], data["title"], data["id"])
		resData = map[string]interface{}{
			"result": drafts,
		}
	} else {
		dots := mysqlConn.DotTeaFuzzyFind(fuzzy, data["state"], data["title"], data["id"])
		resData = map[string]interface{}{
			"result": dots,
		}
	}
	c.AsciiJSON(http.StatusOK, resData)
}
