package models

import (
	"time"
)

func (Buyer) TableName() string {
	return "buyers"
}

type Buyers []*Buyer

type Buyer struct {
	Id      int64  `json:"id" gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT"`
	PetName string `json:"name" gorm:"column:name_new;type:varchar(255);`
	//Name              string                 `json:"name" gorm:"column:name_new;type:varchar(255);uniqueindex;not null"`
	Price             float64                `json:"price" gorm:"column:pricees"`
	WebhookUrl        string                 `json:"webhookUrl" gorm:"column:webhookUrl;type:text"`
	WebhookMethod     string                 `json:"webhookMethod" gorm:"column:webhookMethod;type:varchar(20)"`
	WebhookTimeout    string                 `json:"webhookTimeout" gorm:"column:webhookTimeout;type:varchar(30)"`
	HeadersStruct     []Header               `json:"headers" gorm:"-"`
	HeadersDb         *string                `json:"-" gorm:"column:headersnew;type:text"`
	RequestBodyStruct map[string]interface{} `json:"requestBody" gorm:"-"`
	RequestBodyDb     *string                `json:"-" gorm:"column:requestBody;type:text"`
	HttpAuthUsername  string                 `json:"httpAuthUsername" gorm:"column:httpAuthUsername;type:varchar(20)"`
	//HttpAuthPassword    string                 `json:"httpAuthPassword" gorm:"column:httpAuthPassword;type:varchar(20)"`
	ResponseModelStruct []ResponseDef `json:"responseModel" gorm:"-"`
	ResponseModelDb     *string       `json:"-" gorm:"column:responseModel;type:text"`
	BuyerGroupId        int64         `json:"-" gorm:"-"`
	CreatedAt           time.Time     `json:"created_at" gorm:"column:created_at;index:idx_created_at"`
	UpdatedAt           time.Time     `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

type Header struct {
	HeaderName  string `json:"headerName" gorm:"column:headerName;type:varchar(100)"`
	HeaderValue string `json:"headerValue" gorm:"column:headerValue;type:varchar(100)"`
}

type RequestBody struct {
}
