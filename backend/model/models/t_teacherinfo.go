package models

import "time"

type TeacherInfo struct {
	KsID       int       `gorm:"column:ks_ID;auto_increment"`
	Teacher_ID string    `gorm:"column:teacher_ID"`
	Yb_ID      string    `gorm:"column:yb_ID"`
	Name       string    `gorm:"column:NAME"`
	Email      string    `gorm:"column:EMAIL"`
	Phone      string    `gorm:"PHONE"`
	CreateTime time.Time `gorm:"column:CREATETIME;type:datetime(0);autoUpdateTime"`
	UpdateTime time.Time `gorm:"column:UPDATETIME;type:datetime(0);autoUpdateTime"`
}

func (receiver TeacherInfo) TableName() string {
	return "teacherinfo"
}
