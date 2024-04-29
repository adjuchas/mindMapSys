package models

import "time"

type Message struct {
	MessageID  int       `gorm:"column:messageID;auto_increment;primary" json:"messageID"`
	KsTea      int       `gorm:"column:ks_tea" json:"ks_tea"`
	KsStu      int       `gorm:"column:ks_stu" json:"ks_stu"`
	Msg        string    `gorm:"column:msg" json:"msg"`
	DraftID    int       `gorm:"column:draftID" json:"draftID"`
	CreateTime time.Time `gorm:"column:createTime;type:datetime(0);autoCreateTime" json:"createTime"`
}

func (receiver Message) TableName() string {
	return "message"
}
