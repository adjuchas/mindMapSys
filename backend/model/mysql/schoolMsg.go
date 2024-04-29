package mysqlConn

import (
	"backend/model/models"
	"time"
)

func UpdateInstitute(info models.InstituteInfo) {
	var exist models.InstituteInfo
	result := DB.Where(&info).First(&exist)
	if result.Error != nil { // 不存在
		DB.Create(&info)
	} else {
		DB.Model(&exist).Updates(models.InstituteInfo{UpdateTime: time.Now()})
	}
}

func UpdateClass(info models.ClassInfo) {
	var exist models.ClassInfo
	result := DB.Where(&info).First(&exist)
	if result.Error != nil { // 不存在
		DB.Create(&info)
	} else {
		DB.Model(&exist).Updates(models.ClassInfo{UpdateTime: time.Now()})
	}
}
