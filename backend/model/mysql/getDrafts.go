package mysqlConn

import "backend/model/models"

func GetTeaDrafts(Id string) []models.DraftInfo {
	tea := models.TeacherInfo{
		Teacher_ID: Id,
	}
	DB.Select("ks_ID").First(&tea)
	var drafts []models.DraftInfo
	DB.Select("draft_ID", "TITLE", "STATE", "TAGS", "DESCRIPTION", "CREATETIME", "UPDATETIME").Where("ks_ID = ?", tea.KsID).Find(&drafts)
	return drafts
}

func GetDrafts(stuId string) []models.DraftInfo {
	stu := models.Stu{
		StuId: stuId,
	}
	DB.Select("ks_ID").First(&stu)
	var drafts []models.DraftInfo
	DB.Select("draft_ID", "TITLE", "STATE", "TAGS", "DESCRIPTION", "CREATETIME", "UPDATETIME", "nodeTreePath").Where("ks_ID = ?", stu.KsID).Find(&drafts)
	return drafts
}
