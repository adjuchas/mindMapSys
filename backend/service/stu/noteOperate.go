package stu

import (
	"backend/model/models"
	mysqlConn "backend/model/mysql"
	redisConn "backend/model/redis"
	"backend/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNotes(c *gin.Context) {
	var resData map[string]interface{}
	var data models.RequestGetDraftInfo
	tools.MapToStruct(c.GetStringMapString("data"), data)
	result := mysqlConn.GetNotesByStuId(data.Id)
	resData = map[string]interface{}{
		"result": result,
	}
	c.AsciiJSON(http.StatusOK, resData)
}

func StarNote(c *gin.Context) {
	var resData map[string]interface{}
	var data models.RequestStartNote
	tools.MapToStruct(c.GetStringMapString("data"), data)
	result := redisConn.SetStarNote(data.StuId, data.NoteId)
	resData = map[string]interface{}{
		"result": result,
	}
	c.AsciiJSON(http.StatusOK, resData)
}

func SetNoteState(c *gin.Context) {
	var resData map[string]interface{}
	var data models.RequestSetNoteState
	tools.MapToStruct(c.GetStringMapString("data"), data)
	result := mysqlConn.SetNotState(data.NoteId, data.State)
	resData = map[string]interface{}{
		"result": result,
	}
	c.AsciiJSON(http.StatusOK, resData)
}

func GetStartNotes(c *gin.Context) {
	var resData map[string]interface{}
	var data models.RequestSetDraft
	tools.MapToStruct(c.GetStringMapString("data"), data)
	noteArray := redisConn.GetStarNotes(data.StuId)
	result := mysqlConn.GetStarNote(noteArray)
	resData = map[string]interface{}{
		"result": result,
	}
	c.AsciiJSON(http.StatusOK, resData)
}

//func GetNoteBody(c *gin.Context) {
//	var data map[string]interface{}
//	var requestInfo models.RequestStartNote
//	c.ShouldBindJSON(&requestInfo)
//	if !redisConn.CheckUserByStuid(requestInfo.AccessToken, requestInfo.StuId) {
//		data = map[string]interface{}{
//			"result": false,
//		}
//	} else {
//		path := mysqlConn.GetNoteNodeTreePath(requestInfo.NoteId)
//
//		var nodeTree *os.File
//		var jsonByte []byte
//		var err error
//		nodeTree, err = os.Open(path)
//		if err != nil {
//			fmt.Println("打开文件时出错:", err)
//			return
//		}
//		defer nodeTree.Close()
//
//		jsonByte, err = ioutil.ReadAll(nodeTree)
//		if err != nil {
//			fmt.Println("读取文件时出错:", err)
//			return
//		}
//		var result map[string]interface{}
//		err = json.Unmarshal(jsonByte, &result)
//		if err != nil {
//			fmt.Println("解码JSON时出错:", err)
//			return
//		}
//		data = map[string]interface{}{
//			"result": result,
//		}
//	}
//	c.AsciiJSON(http.StatusOK, data)
//}
