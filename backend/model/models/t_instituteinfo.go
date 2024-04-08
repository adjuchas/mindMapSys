package models

type InstituteInfo struct {
	InstituteID string `gorm:"column:institute_ID"`
	Name        string `gorm:"column:NAME"`
	Description string `gorm:"column:DESCRIPTION"`
	LeaderID    string `gorm:"column:leader_ID"`
	STATE       string `gorm:"column:STATE"`
	CreateTime  string `gorm:"column:CREATETIME"`
	UpdateTime  string `gorm:"column:UPDATETIME"`
}

func (receiver InstituteInfo) TableName() string {
	return "instituteinfo"
}
