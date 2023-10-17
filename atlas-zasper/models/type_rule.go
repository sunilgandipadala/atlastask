package models

import (
	"time"
)

//	func (Ruledata) TableName() string {
//		return "rules"
//	}
func (Rules) TableName() string {
	return "rules_Info"
}

type Rule struct {
	Name       string         `json:"name,omitempty"`
	Conditions []Condition    `json:"conditions,omitempty"`
	Manifest   *RulesManifest `json:"then,omitempty"`
}

type Rules struct {
	Id        int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	FlowName  string    `gorm:"column:flowName1;type:varchar(255)"`
	Manifest  string    `gorm:"column:manifest;type:TEXT"`
	CreatedAt time.Time `gorm:"column:created_at;index:idx_name_created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type MyPoGo struct {
	Result string
}
