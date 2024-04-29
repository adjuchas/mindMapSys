package models

import "time"

type InstituteInfo struct {
	InstituteID int       `gorm:"column:institute_ID"`
	Name        string    `gorm:"column:NAME"`
	Description string    `gorm:"column:DESCRIPTION"`
	LeaderID    int       `gorm:"column:leader_ID"`
	STATE       string    `gorm:"column:STATE; default:1"`
	CreateTime  time.Time `gorm:"column:CREATETIME;type:datetime(0);autoCreateTime"`
	UpdateTime  time.Time `gorm:"column:UPDATETIME;type:datetime(0);autoUpdateTime"`
}

func (receiver InstituteInfo) TableName() string {
	return "instituteinfo"
}
