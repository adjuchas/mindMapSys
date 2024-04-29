package models

import "time"

type ClassInfo struct {
	ClassID    int       `gorm:"column:class_ID"`
	Name       string    `gorm:"column:NAME"`
	LeaderID   int       `gorm:"column:leader_ID"`
	STATE      string    `gorm:"column:STATE;default:1"`
	CreateTime time.Time `gorm:"column:CREATETIME;type:datetime(0);autoCreateTime"`
	UpdateTime time.Time `gorm:"column:UPDATETIME;type:datetime(0);autoUpdateTime"`
}

func (receiver ClassInfo) TableName() string {
	return "classinfo"
}
