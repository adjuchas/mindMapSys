package mysqlConn

import "backend/model/models"

func CreateTag(id int, name string) {
	tag := models.Tag{
		LeaderID: id,
		Name:     name,
		DotNum:   "0",
	}
	DB.Create(&tag)
}
