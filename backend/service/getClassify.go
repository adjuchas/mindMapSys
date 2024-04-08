package service

import (
	mysqlConn "backend/model/mysql"
	"backend/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type responseClassify struct {
	DotID       string `gorm:"column:dot_ID"`
	TITLE       string `gorm:"column:TITLE"`
	KsID        string `gorm:"column:ks_ID"`
	TAGS        string `gorm:"column:TAGS"`
	DESCRIPTION string `gorm:"column:DESCRIPTION"`
	UpdateTime  string `gorm:"column:UPDATETIME"`
}

func GetClassify(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	tagInfo := mysqlConn.GetTagNameById(data["classifyId"])
	classifys := mysqlConn.GetClassify(tagInfo.Name)
	respClassifys := []responseClassify{}
	for _, item := range classifys {
		var respclassify responseClassify
		err := tools.CopyStructValue(&item, &respclassify)
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
