package models

import (
	"time"
)

type RuleTopicDetails []RuleTopicDetails

type RuleTopicDetail struct {
	RulesManifest          *RulesManifest          `json:"rules_manifest,omitempty,omitempty"`
	TraceId                string                  `json:"traceId,omitempty"`
	Cid                    string                  `json:"cid,omitempty"`
	Status                 string                  `json:"status,omitempty"`
	Endpoint               string                  `json:"endpoint,omitempty"`
	Method                 string                  `json:"method,omitempty"`
	QueryParamsMap         *map[string]interface{} `json:"queryParams,omitempty"`
	CustomerDetails        *Customer               `json:"customer_details,omitempty"`
	CustomerProfileDetails *CustomerProfile        `json:"customer_profile_details,omitempty"`
	LeadUserDetails        *LeadUser               `json:"lead_user_details,omitempty"`
	GoRoutine              string                  `json:"goRoutine,omitempty"`
	TimePicked             time.Time               `json:"timePicked,omitempty"`
	TimeFired              time.Time               `json:"timeFired,omitempty"`
	TimeEnded              time.Time               `json:"timeEnded,omitempty"`
}
