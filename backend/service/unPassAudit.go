package service

import (
	mysqlConn "backend/model/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UnPassAudit(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	draftid, _ := strconv.Atoi(data["draftId"])
	teaId := data["tea"]
	mysqlConn.UnPassDraft(draftid)
	msg := data["msg"]
	if len(msg) != 0 {
		mysqlConn.SetMsg(msg, draftid, teaId)
	}
	resData = map[string]interface{}{
		"result": true,
	}
	c.AsciiJSON(http.StatusOK, resData)
}
