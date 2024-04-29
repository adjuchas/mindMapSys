package mysqlConn

import "backend/model/models"

func StuToKs(stuid string) int {
	var stu models.Stu
	DB.Where("stu_ID = ?", stuid).First(&stu)
	return stu.KsID
}
