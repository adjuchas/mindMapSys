package mysqlConn

import (
	"backend/model/models"
	"fmt"
)

func GetMsg(ks int) []models.Message {
	var msgs []models.Message
	DB.Where("ks_stu = ?", ks).Find(&msgs)
	fmt.Println(msgs)
	return msgs
}
