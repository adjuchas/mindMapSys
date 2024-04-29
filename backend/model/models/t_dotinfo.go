package models

import "time"

type DotInfo struct {
	DotID        int       `gorm:"column:dot_ID;auto_increment;primary" json:"dotID"`
	TITLE        string    `gorm:"column:TITLE" json:"title"`
	KsID         int       `gorm:"column:ks_ID" json:"ksID"`
	STATE        string    `gorm:"column:STATE" json:"state"`
	TAGS         string    `gorm:"column:TAGS" json:"tags"`
	DESCRIPTION  string    `gorm:"column:DESCRIPTION" json:"description"`
	NodeTreePath string    `gorm:"column:nodeTreePath" json:"nodeTreePath"`
	PreViewPath  string    `gorm:"column:preViewPath" json:"preViewPath"`
	CreateTime   time.Time `gorm:"column:CREATETIME;type:datetime(0);autoCreateTime" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:UPDATETIME;type:datetime(0);autoUpdateTime" json:"updateTime"`
}

func (receiver DotInfo) TableName() string {
	return "dotinfo"
}
