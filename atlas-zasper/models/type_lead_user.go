package models

type LeadUser struct {
	CustomerId      int64  `json:"customer_id" form:"customer_id"`
	UserId          int64  `json:"user_id,omitempty" form:"user_id"`
	Age             int64  `json:"age,omitempty" form:"age"`
	Email           string `json:"email" form:"email"  binding:"required"`
	FirstName       string `json:"first_name,omitempty" form:"first_name"`
	LastName        string `json:"last_name,omitempty" form:"last_name"`
	Gender          string `json:"gender,omitempty" form:"gender"`
	Address1        string `json:"address1,omitempty" form:"address1"`
	Address2        string `json:"address2,omitempty" form:"address2"`
	Zipcode         string `json:"zipcode,omitempty" form:"zipcode"`
	City            string `json:"city,omitempty" form:"city"`
	State           string `json:"state,omitempty" form:"state"`
	County          string `json:"county,omitempty" form:"county"`
	EduLevel        string `json:"edu_level,omitempty" form:"edu_level"`
	EduInterest     string `json:"edu_interest,omitempty" form:"edu_interest"`
	EduUpdatedAt    string `json:"edu_updated_at,omitempty" form:"edu_updated_at"`
	CreatedAt       string `json:"created_at,omitempty" form:"created_at"`
	HealthInsurance string `json:"health_insurance,omitempty" form:"health_insurance"`
	WorkStatus      string `json:"work_status,omitempty" form:"work_status"`
	FinanceDebt     string `json:"finance_debt,omitempty" form:"finance_debt"`
	CarInsurance    string `json:"car_insurance,omitempty" form:"car_insurance"`
	CarLicense      string `json:"car_license,omitempty" form:"car_license"`
	Medicare        string `json:"medicare,omitempty" form:"medicare"`
	HomeOwner       string `json:"home_owner,omitempty" form:"home_owner"`
	CarOwner        string `json:"car_owner,omitempty" form:"car_owner"`
}
