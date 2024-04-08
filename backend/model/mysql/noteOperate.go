package mysqlConn

import (
	"backend/model/models"
	"fmt"
)

func GetNotesByStuId(stuId string) []models.NoteInfo {
	stu := models.Stu{
		StuId: stuId,
	}
	DB.Select("ks_ID").First(&stu)
	var notes []models.NoteInfo
	DB.Select("note_ID", "TITLE", "dot_ID", "STATE", "CREATETIME", "UPDATETIME").Where("ks_ID = ?", stu.KsID).Find(&notes)
	return notes
}

func SetNotState(noteId string, state string) bool {
	err := DB.Model(&models.NoteInfo{}).Where("note_ID = ?", noteId).Update("STATE", state)
	if err != nil {
		fmt.Println("setState is err")
		return false
	} else {
		return true
	}
}

func GetStarNote(noteArray []string) []models.NoteInfo {
	var notes []models.NoteInfo
	for _, val := range noteArray {
		note := models.NoteInfo{
			DotID: val,
		}
		DB.First(&note)
		notes = append(notes, note)
	}
	return notes
}

func GetNoteNodeTreePath(noteId string) string {
	var note models.NoteInfo
	DB.Select("notePath").Where("note_ID = ?", noteId).First(&note)
	return note.NotePath
}
