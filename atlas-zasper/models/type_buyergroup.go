package models

import "time"

type BuyerGroups []*BuyerGroup

func (BuyerGroup) TableName() string {
	return "buyergroups"
}

type BuyerGroup struct {
	Id        int64     `json:"id" gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT"`
	GroupName string    `json:"groupName" gorm:"column:groupName;type:varchar(100);uniqueindex;not null"`
	BuyerIds  string    `json:"buyerIds" gorm:"column:buyerIds;type:varchar(100)"`
	Priority  string    `json:"priority" gorm:"column:priority;type:varchar(100)"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;index:idx_created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}
