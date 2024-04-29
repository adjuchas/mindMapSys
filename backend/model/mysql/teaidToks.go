package mysqlConn

import "backend/model/models"

func TeaidToKs(teaid string) int {
	var tea models.TeacherInfo
	DB.Where("teacher_ID = ?", teaid).First(&tea)
	return tea.KsID
}
