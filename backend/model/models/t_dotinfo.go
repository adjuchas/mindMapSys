package models

type DotInfo struct {
	DotID        string `gorm:"column:dot_ID" json:"dotID"`
	TITLE        string `gorm:"column:TITLE" json:"title"`
	KsID         string `gorm:"column:ks_ID" json:"ksID"`
	STATE        string `gorm:"column:STATE" json:"state"`
	TAGS         string `gorm:"column:TAGS" json:"tags"`
	DESCRIPTION  string `gorm:"column:DESCRIPTION" json:"description"`
	NodeTreePath string `gorm:"column:nodeTreePath" json:"nodeTreePath"`
	PreViewPath  string `gorm:"column:preViewPath" json:"preViewPath"`
	CreateTime   string `gorm:"column:CREATETIME" json:"createTime"`
	UpdateTime   string `gorm:"column:UPDATETIME" json:"updateTime"`
}

func (receiver DotInfo) TableName() string {
	return "dotinfo"
}
