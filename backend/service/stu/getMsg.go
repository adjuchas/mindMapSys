package stu

import (
	mysqlConn "backend/model/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type responseMsg struct {
	MessageID  int       `json:"messageID"`
	CreateTime time.Time `json:"createTime"`
	DraftID    int       `json:"draftID"`
	From       string    `json:"from"`
	Msg        string    `json:"msg"`
}

func GetMsg(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	ks := mysqlConn.StuToKs(data["Id"])
	var results []responseMsg
	msgs := mysqlConn.GetMsg(ks)
	for _, item := range msgs {
		var res responseMsg
		res.Msg = item.Msg
		res.MessageID = item.MessageID
		res.DraftID = item.DraftID
		res.CreateTime = item.CreateTime
		name := mysqlConn.KsToAuth(item.KsTea)
		res.From = name
		results = append(results, res)
	}

	resData = map[string]interface{}{
		"result": results,
	}
	c.AsciiJSON(http.StatusOK, resData)
}
