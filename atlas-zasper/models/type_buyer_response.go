package models

type BuyerApiResponse struct {
	StatusCode int64
	Message    string
	Response   map[string]interface{}
}
