package mysqlConn

import (
	"backend/model/models"
)

func SetMsg(msg string, draftId int, teaId string) {

	ksStu := DraftToKs(draftId)
	ksTea := TeaidToKs(teaId)
	msgObj := models.Message{
		Msg:     msg,
		KsStu:   ksStu,
		KsTea:   ksTea,
		DraftID: draftId,
	}
	DB.Select("msg", "ks_tea", "ks_stu", "draftID").Create(&msgObj)
}
