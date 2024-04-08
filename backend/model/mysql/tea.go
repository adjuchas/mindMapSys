package mysqlConn

import "backend/model/models"

func CheckTeaInfo(teamsg models.TeacherInfo) (models.TeacherInfo, bool) {
	isIn := true
	var resultTea models.TeacherInfo
	DB.Table("teacherinfo").Where("teacher_ID = ?", teamsg.Teacher_ID).Find(&resultTea)
	if resultTea == (models.TeacherInfo{}) {
		DB.Table("teacherinfo").Select("yb_ID", "teacher_ID", "NAME").Create(&teamsg)
		isIn = false
	}
	DB.Table("teacherinfo").Where("teacher_ID = ?", teamsg.Teacher_ID).Find(&teamsg)
	return teamsg, isIn
}

func FindTea(id string) bool {
	result := DB.Where("teacher_ID = ?", id).First(&models.TeacherInfo{})
	if result.Error != nil {
		return false
	}
	return true
}

func SetTea(tea models.TeacherInfo) {
	DB.Select("yb_ID", "teacher_ID", "NAME").Create(&tea)
}
