package mysqlConn

import "backend/model/models"

func Audit() []models.DraftInfo {
	var drafts []models.DraftInfo
	DB.Where("STATE = ?", "1").Find(&drafts)

	return drafts
}
