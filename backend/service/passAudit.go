package service

import (
	mysqlConn "backend/model/mysql"
	"backend/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PassAudit(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	nodePath := mysqlConn.GetDraftNoteTreePath(data["draftId"])
	newDotPath, err := tools.DraftToDot(nodePath)
	draftInfo := mysqlConn.FindDraft(data["draftId"])
	mysqlConn.PassDraft(draftInfo.DraftID)
	mysqlConn.CreateDot(draftInfo.DraftID, draftInfo.KsID, draftInfo.TAGS, draftInfo.TITLE, draftInfo.DESCRIPTION, newDotPath, "0")
	if err != nil {
		fmt.Println(err)
	}
	resData = map[string]interface{}{
		"result": true,
	}
	c.AsciiJSON(http.StatusOK, resData)
}
