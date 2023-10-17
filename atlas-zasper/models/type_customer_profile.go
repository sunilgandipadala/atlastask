package models

import (
	"time"
)

func (CustomerProfile) TableName() string {
	return "customer_profiles"
}

type CustomerProfiles []*CustomerProfile

type CustomerProfile struct {
	CustomerId                    int64                    `json:"customer_id,omitempty" gorm:"primary_key"`
	CustomerProfileManifestStruct *CustomerProfileManifest `json:"customer_profile_manifest" gorm:"embedded"`
	AdditionalFieldsStruct        map[string]*string       `json:"additional_fields,omitempty" gorm:"-"`
	AdditionalFieldsDb            *string                  `json:"-" gorm:"additionalFieldsManifest,type:TEXT"`
	CreatedAt                     time.Time                `json:"created_at,omitempty"`
	UpdatedAt                     time.Time                `json:"updated_at,omitempty"`
}

type CustomerProfileManifest struct {
	UserID                 *int64     `json:"user_id" gorm:"column:user_id"`
	EduLevel               *string    `json:"edu_level,omitempty" gorm:"column:edu_level;type:varchar(255)"`
	EduInterest            *string    `json:"edu_interest,omitempty" gorm:"column:edu_interest;type:varchar(255)"`
	EduUpdatedAt           *time.Time `json:"edu_updated_at,omitempty" gorm:"column:edu_updated_at"`
	DeletedAt              *time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
	TestID                 *int64     `json:"test_id,omitempty" gorm:"column:test_id;type:bigint"`
	HealthInsurance        *string    `json:"health_insurance,omitempty" gorm:"column:health_insurance;type:varchar(255)"`
	HealthMultiplePresc    *string    `json:"health_multiple_prescription,omitempty" gorm:"column:health_multiple_prescription;type:varchar(255)"`
	HealthDiabetic         *string    `json:"health_diabetic,omitempty" gorm:"column:health_diabetic;type:varchar(255)"`
	HealthCPAP             *string    `json:"health_cpap,omitempty" gorm:"column:health_cpap;type:varchar(255)"`
	HealthLowIncome        *string    `json:"health_low_income,omitempty" gorm:"column:health_low_income;type:varchar(255)"`
	HealthBackPain         *string    `json:"health_back_pain,omitempty" gorm:"column:health_back_pain;type:varchar(255)"`
	HealthUpdatedAt        *time.Time `json:"health_updated_at,omitempty" gorm:"column:health_updated_at"`
	FinanceCheckingAccount *string    `json:"finance_checking_account,omitempty" gorm:"column:finance_checking_account;type:varchar(255)"`
	FinanceDebtAmount      *int64     `json:"finance_debt_amount,omitempty" gorm:"column:finance_debt_amount;type:varchar(255)"`
	FinanceDebt            *bool      `json:"finance_debt,omitempty" gorm:"column:finance_debt"`
	FinanceUpdatedAt       *time.Time `json:"finance_updated_at,omitempty" gorm:"column:finance_updated_at"`
	CarOwner               *string    `json:"car_owner,omitempty" gorm:"column:car_owner;type:varchar(255)"`
	CarUpdatedAt           *time.Time `json:"car_updated_at,omitempty" gorm:"column:car_updated_at"`
	CarInsurance           *string    `json:"car_insurance,omitempty" gorm:"column:car_insurance;type:varchar(255)"`
	CarLicense             *string    `json:"car_license,omitempty" gorm:"column:car_license"`
	WorkStatus             *string    `json:"work_status,omitempty" gorm:"column:work_status;type:varchar(255)"`
	WorkUpdatedAt          *time.Time `json:"work_updated_at,omitempty" gorm:"column:work_updated_at"`
	EduJornayaToken        *string    `json:"edu_jornaya_token,omitempty" gorm:"column:edu_jornaya_token;type:varchar(255)"`
	CaptchaScore           *float64   `json:"captcha_score,omitempty" gorm:"column:captcha_score"`
	CaptchaUpdatedAt       *time.Time `json:"captcha_updated_at,omitempty" gorm:"column:captcha_updated_at"`
	NavyProgram            *string    `json:"navy_program,omitempty" gorm:"column:navy_program;type:varchar(255)"`
	USCitizen              *bool      `json:"us_citizen,omitempty" gorm:"column:us_citizen"`
	GPAAverage             *float64   `json:"gpa_average,omitempty" gorm:"column:gpa_average"`
	FQScore                *int64     `json:"fq_score,omitempty" gorm:"column:fq_score;type:bigint"`
	FQUpdatedAt            *time.Time `json:"fq_updated_at,omitempty" gorm:"column:fq_updated_at"`
	Ranch                  *string    `json:"ranch,omitempty" gorm:"column:ranch;type:varchar(255)"`
	BRTAFilt               *string    `json:"brta_filt,omitempty" gorm:"column:brta_filt;type:varchar(255)"`
	GLAD1                  *string    `json:"glad1,omitempty" gorm:"column:glad1;type:varchar(255)"`
	GLAD2                  *string    `json:"glad2,omitempty" gorm:"column:glad2;type:varchar(255)"`
	HomeOwner              *string    `json:"home_owner,omitempty" gorm:"column:home_owner;type:varchar(255)"`
	HomeUpdatedAt          *time.Time `json:"home_updated_at,omitempty" gorm:"column:home_updated_at"`
	DebtType               *string    `json:"debt_type,omitempty" gorm:"column:debt_type;type:varchar(255)"`
	EduGrant               *string    `json:"edu_grant,omitempty" gorm:"column:edu_grant;type:varchar(255)"`
	Finex                  *string    `json:"finex,omitempty" gorm:"column:finex;type:varchar(255)"`
	Demply                 *string    `json:"demply,omitempty" gorm:"column:demply;type:varchar(255)"`
	CLEJ                   *string    `json:"clej,omitempty" gorm:"column:clej;type:varchar(255)"`
	AutoAccident           *string    `json:"auto_accident,omitempty" gorm:"column:auto_accident;type:varchar(255)"`
	Accident               *string    `json:"accident,omitempty" gorm:"column:accident;type:varchar(50)"`
	Medicaid               *string    `json:"medicaid,omitempty" gorm:"column:medicaid;type:varchar(50)"`
	Medicare               *string    `json:"medicare,omitempty" gorm:"column:medicare;type:varchar(50)"`
	AutoYear               *string    `json:"auto_year,omitempty" gorm:"column:auto_year;type:varchar(50)"`
	AutoMake               *string    `json:"auto_make,omitempty" gorm:"column:auto_make;type:varchar(50)"`
	AutoModel              *string    `json:"auto_model,omitempty" gorm:"column:auto_model;type:varchar(50)"`
}
