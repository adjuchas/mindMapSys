package service

import (
	mysqlConn "backend/model/mysql"
	"backend/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type responseDraft struct {
	DraftID      int       `json:"draftID"`
	TITLE        string    `json:"title"`
	STATE        string    `json:"state"`
	TAGS         string    `json:"tags"`
	DESCRIPTION  string    `json:"description"`
	CreateTime   time.Time `json:"createTime"`
	NodeTreePath string    `json:"nodeTreePath"`
	UpdateTime   time.Time `json:"updateTime"`
}

func GetDrafts(c *gin.Context) {
	var responseDrafts []responseDraft
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	if c.GetString("Identity") == "学生" {
		drafts := mysqlConn.GetDrafts(data["Id"])
		for _, item := range drafts {
			var respDraft responseDraft
			err := tools.CopyStructValue(&item, &respDraft)
			if err != nil {
				fmt.Println(err)
			}
			responseDrafts = append(responseDrafts, respDraft)
		}
		resData = map[string]interface{}{
			"result": responseDrafts,
		}
	} else {
		drafts := mysqlConn.GetTeaDrafts(data["Id"])
		for _, item := range drafts {
			var respDraft responseDraft
			err := tools.CopyStructValue(&item, &respDraft)
			if err != nil {
				fmt.Println(err)
			}
			responseDrafts = append(responseDrafts, respDraft)
		}
		resData = map[string]interface{}{
			"result": responseDrafts,
		}
	}
	c.AsciiJSON(http.StatusOK, resData)
}
