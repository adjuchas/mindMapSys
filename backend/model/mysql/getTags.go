package mysqlConn

import "backend/model/models"

func GetTags() []models.Tag {
	var tags []models.Tag
	DB.Find(&tags)
	return tags
}
