package models

type RulesManifest struct {
	Id               string `json:"id,omitempty"`
	Completed        string `json:"completed,omitempty"`
	IsConditionBlock bool   `json:"is_condition_block,omitempty"`
	IsActionBlock    bool   `json:"is_action_block,omitempty"`
	//ActionDetailsManifest    *ActionDetails    `json:"action_details_manifest,omitempty"`
	ConditionDetailsManifest *ConditionDetails `json:"condition_details_manifest,omitempty"`
	KnowledgeBaseName        string            `json:"knowledge_base_name,omitempty"`
	TraceId                  string            `json:"trace_id,omitempty"`
}
