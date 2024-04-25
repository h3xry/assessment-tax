package model

type Deductions struct {
	Name   string  `json:"name" gorm:"unique,type:varchar(20)"`
	Amount float64 `json:"amount" gorm:"type:decimal(10,2)"`
}
