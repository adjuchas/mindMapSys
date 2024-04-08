package models

type NoteInfo struct {
	NoteID     string `gorm:"column:note_ID"`
	TITLE      string `gorm:"column:TITLE"`
	KsID       string `gorm:"column:ks_ID"`
	STATE      string `gorm:"column:STATE"`
	DotID      string `gorm:"column:dot_ID"`
	NotePath   string `gorm:"column:notePath"`
	CreateTime string `gorm:"column:CREATETIME"`
	UpdateTime string `gorm:"column:UPDATETIME"`
}

func (receiver NoteInfo) TableName() string {
	return "noteinfo"
}
