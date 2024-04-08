package models

type Stu struct {
	KsID        string `gorm:"column:ks_ID"`
	YbId        string `gorm:"column:yb_ID"`
	StuId       string `gorm:"column:stu_ID"`
	Name        string `gorm:"column:NAME"`
	InstituteId string `gorm:"column:institute_ID"`
	ClassId     string `gorm:"column:class_ID"`
	Email       string `gorm:"column:EMAIL"`
	CreateTime  string `gorm:"column:CREATETIME"`
	UpdateTime  string `gorm:"column:UPDATETIME"`
}

func (receiver Stu) TableName() string {
	return "stuinfo"
}
