package models

import "time"

type Stu struct {
	KsID        int       `gorm:"column:ks_ID;auto_increment"`
	YbId        string    `gorm:"column:yb_ID"`
	StuId       string    `gorm:"column:stu_ID"`
	Name        string    `gorm:"column:NAME"`
	InstituteId string    `gorm:"column:institute_ID"`
	ClassId     string    `gorm:"column:class_ID"`
	Email       string    `gorm:"column:EMAIL"`
	CreateTime  time.Time `gorm:"column:CREATETIME;type:datetime(0);autoUpdateTime"`
	UpdateTime  time.Time `gorm:"column:UPDATETIME;type:datetime(0);autoUpdateTime"`
}

func (receiver Stu) TableName() string {
	return "stuinfo"
}
