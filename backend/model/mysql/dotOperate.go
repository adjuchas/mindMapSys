package mysqlConn

import (
	"backend/model/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
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

func GetDotNodeTreePath(dotid int) string {
	var dot models.DotInfo
	DB.Select("nodeTreePath").Where("dot_ID = ?", dotid).First(&dot)
	return dot.NodeTreePath
}

func GetTitle(dotid int) string {
	var dot models.DotInfo
	DB.Select("TITLE").Where("dot_ID = ?", dotid).First(&dot)
	return dot.NodeTreePath
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
		query := DB.Select("dot_ID", "TITLE", "STATE", "TAGS", "DESCRIPTION", "CREATETIME", "UPDATETIME").Where("ks_id = ? and TITLE = ?", tea.KsID, key).Find(&dots)
		if state != "" {
			query = query.Where("STATE = ?", state)
		}
		query.Find(&dots)
		return dots
	}
}

func CreateDot(dotId int, Id int, tags string, title string, description string, path string, state string) bool {
	dot := models.DotInfo{
		DotID:        dotId,
		KsID:         Id,
		TAGS:         tags,
		TITLE:        title,
		DESCRIPTION:  description,
		STATE:        state,
		NodeTreePath: path,
	}
	var aa models.DotInfo
	err := DB.Where("dot_ID = ?", dotId).First(&aa).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {

		result := DB.Create(&dot)

		if result.Error != nil {
			fmt.Println("create draft err")
			return false
		}
	} else {
		DB.Save(&dot)
	}
	return true
}

func DelDot(dotid string) bool {
	result := DB.Delete(&models.DotInfo{}, "dot_ID = ?", dotid)
	if result.Error != nil {
		fmt.Println("is draft del error")
		fmt.Println(result.Error)
		return false
	}
	return true
}
