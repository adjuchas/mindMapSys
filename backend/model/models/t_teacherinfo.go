package models

type TeacherInfo struct {
	KsID       string `gorm:"column:ks_ID"`
	Teacher_ID string `gorm:"column:teacher_ID"`
	Yb_ID      string `gorm:"column:yb_ID"`
	Name       string `gorm:"column:NAME"`
	Email      string `gorm:"column:EMAIL"`
	Phone      string `gorm:"PHONE"`
	CreateTime string `gorm:"column:CREATETIME"`
	UpdateTime string `gorm:"column:UPDATETIME"`
}

func (receiver TeacherInfo) TableName() string {
	return "teacherinfo"
}
