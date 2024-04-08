package models

type Tag struct {
	TagID      string `gorm:"column:tag_ID"`
	Name       string `gorm:"column:NAME"`
	DotNum     string `gorm:"column:dot_NUM"`
	LeaderID   string `gorm:"column:leader_ID"`
	CreateTime string `gorm:"column:CREATETIME"`
	UpdateTime string `gorm:"column:UPDATETIME"`
}

func (receiver Tag) TableName() string {
	return "taginfo"
}
