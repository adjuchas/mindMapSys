package mysqlConn

import (
	"backend/model/models"
)

func GetClassify(tagName string) []models.DotInfo {
	var dots []models.DotInfo
	DB.Order("UPDATETIME desc").Where("STATE = ? and TAGS like ?", "0", "%"+tagName+"%").Where("STATE = ?", 0).Find(&dots)
	return dots
}

func GetTagNameById(tagId int) models.Tag {
	var tag models.Tag
	DB.Select("NAME").Where("tag_ID = ?", tagId).First(&tag)
	return tag
}
