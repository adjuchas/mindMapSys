package stu

import (
	"backend/model/models"
	mysqlConn "backend/model/mysql"
	redisConn "backend/model/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStarDots(c *gin.Context) {
	var resData map[string]interface{}
	var data models.RequestSetDraft
	dotArray := redisConn.GetStarDots(data.StuId)
	result := mysqlConn.GetStarDot(dotArray)
	// 这里返回了所有的数据，而没有经过responseInfo的coryStruct的处理
	resData = map[string]interface{}{
		"result": result,
	}
	c.AsciiJSON(http.StatusOK, resData)
}

// 这个都没测试过，不知道对错
// 好像是错了
//func GetDotBody(c *gin.Context) {
//	var data map[string]interface{}
//	var requestInfo models.RequestGetDot
//	c.ShouldBindJSON(requestInfo)
//	if !redisConn.CheckUserByStuid(requestInfo.AccessToken, requestInfo.Id) {
//		data = map[string]interface{}{
//			"result": false,
//		}
//	} else {
//		path := mysqlConn.GetDotNoteTreePath(requestInfo.DotId)
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
