package models

import "time"

type Tag struct {
	TagID      int       `gorm:"column:tag_ID;auto_increment"`
	Name       string    `gorm:"column:NAME"`
	DotNum     string    `gorm:"column:dot_NUM"`
	LeaderID   int       `gorm:"column:leader_ID"`
	CreateTime time.Time `gorm:"column:CREATETIME;type:datetime(0);autoUpdateTime" `
	UpdateTime time.Time `gorm:"column:UPDATETIME;type:datetime(0);autoUpdateTime"`
}

func (receiver Tag) TableName() string {
	return "taginfo"
}
