package mysqlConn

import "backend/model/models"

func GetDots(stuId string) []models.DotInfo {
	stu := models.Stu{
		StuId: stuId,
	}
	DB.Select("ks_ID").First(&stu)
	var dots []models.DotInfo
	DB.Select("dot_ID", "TITLE", "STATE", "TAGS", "DESCRIPTION", "CREATETIME", "UPDATETIME").Where("ks_ID = ?", stu.KsID).Find(&dots)
	return dots
}

func GetTeaDots(id string) []models.DotInfo {
	tea := models.TeacherInfo{
		Teacher_ID: id,
	}
	DB.Select("ks_ID").First(&tea)
	var dots []models.DotInfo
	DB.Select("dot_ID", "TITLE", "STATE", "TAGS", "DESCRIPTION", "CREATETIME", "UPDATETIME").Where("ks_ID = ?", tea.KsID).Find(&dots)
	return dots
}
