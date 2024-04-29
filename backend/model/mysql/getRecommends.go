package mysqlConn

import "backend/model/models"

func GetRecommends() []models.DotInfo {
	var recommends []models.DotInfo
	DB.Order("UPDATETIME desc").Where("STATE = ?", 0).Find(&recommends)
	return recommends
}
