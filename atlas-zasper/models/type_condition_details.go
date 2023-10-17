package models

type ConditionDetails struct {
	ConditionId   string      `json:"condition_id,omitempty"`
	ConditionName string      `json:"condition_name,omitempty"`
	Conditions    []Condition `json:"conditions,omitempty"`
}

type Condition struct {
	DataAttribute string `json:"dataAttribute"`
	Operator      string `json:"operator"`
	Value         string `json:"value"`
	Action        string `json:"action"`
}
