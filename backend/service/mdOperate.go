package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func SetMd(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)

	mdPath := data["mdPath"]
	mdData := data["mdData"]

	filePath := filepath.Join("uploads", "md", mdPath)
	_ = ioutil.WriteFile(filePath, []byte(mdData), 0644)

	resData = map[string]interface{}{
		"result": true,
	}
	c.AsciiJSON(http.StatusOK, resData)
}

func GetMd(c *gin.Context) {
	var resData map[string]interface{}
	mid, _ := c.Get("data")
	data, _ := mid.(map[string]string)

	mdPath := data["mdPath"]
	filePath := filepath.Join("uploads", "md", mdPath)
	mdData, _ := ioutil.ReadFile(filePath)

	resData = map[string]interface{}{
		"result": string(mdData),
	}
	c.AsciiJSON(http.StatusOK, resData)
}

func UploadImg(c *gin.Context) {
	id := uuid.New().String()
	file, _ := c.FormFile("file")
	originalFilename := file.Filename
	fileExt := filepath.Ext(originalFilename)
	filePath := filepath.Join("uploads", "img", id+fileExt)
	err := c.SaveUploadedFile(file, filePath)
	if err != nil {
		fmt.Println(err)
	}
	c.AsciiJSON(http.StatusOK, map[string]string{
		"result": filePath,
	})
}
