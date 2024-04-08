package mysqlConn

import (
	"backend/model/models"
	"fmt"
)

func SetDotState(dotId string, state string) bool {
	result := DB.Model(&models.DotInfo{}).Where("dot_ID = ?", dotId).Update("STATE", state)
	if result.Error != nil {
		fmt.Println("setState is err")
		return false
	} else {
		return true
	}
}

func GetDotNoteTreePath(dotid string) string {
	var dot models.DotInfo
	DB.Select("nodeTreePath").Where("dot_ID = ?", dotid).First(&dot)
	return dot.NodeTreePath
}

func GetStarDot(idArray []string) []models.DotInfo {
	var dots []models.DotInfo
	for _, val := range idArray {
		dot := models.DotInfo{
			DotID: val,
		}
		DB.First(&dot)
		dots = append(dots, dot)
	}
	return dots
}

func DotFuzzyFind(fuzzy bool, state string, key string, id string) []models.DotInfo {
	stu := models.Stu{
		StuId: id,
	}
	DB.Select("ks_ID").First(&stu)
	var dots []models.DotInfo

	if fuzzy {
		query := DB.Select("dot_ID", "TITLE", "STATE", "TAGS", "DESCRIPTION", "CREATETIME", "UPDATETIME", "nodeTreePath").Where("ks_id = ? and TITLE like ?", stu.KsID, "%"+key+"%").Find(&dots)
		if state != "" {
			query = query.Where("STATE = ?", state)
		}
		query.Find(&dots)
		return dots
	} else {
		query := DB.Select("dot_ID", "TITLE", "STATE", "TAGS", "DESCRIPTION", "CREATETIME", "UPDATETIME").Where("ks_id = ? and TITLE = ?", stu.KsID, key).Find(&dots)
		if state != "" {
			query = query.Where("STATE = ?", state)
		}
		query.Find(&dots)
		return dots
	}
}

func DotTeaFuzzyFind(fuzzy bool, state string, key string, id string) []models.DotInfo {
	tea := models.TeacherInfo{
		Teacher_ID: id,
	}
	DB.Select("ks_ID").First(&tea)
	var dots []models.DotInfo

	if fuzzy {
		query := DB.Select("dot_ID", "TITLE", "STATE", "TAGS", "DESCRIPTION", "CREATETIME", "UPDATETIME", "nodeTreePath").Where("ks_id = ? and TITLE like ?", tea.KsID, "%"+key+"%").Find(&dots)
		if state != "" {
			query = query.Where("STATE = ?", state)
		}
		query.Find(&dots)
		return dots
	} else {
		query := DB.Select("draft_ID", "TITLE", "STATE", "TAGS", "DESCRIPTION", "CREATETIME", "UPDATETIME").Where("ks_id = ? and TITLE = ?", tea.KsID, key).Find(&dots)
		if state != "" {
			query = query.Where("STATE = ?", state)
		}
		query.Find(&dots)
		return dots
	}
}
