package models

import "time"

type DraftInfo struct {
	DraftID      int       `gorm:"column:draft_ID;primaryKey;auto_increment" json:"draftID"`
	TITLE        string    `gorm:"column:TITLE" json:"title"`
	KsID         int       `gorm:"column:ks_ID" json:"ksID"`
	STATE        string    `gorm:"column:STATE" json:"state"`
	TAGS         string    `gorm:"column:TAGS" json:"tags"`
	DESCRIPTION  string    `gorm:"column:DESCRIPTION" json:"description"`
	NodeTreePath string    `gorm:"column:nodeTreePath" json:"nodeTreePath"`
	CreateTime   time.Time `gorm:"column:CREATETIME;type:datetime(0);autoCreateTime" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:UPDATETIME;type:datetime(0);autoUpdateTime" json:"updateTime"`
}

func (receiver DraftInfo) TableName() string {
	return "draftinfo"
}
