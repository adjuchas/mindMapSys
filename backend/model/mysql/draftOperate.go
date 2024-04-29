package mysqlConn

import (
	"backend/model/models"
	"backend/tools"
	"fmt"
	"strconv"
	"time"
)

func DraftCommit(draftId string) bool {
	result := DB.Model(&models.DraftInfo{}).Where("draft_ID = ?", draftId).Update("STATE", "1")
	if result.Error != nil {
		fmt.Println("is draft commit error")
		fmt.Println(result.Error)
		return false
	}
	return true
}

func DraftDel(draftId string) bool {
	result := DB.Delete(&models.DraftInfo{}, "draft_ID = ?", draftId)
	if result.Error != nil {
		fmt.Println("is draft del error")
		fmt.Println(result.Error)
		return false
	}
	return true
}

func DraftFuzzyFind(fuzzy bool, state string, key string, stuId string) []models.DraftInfo {
	stu := models.Stu{
		StuId: stuId,
	}
	DB.Select("ks_ID").First(&stu)
	var drafts []models.DraftInfo

	if fuzzy {
		query := DB.Select("draft_ID", "TITLE", "STATE", "TAGS", "DESCRIPTION", "CREATETIME", "UPDATETIME").Where("ks_id = ? and TITLE like ?", stu.KsID, "%"+key+"%").Find(&drafts)
		if state != "" {
			query = query.Where("STATE = ?", state)
		}
		query.Find(&drafts)
		return drafts
	} else {
		query := DB.Select("draft_ID", "TITLE", "STATE", "TAGS", "DESCRIPTION", "CREATETIME", "UPDATETIME").Where("ks_id = ? and TITLE = ?", stu.KsID, key).Find(&drafts)
		if state != "" {
			query = query.Where("STATE = ?", state)
		}
		query.Find(&drafts)
		return drafts
	}
}

func DraftTeaFuzzyFind(fuzzy bool, state string, key string, Id string) []models.DraftInfo {
	tea := models.TeacherInfo{
		Teacher_ID: Id,
	}
	DB.Select("ks_ID").First(&tea)
	var drafts []models.DraftInfo
	if fuzzy {
		DB.Where("ks_id = ? and STATE = ? and TITLE like ?", tea.KsID, state, "%"+key+"%").Find(&drafts)
		return drafts
	} else {
		DB.Where("ks_id = ? and STATE = ? and TITLE = ?", tea.KsID, state, key).Find(&drafts)
		return drafts
	}
}

func GetDraftNoteTreePath(draftid string) string {
	var draft models.DraftInfo
	DB.Select("nodeTreePath").Where("draft_ID = ?", draftid).First(&draft)
	return draft.NodeTreePath
}

func CreateDraft(stuId string, title string, tags string, description string) bool {
	stu := models.Stu{
		StuId: stuId,
	}
	DB.Select("ks_ID", "NAME").First(&stu)
	draft := &models.DraftInfo{
		KsID:        stu.KsID,
		TAGS:        tags,
		TITLE:       title,
		DESCRIPTION: description,
		STATE:       "0",
	}
	result := DB.Create(draft)

	filePath := tools.InitDraft(stuId, title, stu.Name, draft.DraftID)
	DB.Model(&models.DraftInfo{DraftID: draft.DraftID}).Update("nodeTreePath", filePath)
	if result.Error != nil {
		fmt.Println("create draft err")
		return false
	}
	return true
}

func CreateTeaDraft(Id string, title string, tags string, description string) (int, string) {
	tea := models.TeacherInfo{
		Teacher_ID: Id,
	}
	DB.Select("ks_ID").First(&tea)
	draft := models.DraftInfo{
		KsID:        tea.KsID,
		TAGS:        tags,
		TITLE:       title,
		DESCRIPTION: description,
		STATE:       "3",
	}
	result := DB.Create(&draft)

	if result.Error != nil {
		fmt.Println("create draft err")
		return 0, ""
	}
	filePath := tools.InitDraft(Id, title, tea.Name, draft.DraftID)
	DB.Model(&models.DraftInfo{DraftID: draft.DraftID}).Update("nodeTreePath", filePath)
	return draft.DraftID, filePath
}

func DelDraft(dotId string) bool {
	result := DB.Delete(&models.DraftInfo{}, "draft_ID = ?", dotId)
	if result.Error != nil {
		fmt.Println("is draft del error")
		fmt.Println(result.Error)
		return false
	}
	return true
}

func FindDraft(draftId string) models.DraftInfo {
	d, _ := strconv.Atoi(draftId)
	draft := models.DraftInfo{
		DraftID: d,
	}
	DB.First(&draft)
	return draft
}

func PassDraft(draftId int) {
	draft := models.DraftInfo{
		DraftID: draftId,
	}
	DB.Model(&draft).Update("STATE", "3")
}

func UnPassDraft(draftId int) {
	draft := models.DraftInfo{
		DraftID: draftId,
	}

	DB.Model(&draft).Update("STATE", "2")
}

func UpdateDraft(ID int) {
	DB.Model(&models.DraftInfo{}).Where("draft_ID = ?", ID).Update("UPDATETIME", time.Now())
}
