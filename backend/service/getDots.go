package service

import (
	mysqlConn "backend/model/mysql"
	"backend/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type responseDot struct {
	DotID        int       `json:"dotID"`
	TITLE        string    `json:"title"`
	STATE        string    `json:"state"`
	TAGS         string    `json:"tags"`
	DESCRIPTION  string    `json:"description"`
	CreateTime   time.Time `json:"createTime"`
	UpdateTime   time.Time `json:"updateTime"`
	NodeTreePath string    `json:"nodeTreePath"`
}

func GetDots(c *gin.Context) {
	var resData map[string]interface{}
	var responseDots []responseDot
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	if c.GetString("Identity") == "学生" {
		dots := mysqlConn.GetDots(data["id"])
		for _, item := range dots {
			var responseDot responseDot
			err := tools.CopyStructValue(&item, &responseDot)
			if err != nil {
				fmt.Println(err)
			}
			responseDots = append(responseDots, responseDot)
		}
		resData = map[string]interface{}{
			"result": responseDots,
		}

	} else {
		dots := mysqlConn.GetTeaDots(data["Id"])
		for _, item := range dots {
			var responseDot responseDot
			err := tools.CopyStructValue(&item, &responseDot)
			if err != nil {
				fmt.Println(err)
			}
			responseDots = append(responseDots, responseDot)
		}
		resData = map[string]interface{}{
			"result": responseDots,
		}
	}
	c.AsciiJSON(http.StatusOK, resData)
}
