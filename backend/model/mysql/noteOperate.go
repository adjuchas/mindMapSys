package mysqlConn

import (
	"backend/model/models"
	"strconv"
	"time"
)

func UpdateNote(path string, stuId string, dotId string) {
	dot, _ := strconv.Atoi(dotId)
	var exist models.NoteInfo
	note := models.NoteInfo{
		KsID:     StuToKs(stuId),
		NotePath: path,
		DotID:    dot,
	}
	result := DB.Where(&note).First(&exist)
	if result.Error != nil { // 不存在
		DB.Create(&note)
	} else {
		DB.Model(&exist).Updates(models.ClassInfo{UpdateTime: time.Now()})
	}
}

func GetNotes(ksId int) []models.NoteInfo {
	var notes []models.NoteInfo
	DB.Where("ks_ID = ?", ksId).Find(&notes)
	return notes
}

func GetDotTitle(dotID int) string {
	var dot models.DotInfo
	DB.Where("dot_ID = ?", dotID).First(&dot)
	return dot.TITLE
}

func GetNodeTreePath(dotID int) string {
	var dot models.DotInfo
	DB.Where("dot_ID = ?", dotID).First(&dot)
	return dot.NodeTreePath
}
