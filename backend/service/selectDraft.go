package service

import (
	mysqlConn "backend/model/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DraftSelect(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	fuzzy, _ := strconv.ParseBool(data["fuzzy"])
	if c.GetString("Identity") == "学生" {
		drafts := mysqlConn.DraftFuzzyFind(fuzzy, data["state"], data["title"], data["id"])
		resData = map[string]interface{}{
			"result": drafts,
		}
	} else {
		drafts := mysqlConn.DraftTeaFuzzyFind(fuzzy, data["state"], data["title"], data["id"])
		resData = map[string]interface{}{
			"result": drafts,
		}
	}
	c.AsciiJSON(http.StatusOK, resData)
}
