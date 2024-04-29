package service

import (
	mysqlConn "backend/model/mysql"
	"backend/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

type respAudit struct {
	DraftID      int    `json:"draftID"`
	TITLE        string `json:"title"`
	Auth         string `json:"auth"`
	TAGS         string `json:"tags"`
	DESCRIPTION  string `json:"description"`
	NodeTreePath string `json:"nodeTreePath"`
}

func Audit(c *gin.Context) {
	var resData map[string]interface{}
	var auditData []respAudit
	drafts := mysqlConn.Audit()
	for _, item := range drafts {
		var info respAudit
		tools.CopyStructValue(&item, &info)
		name := mysqlConn.KsToAuth(item.KsID)
		info.Auth = name
		auditData = append(auditData, info)
	}
	resData = map[string]interface{}{
		"result": auditData,
	}
	c.AsciiJSON(http.StatusOK, resData)
}
