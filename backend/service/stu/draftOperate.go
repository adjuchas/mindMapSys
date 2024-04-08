package stu

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
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
		fmt.Println("err1")
		fmt.Println(err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("err2")
		fmt.Println(err)
	}

	resData = map[string]interface{}{
		"result": true,
	}
	c.AsciiJSON(http.StatusOK, resData)
}
