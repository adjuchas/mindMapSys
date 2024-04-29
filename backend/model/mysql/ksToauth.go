package mysqlConn

import (
	"backend/model/models"
)

func KsToAuth(ks int) string {
	if ks < 1000 {
		tea := models.TeacherInfo{}
		DB.Select("NAME").Where("ks_ID = ?", ks).First(&tea)
		return tea.Name
	} else {
		stu := models.Stu{}
		DB.Select("NAME").Where("ks_ID = ?", ks).First(&stu)
		return stu.Name
	}
}
