package models

type RuleGeneration struct {
	When []*ConditionDef `json:"when"`
	Then RulesManifest   `json:"then"`
	Name string          `json:"name"`
}

type ThenScope struct {
	Conditions  []*ConditionDef
	TriggerData *TriggerData `json:"triggerData"`
}
