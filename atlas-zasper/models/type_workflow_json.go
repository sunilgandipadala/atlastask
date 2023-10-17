package models

import (
	"time"

	"github.com/google/uuid"
)

type WorkFlowJsonDetails []*WorkFlowJson

func (WorkFlowJson) TableName() string {
	return "work_flow_json_details"
}

type WorkFlowJson struct {
	Id     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name   string    `json:"name,omitempty" gorm:"column:name;index:idx_name_created_at;uniqueindex;not null"`
	JsonDb *string   `json:"-" gorm:"column:json"`
	//JsonStruct *Root     `json:"json" gorm:"-"`
	JsonStruct *string   `json:"json_struct,omitempty" gorm:"-"`
	CreatedAt  time.Time `json:"created_at,omitempty" gorm:"column:created_at;index:idx_name_created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}
