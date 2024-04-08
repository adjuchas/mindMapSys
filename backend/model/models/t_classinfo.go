package models

type ClassInfo struct {
	ClassID    string `gorm:"column:class_ID"`
	Name       string `gorm:"column:NAME"`
	LeaderID   string `gorm:"column:leader_ID"`
	STATE      string `gorm:"column:STATE"`
	StuNum     string `gorm:"column:stu_NUM"`
	CreateTime string `gorm:"column:CREATETIME"`
	UpdateTime string `gorm:"column:UPDATETIME"`
}

func (receiver ClassInfo) TableName() string {
	return "classinfo"
}
