package mysqlConn

import (
	"backend/model/models"
)

func FindStu(id string) bool {
	result := DB.Where("stu_ID = ?", id).First(&models.Stu{})
	if result.Error != nil {
		return false
	}
	return true
}

func SetStu(stu models.Stu) {
	DB.Select("yb_ID", "stu_ID", "NAME", "institute_ID", "class_ID").Create(&stu)
}
