package mysqlConn

import (
	"backend/model/models"
)

func GetDots(stuId string) []models.DotInfo {
	var stu models.Stu
	DB.Select("ks_ID").Where("stu_ID = ?", stuId).First(&stu)
	var dots []models.DotInfo
	DB.Select("dot_ID", "TITLE", "STATE", "TAGS", "DESCRIPTION", "nodeTreePath", "CREATETIME", "UPDATETIME").Where("ks_ID = ?", stu.KsID).Find(&dots)
	return dots
}

func GetTeaDots(id string) []models.DotInfo {
	var tea models.TeacherInfo
	DB.Debug().Select("ks_ID").Where("teacher_ID", id).First(&tea)
	var dots []models.DotInfo
	DB.Select("dot_ID", "TITLE", "STATE", "TAGS", "DESCRIPTION", "CREATETIME", "UPDATETIME", "nodeTreePath").Where("ks_ID = ?", tea.KsID).Find(&dots)
	return dots
}

func IdToDot(dotId int) models.DotInfo {
	dot := models.DotInfo{}
	DB.First(&dot, dotId)
	return dot
}
