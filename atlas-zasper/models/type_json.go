package models

type Trigger struct {
	ID       string       `json:"id"`
	Type     string       `json:"type"`
	Data     TriggerData  `json:"data"`
	Children []*ChildNode `json:"children"`
	Meta     *interface{} `json:"meta,omitempty"`
}

type TriggerData struct {
	ID       string       `json:"id"`
	NodeData *NodeDataDef `json:"nodeData,omitempty"`
	Meta     *interface{} `json:"meta,omitempty"`
}

type NodeDataDef struct {
	Title         string        `json:"title,omitempty"`
	Description   string        `json:"description,omitempty"`
	CreateTime    int64         `json:"createTime,omitempty"`
	Name          string        `json:"name,omitempty"`
	BuyerGroups   []int64       `json:"buyerGroups,omitempty"`
	Limit         string        `json:"limit,omitempty"`
	ShowJSONInput bool          `json:"showJSONInput,omitempty"`
	ResponseModel []ResponseDef `json:"responseModel,omitempty"`
}

type BuyerNodeDef struct {
	ID          string                     `json:"id"`
	EndFlow     bool                       `json:"endFlow"`
	BuyerGroups []AttachedBuyerGroup       `json:"attachedBuyerGroups"`
	MapResponse map[string]ResponseMapping `json:"map_response,omitempty"`
	Meta        *interface{}               `json:"meta,omitempty"`
}

type AttachedBuyerGroup struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Buyers        []int  `json:"buyers"`
	Active        bool   `json:"active"`
	BuyerPriority int    `json:"buyerPriority"`
}

type ResponseMapping struct {
	MappedValue string `json:"mappedValue"`
	Response    string `json:"response"`
}

type ResponseDef struct {
	DataType     string `json:"dataType"`
	PropertyName string `json:"propertyName"`
}

type ChildNode struct {
	ID       string        `json:"id"`
	Type     string        `json:"type"`
	Data     ChildNodeData `json:"data"`
	Children []*ChildNode  `json:"children"`
	Meta     *interface{}  `json:"meta,omitempty"`
}

type ChildNodeData struct {
	ConditionName string         `json:"title"`
	Conditions    []ConditionDef `json:"actions"`
	Meta          *interface{}   `json:"meta,omitempty"`
	BuyerNodeDef
}

type ConditionDef struct {
	DataAttribute string `json:"attribute"`
	Operator      string `json:"operator"`
	Value         string `json:"value"`
	Action        string `json:"action"`
}

type Root struct {
	Root       Trigger       `json:"root"`
	Connectors []interface{} `json:"connectors"`
}

type Nested struct {
	IsNested bool `json:"isNested,omitempty"`
}

type MetaData struct {
	Key      string   `json:"key"`
	ID       string   `json:"id"`
	Data     Data     `json:"data"`
	Position Position `json:"position"`
	Type     string   `json:"type"`
}

type Data struct {
	ID string `json:"id"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}
