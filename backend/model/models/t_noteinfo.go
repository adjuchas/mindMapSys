package models

import "time"

type NoteInfo struct {
	NoteID     int       `gorm:"column:note_ID;primaryKey;auto_increment" json:"noteID"`
	TITLE      string    `gorm:"column:TITLE;null" json:"title"`
	KsID       int       `gorm:"column:ks_ID" json:"ksID"`
	STATE      int       `gorm:"column:STATE" json:"state"`
	DotID      int       `gorm:"column:dot_ID" json:"dotID"`
	NotePath   string    `gorm:"column:notePath" json:"notePath"`
	CreateTime time.Time `gorm:"column:CREATETIME;type:datetime(0);autoCreateTime" json:"createTime"`
	UpdateTime time.Time `gorm:"column:UPDATETIME;type:datetime(0);autoUpdateTime" json:"updateTime"`
}

func (receiver NoteInfo) TableName() string {
	return "noteinfo"
}
