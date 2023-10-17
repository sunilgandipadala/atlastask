package models

import (
	"time"
)

type LeadTopicDetail struct {
	TraceId                string           `json:"traceId,omitempty"`
	RulesManifest          *RulesManifest   `json:"rules_manifest,omitempty"`
	CustomerDetails        *Customer        `json:"customerdetails,omitempty"`
	CustomerProfileDetails *CustomerProfile `json:"customer_profile_details,omitempty"`
	LeadUserDetails        *LeadUser        `json:"lead_user_details,omitempty"`
	Cid                    string           `json:"cid,omitempty"`
	Status                 string           `json:"status,omitempty"`
	FlowEnded              bool             `json:"flowEnded,omitempty"`
	StoreResponse          string           `json:"storeResponse,omitempty"`
	TimePicked             time.Time        `json:"timePicked,omitempty"`
	TimeFired              time.Time        `json:"timeFired,omitempty"`
	TimeEnded              time.Time        `json:"timeEnded,omitempty"`
}
