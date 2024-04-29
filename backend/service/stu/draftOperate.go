package stu

import (
	mysqlConn "backend/model/mysql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func GetDraftBody(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	jsonData, _ := ioutil.ReadFile(data["draftPath"])
	resData = map[string]interface{}{
		"result": string(jsonData),
	}
	c.AsciiJSON(http.StatusOK, resData)
}

func SetDraftBody(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	filePath := data["draftPath"]
	jsonMsg, _ := c.Get("nodeJson")
	jsonData, _ := json.Marshal(jsonMsg)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
	}

	_, err = file.Write(jsonData)
	if err != nil {
	}
	draftid, _ := strconv.Atoi(data["draftId"])
	mysqlConn.UpdateDraft(draftid)
	resData = map[string]interface{}{
		"result": true,
	}
	c.AsciiJSON(http.StatusOK, resData)
}
