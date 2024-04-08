package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func InitDraft(stuId string, title string, name string, draftId int) string {
	fileName := fmt.Sprintf("%sDraft%d.json", stuId, draftId)

	draftInit := map[string]interface{}{
		"meta": map[string]interface{}{
			"name":   title,
			"author": name,
		},
		"format": "node_tree",
		"data":   "",
	}

	jsonData, _ := json.MarshalIndent(draftInit, "", "  ")

	filePath := filepath.Join("uploads", "draft", fileName)
	_ = ioutil.WriteFile(filePath, jsonData, 0644)
	return filePath
}
