package service

import (
	mysqlConn "backend/model/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateDraft(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)

	if c.GetString("Identity") == "学生" {
		mysqlConn.CreateDraft(data["id"], data["title"], data["tags"], data["description"])
		resData = map[string]interface{}{
			"result": true,
		}
	} else {
		draftid, filePath := mysqlConn.CreateTeaDraft(data["id"], data["title"], data["tags"], data["description"])
		ks := mysqlConn.TeaidToKs(data["id"])
		mysqlConn.CreateDot(draftid, ks, data["tags"], data["title"], data["description"], filePath, "1")
		resData = map[string]interface{}{
			"result": true,
		}
	}
	c.AsciiJSON(http.StatusOK, resData)
}

func NewDraftJson(draftJson map[string]interface{}) {

}
