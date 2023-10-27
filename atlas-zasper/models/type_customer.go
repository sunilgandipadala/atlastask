package models

import (
	"time"
)

func (Customer) TableName() string {
	return "customers"
}

func (Customer) GetObject() Customer {
	return Customer{}
}

type Customer struct {
	ID                 int64      `json:"id" gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	FirstName          *string    `json:"first_name,omitempty" gorm:"column:first_name;type:varchar(255)"`
	LastName           *string    `json:"last_name,omitempty" gorm:"column:last_name;type:varchar(255)"`
	Age                *int64     `json:"age,omitempty" gorm:"column:age"`
	Email              string     `json:"email,omitempty" gorm:"column:email;type:varchar(50);uniqueindex;not null"`
	Gender             *string    `json:"gender,omitempty" gorm:"column:gender;type:varchar(10)"`
	Address1           *string    `json:"address1,omitempty" gorm:"column:address1;type:varchar(255)"`
	Address2           *string    `json:"address2,omitempty" gorm:"column:address2;type:varchar(255)"`
	City               *string    `json:"city,omitempty" gorm:"column:city;type:varchar(255)"`
	State              *string    `json:"state,omitempty" gorm:"column:state;type:varchar(100)"`
	Country            *string    `json:"country,omitempty" gorm:"column:country;type:varchar(100)"`
	ZipCode            *string    `json:"zipcode,omitempty" gorm:"column:zipcode;type:varchar(50)"`
	Phone              *string    `json:"phone,omitempty" gorm:"column:phone;type:varchar(20);index:idx_phone"`
	Subscribed         *bool      `json:"subscribed,omitempty" gorm:"column:subscribed"`
	SubscribedAt       *time.Time `json:"subscribed_at" gorm:"column:subscribed_at"`
	S1                 *int64     `json:"s1,omitempty" gorm:"column:s1;type:bigint"`
	S2                 *int64     `json:"s2,omitempty" gorm:"column:s2;type:bigint"`
	S3                 *string    `json:"s3,omitempty" gorm:"column:s3;type:varchar(50)"`
	Q                  *string    `json:"q,omitempty" gorm:"column:q;type:varchar(255)"`
	L                  *string    `json:"l,omitempty" gorm:"column:l;type:varchar(255)"`
	OngageContactId    *string    `json:"ongage_contact_id,omitempty" gorm:"column:ongage_contact_id;type:varchar(255)"`
	OngageStatus       *int64     `json:"ongage_status,omitempty" gorm:"column:ongage_status;type:bigint"`
	OversightContactId *string    `json:"oversight_contact_id,omitempty" gorm:"column:oversight_contact_id;type:varchar(255)"`
	OversightStatus    *int64     `json:"oversight_status,omitempty" gorm:"column:oversight_status;type:bigint"`
	Validation         *string    `json:"validation,omitempty" gorm:"column:validation;type:varchar(255)"`
	ValidationDate     *time.Time `json:"validation_date,omitempty" gorm:"column:validation_date"`
	Processed          bool       `json:"processed,omitempty" gorm:"column:processed"`
	ProcessedDate      *time.Time `json:"processed_date,omitempty" gorm:"column:processed_date"`
	Tz                 *string    `json:"tz,omitempty" gorm:"column:tz"`
	OptInPhone         *bool      `json:"opt_in_phone,omitempty" gorm:"column:opt_in_phone"`
	Dob                *string    `json:"dob,omitempty" gorm:"column:dob"`
	ClickIp            *string    `json:"click_ip,omitempty" gorm:"column:click_ip"`
	PushnamiId         *string    `json:"pushnami_id,omitempty" gorm:"column:pushnami_id;type:varchar(255)"`
	OptinPhone         *int64     `json:"optin_phone,omitempty" gorm:"column:optin_phone;type:bigint"`
	LastVisit          *time.Time `json:"last_visit,omitempty" gorm:"column:last_visit"`
	UserAgent          *string    `json:"useragent,omitempty" gorm:"column:user_agent;type:varchar(100)"`
	OptinIP            *string    `json:"optin_ip,omitempty" gorm:"column:optin_ip;type:varchar(50)"`
	OptinURL           *string    `json:"optin_url,omitempty" gorm:"column:optin_url;type:varchar(100)"`
	OptinTime          *time.Time `json:"optin_time,omitempty" gorm:"column:optin_time"`
	UTMSource          *string    `json:"utm_source,omitempty" gorm:"column:utm_source;type:varchar(100)"`
	UTMMedium          *string    `json:"utm_medium,omitempty" gorm:"column:utm_medium;type:varchar(100)"`
	UTMCampaign        *string    `json:"utm_campaign,omitempty" gorm:"column:utm_campaign;type:varchar(100)"`
	CreatedAt          time.Time  `json:"created_at,omitempty" gorm:"index:idx_created_at"`
	UpdatedAt          time.Time  `json:"updated_at,omitempty"`
	Pet_Name           *string    `json:"pet_name,omitempty" gorm:"column:pet_name;type:varchar(255)"` //it is newly added---
}

type Customers []*Customers
