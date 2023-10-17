package models

// We have to work on this struct from now with MYSQL DB
type Employee struct {
	EmpNo     int    `gorm:"primarykey;type:INT"`
	BirthDate string `gorm:"type:date;NOT NULL"`
	FirstName string `gorm:"type:Varchar(14);NOT NULL"`
	LastName  string `gorm:"type:Varchar(16);NOT NULL"`
	HireDate  string `gorm:"type:date;NOT NULL"`
	Gender    string `gorm:"NOT NULL;type:ENUM ('M','F')"`
}

//Lets define remaining structs also
type Departments struct {
	DeptNo   string `gorm:"primarykey;type:char(4)"`
	DeptName string `gorm:"unique;type:varchar(40);NOT NULL""`
}

type Salaries struct {
	EmpNo    int    `gorm:"primarykey;type:INT"`
	Salary   int    `gorm:"not null;type:INT"`
	FromDate string `gorm:"type:date;primarykey;not null"`
	ToDate   string `gorm:"type:date;not null"`

	// Define foreign key
	Employee Employee `gorm:"foreignKey:EmpNo;constraint:OnUpdate:NO ACTION,OnDelete:CASCADE"`
}

type Titles struct {
	EmpNo    int     `gorm:"primarykey;type:INT"`
	Title    string  `gorm:"primarykey;size:50;not null"`
	FromDate string  `gorm:"type:date;primarykey;not null"`
	ToDate   *string `gorm:"type:date"` // Use a pointer for nullable date

	// Define foreign key
	Employee Employee `gorm:"foreignKey:EmpNo;constraint:OnUpdate:NO ACTION,OnDelete:CASCADE"`
}

type DeptEmp struct {
	EmpNo    int    `gorm:"primarykey"`
	DeptNo   string `gorm:"primarykey;type:char(4);index"`
	FromDate string `gorm:"NOT NULL;type:date"`
	ToDate   string `gorm:"NOT NULL;type:date"`
	// Define foreign keys
	Employee   Employee    `gorm:"foreignKey:EmpNo;constraint:OnUpdate:NO ACTION,OnDelete:CASCADE"`
	Department Departments `gorm:"foreignKey:DeptNo;constraint:OnUpdate:NO ACTION,OnDelete:CASCADE"`
}

type DeptManager struct {
	EmpNo    int    `gorm:"primarykey"`
	DeptNo   string `gorm:"primarykey;type:char(4);index"`
	FromDate string `gorm:"NOT NULL;type:date"`
	ToDate   string `gorm:"NOT NULL;type:date"`
	// Define foreign keys
	Employee   Employee    `gorm:"foreignKey:EmpNo;constraint:OnUpdate:NO ACTION,OnDelete:CASCADE"`
	Department Departments `gorm:"foreignKey:DeptNo;constraint:OnUpdate:NO ACTION,OnDelete:CASCADE"`
}
