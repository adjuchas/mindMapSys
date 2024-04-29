package service

import (
	mysqlConn "backend/model/mysql"
	"backend/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ResponseClassifyInfo struct {
	TagID      int       `json:"tagID"`
	Name       string    `json:"name"`
	DotNum     string    `json:"dotNum"`
	LeaderID   int       `json:"leaderID"`
	LeaderName string    `json:"leaderName"`
	CreateTime time.Time `json:"createTime"`
}

func GetClassifyInfo(c *gin.Context) {
	var resData map[string]interface{}
	var resInfo []ResponseClassifyInfo
	tags := mysqlConn.GetTags()
	for _, item := range tags {
		var resII ResponseClassifyInfo
		err := tools.CopyStructValue(&item, &resII)
		if err != nil {
			fmt.Println(err)
		}
		resII.LeaderName = mysqlConn.KsToAuth(resII.LeaderID)
		resInfo = append(resInfo, resII)
	}

	resData = map[string]interface{}{
		"result": resInfo,
	}
	c.AsciiJSON(http.StatusOK, resData)
}
