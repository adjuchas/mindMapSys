package service

import (
	mysqlConn "backend/model/mysql"
	"backend/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type responseClassify struct {
	DotID        int       `gorm:"column:dot_ID"`
	TITLE        string    `gorm:"column:TITLE"`
	KsID         int       `gorm:"column:ks_ID"`
	TAGS         string    `gorm:"column:TAGS"`
	Auth         string    `json:"auth"`
	DESCRIPTION  string    `gorm:"column:DESCRIPTION"`
	UpdateTime   time.Time `gorm:"column:UPDATETIME"`
	NodeTreePath string    `json:"nodeTreePath"`
}

func GetClassify(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	tagid, _ := strconv.Atoi(data["classifyId"])
	tagInfo := mysqlConn.GetTagNameById(tagid)
	classifys := mysqlConn.GetClassify(tagInfo.Name)
	respClassifys := []responseClassify{}
	for _, item := range classifys {
		var respclassify responseClassify
		err := tools.CopyStructValue(&item, &respclassify)
		name := mysqlConn.KsToAuth(item.KsID)
		respclassify.Auth = name
		if err != nil {
			fmt.Println(err)
		}
		respClassifys = append(respClassifys, respclassify)
	}
	resData = map[string]interface{}{
		"result": respClassifys,
	}
	c.AsciiJSON(http.StatusOK, resData)
}
