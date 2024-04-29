package stu

import (
	mysqlConn "backend/model/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func SetNote(c *gin.Context) {
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	noteFileName := fmt.Sprintf("%s%s.md", data["id"], data["dotId"])
	mdData := []byte(data["mdData"])
	filePath := filepath.Join("uploads", "note", noteFileName)
	mysqlConn.UpdateNote(filePath, data["id"], data["dotId"])
	_ = ioutil.WriteFile(filePath, mdData, 0644)
	c.AsciiJSON(http.StatusOK, map[string]string{
		"result": "success",
	})
}

func GetNote(c *gin.Context) {
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	noteFileName := fmt.Sprintf("%s%s.md", data["id"], data["dotId"])
	filePath := filepath.Join("uploads", "note", noteFileName)
	mdData, _ := ioutil.ReadFile(filePath)
	c.AsciiJSON(http.StatusOK, map[string]string{
		"result": string(mdData),
	})
}

func GetNoteList(c *gin.Context) {
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)
	ksId := mysqlConn.StuToKs(data["id"])
	notes := mysqlConn.GetNotes(ksId)
	for i := 0; i < len(notes); i++ {
		notes[i].TITLE = mysqlConn.GetDotTitle(notes[i].DotID)
		notes[i].NotePath = mysqlConn.GetNodeTreePath(notes[i].DotID)
	}
	c.AsciiJSON(http.StatusOK, map[string]interface{}{
		"result": notes,
	})
}
