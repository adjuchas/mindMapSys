package models

type DraftInfo struct {
	DraftID      int    `gorm:"column:draft_ID;primaryKey" json:"draftID"`
	TITLE        string `gorm:"column:TITLE" json:"title"`
	KsID         string `gorm:"column:ks_ID" json:"ksID"`
	STATE        string `gorm:"column:STATE" json:"state"`
	TAGS         string `gorm:"column:TAGS" json:"tags"`
	DESCRIPTION  string `gorm:"column:DESCRIPTION" json:"description"`
	NodeTreePath string `gorm:"column:nodeTreePath" json:"nodeTreePath"`
	CreateTime   string `gorm:"column:CREATETIME" json:"createTime"`
	UpdateTime   string `gorm:"column:UPDATETIME" json:"updateTime"`
}

func (receiver DraftInfo) TableName() string {
	return "draftinfo"
}
