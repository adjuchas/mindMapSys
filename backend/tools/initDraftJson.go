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
		"data": map[string]interface{}{
			"id":       "root",
			"topic":    "输入根节点",
			"expanded": true,
			"children": []map[string]interface{}{
				{
					"id":        "left1",
					"topic":     "输入节点内容",
					"direction": "left",
					"expanded":  true,
				},
				{
					"id":        "right1",
					"topic":     "输入节点内容",
					"direction": "right",
					"expanded":  true,
				},
				{
					"id":        "right2",
					"topic":     "输入节点内容",
					"direction": "right",
					"expanded":  true,
				},
				{
					"id":        "left2",
					"topic":     "输入节点内容",
					"direction": "left",
					"expanded":  true,
				}},
		},
	}

	jsonData, _ := json.MarshalIndent(draftInit, "", "  ")

	filePath := filepath.Join("uploads", "draft", fileName)
	_ = ioutil.WriteFile(filePath, jsonData, 0644)
	return filePath
}
