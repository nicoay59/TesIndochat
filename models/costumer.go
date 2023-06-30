package models

type Customer struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment"`
	Name     string `json:"name" gorm:"varchar(255)"`
	Email    string `json:"email" gorm:"varchar(255)"`
	Password string `json:"password" gorm:"varchar(255)"`
}

type CustomerResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (CustomerResponse) TableName() string {
	return "customers"
}
