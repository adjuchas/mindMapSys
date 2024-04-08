package service

import (
	mysqlConn "backend/model/mysql"
	"backend/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type responseTag struct {
	TagID  string
	Name   string
	DotNum string
}

func GetHasTags(c *gin.Context) {
	var data map[string]interface{}
	tags := mysqlConn.GetTags()
	var resptags []responseTag
	for _, item := range tags {
		var resptag responseTag
		err := tools.CopyStructValue(&item, &resptag)
		if err != nil {
			fmt.Println(err)
		}
		resptags = append(resptags, resptag)
	}
	data = map[string]interface{}{
		"result": resptags,
	}
	c.AsciiJSON(http.StatusOK, data)
}
