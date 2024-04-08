package mysqlConn

import "backend/model/models"

func GetRecommends() []models.DotInfo {
	var recommends []models.DotInfo
	DB.Order("UPDATETIME desc").Select("dot_ID", "TITLE", "ks_ID", "TAGS", "DESCRIPTION", "UPDATETIME").Where("STATE = ?", 0).Find(&recommends)
	return recommends
}
