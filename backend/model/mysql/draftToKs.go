package mysqlConn

import "backend/model/models"

func DraftToKs(id int) int {
	d := models.DraftInfo{
		DraftID: id,
	}
	DB.First(&d)
	return d.KsID
}
