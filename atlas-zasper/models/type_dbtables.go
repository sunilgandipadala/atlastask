package models

type DatabaseTableInterface interface {
	TableName() string
}

var DATABASE_APPLICATION_TABLES = []DatabaseTableInterface{
	Customer{},
	WorkFlowJson{},
	BuyerGroup{},
	Buyer{},
	LeadPostHistory{},
	Rules{},
	UserCredentials{},
	CustomerProfile{},
	WorkFlowExecution{},
}
