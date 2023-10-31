package models

import (
	"time"
)

func (LeadPostHistory) TableName() string {
	return "lead_post_history"
}

type LeadPostHistory struct {
	Id                 int64                  `json:"id" gorm:"column:postId;type:bigint(25);primary_key;AUTO_INCREMENT"`
	BuyerId            int64                  `json:"buyer_id" gorm:"column:buyer_id;type:bigint;index:idx_created_at_buyer_id_status"`
	TraceId            string                 `json:"trace_id" gorm:"column:trace_id;index:idx_trace_id"`
	BuyerGroupId       int64                  `json:"buyer_group_id" gorm:"column:buyer_group_id;type:bigint"`
	CustomerId         int64                  `json:"customer_id" gorm:"column:customer_id;type:bigint;index:idx_customer_id;index:idx_rule_manifest_id_customer_id"`
	KnowledgeBaseName  string                 `json:"knowledgebase_name" gorm:"column:knowledge_base_name;type:varchar(150)"`
	RuleManifestId     string                 `json:"rule_manifest_id" gorm:"column:rule_manifest_id;type:varchar(255);index:idx_rule_manifest_id_customer_id"`
	RuleManifestStatus string                 `json:"rule_manifest_status" gorm:"column:rulemanifest_status;type:varchar(100)"`
	ResponseCode       int64                  `json:"response_code" gorm:"column:response_code;type:int"`
	Status             string                 `json:"status" gorm:"column:status;type:varchar(255);index:idx_created_at_buyer_id_status"`
	Response           map[string]interface{} `json:"response" gorm:"-"`
	ResponseDb         *string                `json:"-" gorm:"column:response;type:TEXT"`
	CreatedAt          time.Time              `json:"created_at" gorm:"column:created_at;index:idx_created_at_buyer_id_status"`
	UpdatedAt          time.Time              `json:"updated_at" gorm:"column:updated_at"`
}
