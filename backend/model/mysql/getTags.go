package mysqlConn

import "backend/model/models"

func GetTags() []models.Tag {
	var tags []models.Tag
	DB.Select("tag_ID", "NAME", "dot_NUM").Find(&tags)
	return tags
}
