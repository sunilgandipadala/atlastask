package models

func (User) TableName() string {
	return "userdetails"
}

func (UserCredentials) TableName() string {
	return "userdetails"
}

type User struct {
	Id                int64  `json:"id" gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT"`
	UserName          string `json:"userName" gorm:"column:userName;type:varchar(100)"`
	Email             string `json:"email" gorm:"column:email;type:varchar(50)"`
	Fid               string `json:"Fid" gorm:"column:fid;type:varchar(100)"`
	Age               int64  `json:"age" gorm:"column:age"`
	Phone             string `json:"phone" gorm:"column:phone;type:varchar(13)"`
	Address           string `json:"address" gorm:"column:address;type:varchar(500)"`
	CarBuyingInterest bool   `json:"carBuyingInterest" gorm:"column:carBuyingInterest"`
	HomeBuyInterest   bool   `json:"homeBuyInterest" gorm:"column:homeBuyInterest"`
}

type Users []*UserCredentials

type UserCredentials struct {
	Id       int64  `json:"id" gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT"`
	UserName string `json:"userName" gorm:"column:userName;type:varchar(100)"`
	Password string `json:"password" gorm:"column:password;type:varchar(100)"`
	Email    string `json:"email" gorm:"column:email;type:varchar(50)"`
}
