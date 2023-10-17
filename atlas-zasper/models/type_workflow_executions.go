package models

import (
	"time"

	"github.com/google/uuid"
)

type WorkFlowExecutions []*WorkFlowExecution

func (WorkFlowExecution) TableName() string {
	return "workflow_execution"
}

type WorkFlowExecution struct {
	TraceId    uuid.UUID `gorm:"type:uuid"`
	WorkflowId uuid.UUID `gorm:"type:uuid"`
	Payload    string    `gorm:"type:TEXT"`
	CreatedAt  time.Time `json:"created_at,omitempty" gorm:"column:created_at;index:idx_name_created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}
