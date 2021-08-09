package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id int `json:"id,omitempty",gorm:"type:int"`
	Name string `json:"name,omitempty",gorm:"varchar(25);not null"`
	Password string `json:"password,omitempty",gorm:"varchar(100);not null"`
	Phone string `json:"phone,omitempty",gorm:"size(255);not null"`
}
